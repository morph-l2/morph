#!/usr/bin/env node
/*
 * rpc-consistency-proxy.js — a tiny JSON-RPC proxy that gives read-after-write consistency on a public RPC.
 *
 * Problem: public RPC endpoints sit behind several nodes. A client sends a transaction, sees its receipt from node A,
 * then asks `latest` state from node B which has not seen that block yet and gets stale data or empty return data.
 * Deploy scripts do exactly that (`await contract.deployed(); await contract.owner()`).
 *
 * Fix: remember the highest block we have seen (from receipts and eth_blockNumber) -- the "floor" -- and rewrite every
 * state read that targets `latest`/`pending` to that explicit block number. A node that does not have the block yet
 * answers with an error ("header not found", only eth_estimateGas does that on the endpoint we use) or, for most
 * methods, with `result: null` (a valid eth_call/getBalance/getCode/... result is never null). Either way the proxy
 * waits and retries, but ONLY for requests it rewrote; requests with an explicit block tag pass through unchanged so a
 * client that polls for a future block is not held up. Writes and everything else pass through untouched.
 *
 * Usage: rpc-consistency-proxy.js UPSTREAM_URL[,FALLBACK_URL...] [PORT=18546]      (listens on 127.0.0.1)
 * GET /stats returns counters, including requests per method and rate-limit / failover counts.
 */
const http = require("http");

const upstreams = (process.argv[2] || "").split(",").map((u) => u.trim()).filter(Boolean);
const port = Number(process.argv[3] || 18546);
if (upstreams.length === 0) { console.error("usage: rpc-consistency-proxy.js UPSTREAM_URL[,FALLBACK_URL...] [PORT]"); process.exit(2); }
let current = 0;   // index of the upstream that last answered; sticky until it is rate-limited or unreachable
const MIN_INTERVAL_MS = Number(process.env.PROXY_MIN_INTERVAL_MS || 0);      // optional pacing between upstream calls (0 = off)
const HEAD_REFRESH_MS = Number(process.env.PROXY_HEAD_REFRESH_MS || 1000);      // how often the floor follows the chain head
let nextSlot = 0;

let floor = 0n;
const startAhead = BigInt(process.env.PROXY_TEST_FLOOR_AHEAD || 0);   // testing aid: pretend we already saw head+N
const hex = (n) => "0x" + n.toString(16);
const bump = (v) => { try { const n = BigInt(v); if (n > floor) floor = n; } catch (_) { /* ignore */ } };
const sleep = (ms) => new Promise((r) => setTimeout(r, ms));
const stats = { requests: 0, rewritten: 0, retries: 0, rateLimited: 0, networkErrors: 0, failovers: 0, byMethod: {} };

// method -> index of the block-tag parameter
const TAG_INDEX = {
  eth_call: 1, eth_getCode: 1, eth_getBalance: 1, eth_getStorageAt: 2,
  eth_getTransactionCount: 1, eth_estimateGas: 1, eth_getBlockByNumber: 0,
};
const FLOATING = new Set([undefined, null, "latest", "pending", "safe", "finalized"]);
const NOT_YET = /header not found|missing trie node|unknown block|block not found|not available|no state|future block|too far|beyond current head|extends beyond|after last block/i;

const RATE_LIMITED = /rate.?limit|too many requests|\b429\b|quota|exceeded/i;

