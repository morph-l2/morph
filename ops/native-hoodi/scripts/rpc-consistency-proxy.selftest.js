// Self-test for rpc-consistency-proxy.js against a fake chain. Exits non-zero on failure.
const http = require("http"); const { spawn } = require("child_process");
const PROXY = process.argv[2] || require("path").join(__dirname, "rpc-consistency-proxy.js");
const SENDER = "0x" + "aa".repeat(20);
let head = 100; setInterval(() => head++, 1000);           // a chain at 1 block / second
const CHANGE_AT = 103;                                     // contract state changes in block 103
const fake = http.createServer((q, r) => { let b = ""; q.on("data", c => b += c); q.on("end", () => {
  const req = JSON.parse(b); const P = req.params || []; let result = null; const hexn = n => "0x" + n.toString(16);
  const tagNum = t => (t === "latest" || t === "pending" || t === undefined) ? head : parseInt(t, 16);
  switch (req.method) {
    case "eth_chainId": result = "0x88bb0"; break;
    case "eth_getLogs": { const f = P[0] || {}; const hi = Math.max(tagNum(f.fromBlock), tagNum(f.toBlock));
      if (hi > head) { r.setHeader("content-type", "application/json"); r.end(JSON.stringify({ jsonrpc: "2.0", id: req.id, error: { code: -32000, message: "block range extends beyond current head block" } })); return; }
      result = []; break; }
    case "eth_blockNumber": result = hexn(head); break;
    case "eth_call": { const n = tagNum(P[1]); result = n > head ? null : (n < CHANGE_AT ? "0x01" : "0x02"); break; }
    case "eth_getTransactionCount": { const n = tagNum(P[1]); result = n > head ? null : "0x5"; break; }   // never sees the mempool (stale node)
    case "eth_sendRawTransaction": result = "0x" + "ab".repeat(32); break;
    case "eth_getTransactionByHash": result = { from: SENDER, nonce: "0x5", hash: P[0] }; break;
  }
  r.setHeader("content-type", "application/json"); r.end(JSON.stringify({ jsonrpc: "2.0", id: req.id, result }));
}); });
const rpc = (url, method, params = []) => fetch(url, { method: "POST", headers: { "content-type": "application/json" }, body: JSON.stringify({ jsonrpc: "2.0", id: 1, method, params }) }).then(r => r.json());
const sleep = ms => new Promise(r => setTimeout(r, ms));
(async () => {
  await new Promise(r => fake.listen(19200, "127.0.0.1", r));
  const proxy = spawn("node", [PROXY, "http://127.0.0.1:19200", "19201"], { stdio: ["ignore", "ignore", "inherit"], env: { ...process.env } });
  await sleep(800); const P = "http://127.0.0.1:19201"; let failed = 0;
  const check = (name, ok, detail) => { console.log((ok ? "PASS " : "FAIL ") + name + "  " + detail); if (!ok) failed++; };

  // 1) frozen-floor bug: ask for the block number ONCE, then only poll `latest`
  const bn = parseInt((await rpc(P, "eth_blockNumber")).result, 16);
  const t0 = Date.now(); let value = null, polls = 0;
  while (Date.now() - t0 < 12000) { value = (await rpc(P, "eth_call", [{ to: "0x" + "11".repeat(20), data: "0x" }, "latest"])).result; polls++; if (value === "0x02") break; await sleep(500); }
  check("polling `latest` sees the state change (floor follows the chain)", value === "0x02", `first block number seen=${bn}, change at ${CHANGE_AT}, final value=${value} after ${polls} polls, ${(Date.now() - t0)} ms`);

  // 2) pending nonce after we sent a transaction through the proxy, upstream stays stale at 5
  const before = parseInt((await rpc(P, "eth_getTransactionCount", [SENDER, "pending"])).result, 16);
  await rpc(P, "eth_sendRawTransaction", ["0xdeadbeef"]);
  const after = parseInt((await rpc(P, "eth_getTransactionCount", [SENDER, "pending"])).result, 16);
  check("pending nonce reflects a tx we sent even if the node is stale", before === 5 && after === 6, `before=${before} after=${after} (expect 5 -> 6)`);

  // 3) latest nonce is not inflated by our local record
  const latest = parseInt((await rpc(P, "eth_getTransactionCount", [SENDER, "latest"])).result, 16);
  check("latest nonce unchanged", latest === 5, `latest=${latest} (expect 5)`);

  // 4) explicit block tags pass through untouched (future block -> upstream's null, immediately)
  const s = Date.now(); const fut = (await rpc(P, "eth_call", [{ to: "0x" + "11".repeat(20), data: "0x" }, "0xfffff"])).result;
  check("explicit future block passes straight through", fut === null && Date.now() - s < 1500, `result=${fut} in ${Date.now() - s} ms`);

  // 5) an explicit range at/below the floor that a lagging backend does not have yet is waited for, a real future range is not
  const proxy2 = spawn("node", [PROXY, "http://127.0.0.1:19200", "19202"], { stdio: ["ignore", "ignore", "inherit"], env: { ...process.env, PROXY_TEST_FLOOR_AHEAD: "2" } });
  await sleep(800); const P2 = "http://127.0.0.1:19202";
  const fl = (await rpc(P2, "eth_blockNumber")).result;                         // primes floor = head + 2
  const t1 = Date.now(); const lg = await rpc(P2, "eth_getLogs", [{ fromBlock: fl, toBlock: fl }]);
  check("explicit range <= floor is retried until the backend catches up", Array.isArray(lg.result) && Date.now() - t1 > 500, `result=${JSON.stringify(lg.result || lg.error)} after ${Date.now() - t1} ms`);
  const farBlock = "0x" + (parseInt(fl, 16) + 1000).toString(16); const t2 = Date.now();
  const far = await rpc(P2, "eth_getLogs", [{ fromBlock: farBlock, toBlock: farBlock }]);
  check("explicit range above the floor passes straight through (upstream's error)", far.error && /beyond/.test(far.error.message) && Date.now() - t2 < 1500, `error=${far.error && far.error.message} in ${Date.now() - t2} ms`);
  proxy2.kill();

  console.log(failed ? `\n${failed} FAILED` : "\nall passed"); proxy.kill(); fake.close(); process.exit(failed ? 1 : 0);
})();