// One request to the upstreams. A rate-limit answer (HTTP 429/503 or a JSON-RPC error that says so) or a network error
// moves on to the next upstream at once; when every upstream refuses, back off exponentially (0.5 s .. 8 s) and go
// round again, up to 8 rounds. Anything else -- including ordinary JSON-RPC errors such as a revert -- is the answer.
async function callUpstream(req) {
  stats.byMethod[req.method] = (stats.byMethod[req.method] || 0) + 1;
  if (MIN_INTERVAL_MS > 0) {                       // global pacing: at most one upstream call per MIN_INTERVAL_MS
    const now = Date.now();
    nextSlot = Math.max(nextSlot, now) + MIN_INTERVAL_MS;
    const wait = nextSlot - MIN_INTERVAL_MS - now;
    if (wait > 0) await sleep(wait);
  }
  let delay = 500;
  for (let round = 0; round < 8; round++) {
    for (let k = 0; k < upstreams.length; k++) {
      const idx = (current + k) % upstreams.length;
      try {
        const r = await fetch(upstreams[idx], { method: "POST", headers: { "content-type": "application/json" }, body: JSON.stringify(req) });
        const text = await r.text();
        let j = null;
        try { j = JSON.parse(text); } catch (_) { /* not JSON */ }
        const limited = r.status === 429 || r.status === 503 ||
          (j && j.error && RATE_LIMITED.test(String(j.error.message || "")));
        if (j && !limited) {
          if (idx !== current) { stats.failovers++; console.error(`[failover] now using ${upstreams[idx]}`); current = idx; }
          return j;
        }
        stats.rateLimited++;
        if (stats.rateLimited % 5 === 1) console.error(`[limited] ${req.method} on ${upstreams[idx]} (HTTP ${r.status}): ${j && j.error ? j.error.message : text.slice(0, 80)}`);
      } catch (e) {
        stats.networkErrors++;
        if (stats.networkErrors % 5 === 1) console.error(`[network] ${upstreams[idx]}: ${e.message}`);
      }
    }
    await sleep(delay);
    delay = Math.min(delay * 2, 8000);
  }
  return { jsonrpc: "2.0", id: req.id, error: { code: -32005, message: "proxy: every upstream is rate-limited or unreachable" } };
}

function rewrite(req) {
  if (floor === 0n) return req;
  const p = Array.isArray(req.params) ? req.params.slice() : [];
  if (req.method in TAG_INDEX) {
    const i = TAG_INDEX[req.method];
    // "finalized"/"safe" are deliberately left alone: they are old enough to be consistent everywhere
    if (p[i] === "latest" || p[i] === "pending" || p[i] === undefined || p[i] === null) {
      p[i] = hex(floor);
      stats.rewritten++;
      return { ...req, params: p };
    }
  } else if (req.method === "eth_getLogs" && p[0] && typeof p[0] === "object") {
    const f = { ...p[0] };
    let changed = false;
    if (f.blockHash === undefined) {
      if (FLOATING.has(f.toBlock) && f.toBlock !== "finalized" && f.toBlock !== "safe") { f.toBlock = hex(floor); changed = true; }
      if (f.fromBlock === "latest" || f.fromBlock === "pending") { f.fromBlock = hex(floor); changed = true; }
    }
    if (!changed) return req;                       // explicit range: nothing was rewritten
    p[0] = f;
    stats.rewritten++;
    return { ...req, params: p };
  }
  return req;
}

// ---- the floor must follow the chain, otherwise a client that polls `latest` for a change it made (and never asks for the
// block number or a receipt) would be pinned to a block from before its own transaction and wait forever.
let lastRefresh = 0;
let refreshing = null;
async function refreshFloor() {
  if (Date.now() - lastRefresh < HEAD_REFRESH_MS) return;
  if (!refreshing) {
    refreshing = (async () => {
      try { const r = await callUpstream({ jsonrpc: "2.0", id: 0, method: "eth_blockNumber", params: [] }); if (r.result) bump(r.result); }
      catch (_) { /* keep the old floor */ }
      lastRefresh = Date.now();
      refreshing = null;
    })();
  }
  await refreshing;
}

// ---- nonces: a state read pinned to a block cannot see transactions that were sent but not mined yet, so remember
// the next nonce of every sender we forwarded a transaction for (found via eth_getTransactionByHash).
const localNonce = new Map();   // lower-case address -> next nonce (BigInt)
async function recordSentTx(hash) {
  for (let i = 0; i < 6; i++) {
    const r = await callUpstream({ jsonrpc: "2.0", id: 0, method: "eth_getTransactionByHash", params: [hash] });
    if (r.result && r.result.from) {
      const from = r.result.from.toLowerCase();
      const next = BigInt(r.result.nonce) + 1n;
      if (next > (localNonce.get(from) || 0n)) localNonce.set(from, next);
      return;
    }
    await sleep(500);
  }
}

// An explicit block number that is <= the floor is a block this proxy has already told the client about (or one the client
// derived from it), so some upstream has it and a lagging backend answering "beyond head" / null is only behind.
// A number above the floor is a genuine request for the future and is passed through untouched.
const isNum = (t) => typeof t === "string" && /^0x[0-9a-fA-F]+$/.test(t);
function explicitWithinFloor(req) {
  const p = Array.isArray(req.params) ? req.params : [];
  if (req.method === "eth_getLogs") {
    const f = p[0];
    if (!f || typeof f !== "object" || f.blockHash !== undefined) return false;
    const nums = [f.fromBlock, f.toBlock].filter(isNum).map((x) => BigInt(x));
    return nums.length > 0 && nums.every((n) => n <= floor);
  }
  if (req.method in TAG_INDEX) return isNum(p[TAG_INDEX[req.method]]) && BigInt(p[TAG_INDEX[req.method]]) <= floor;
  return false;
}

// state read pinned to the floor, retried until some upstream has that block
async function pinnedRead(req) {
  const rewritten = rewrite(req);
  const rw = rewritten === req && explicitWithinFloor(req) ? req : rewritten;
  if (rw === req && !explicitWithinFloor(req)) return callUpstream(req);   // nothing to enforce
  for (let attempt = 0; ; attempt++) {
    const resp = await callUpstream(rw);
    const msg = resp.error ? String(resp.error.message || "") : "";
    const notYet = resp.error ? NOT_YET.test(msg) : (resp.result === null || resp.result === undefined);
    if (!notYet || attempt >= 60) return { ...resp, id: req.id };   // 60 x 500ms = 30 s, then hand back what we got
    stats.retries++;
    if (stats.retries % 10 === 1) console.error(`[retry] ${req.method} @${rw.params && (rw.params[TAG_INDEX[req.method]] || JSON.stringify(rw.params[0]))}: ${msg || "null result"}`);
    await sleep(500);
  }
}

async function handle(req) {
  const m = req.method;
  if (m === "eth_blockNumber") {
    const resp = await callUpstream(req);
    if (resp.result) {
      if (startAhead > 0n && floor === 0n) bump(hex(BigInt(resp.result) + startAhead));
      bump(resp.result); resp.result = hex(floor);   // never go backwards
    }
    return resp;
  }
  if (m === "eth_getTransactionReceipt") {
    const resp = await callUpstream(req);
    if (resp.result && resp.result.blockNumber) bump(resp.result.blockNumber);
    return resp;
  }
  if (m === "eth_sendRawTransaction") {
    const resp = await callUpstream(req);
    if (resp.result && typeof resp.result === "string") await Promise.race([recordSentTx(resp.result), sleep(3000)]);
    return resp;
  }
  if (!(m in TAG_INDEX || m === "eth_getLogs")) return callUpstream(req);

  await refreshFloor();
  if (m === "eth_getTransactionCount" && Array.isArray(req.params) && req.params[1] === "pending") {
    // pending = whatever the (possibly stale) node says, the count at the floor block, and what we sent ourselves
    const addr = String(req.params[0]).toLowerCase();
    const [a, b] = await Promise.all([callUpstream(req), pinnedRead({ ...req, params: [req.params[0], "latest"] })]);
    let best = localNonce.get(addr) || 0n;
    for (const r of [a, b]) if (r && r.result) { try { const v = BigInt(r.result); if (v > best) best = v; } catch (_) { /* ignore */ } }
    return { jsonrpc: "2.0", id: req.id, result: hex(best) };
  }
  return pinnedRead(req);
}

const server = http.createServer((req, res) => {
  if (req.method === "GET" && req.url === "/stats") { res.end(JSON.stringify({ floor: hex(floor), ...stats })); return; }
  const chunks = [];
  req.on("data", (c) => chunks.push(c));
  req.on("end", async () => {
    try {
      const body = JSON.parse(Buffer.concat(chunks).toString("utf8"));
      stats.requests++;
      const out = Array.isArray(body) ? await Promise.all(body.map(handle)) : await handle(body);
      res.setHeader("content-type", "application/json");
      res.end(JSON.stringify(out));
    } catch (e) {
      res.statusCode = 502;
      res.end(JSON.stringify({ jsonrpc: "2.0", id: null, error: { code: -32603, message: "proxy error: " + e.message } }));
    }
  });
});
server.listen(port, "127.0.0.1", () => console.error(`rpc-consistency-proxy ${upstreams.join(" , ")} -> http://127.0.0.1:${port}`));
