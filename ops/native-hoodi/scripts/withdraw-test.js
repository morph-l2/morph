#!/usr/bin/env node
/*
 * withdraw-test.js — L2 -> L1 ETH withdrawal, end to end (P7-5).
 *
 *   1. On L2, the deployer calls L2GatewayRouter.withdrawETH(to, amount, gasLimit) with `to` = a brand-new L1 address.
 *   2. From that receipt it reads the L2CrossDomainMessenger SentMessage event (from/to/value/nonce/message) and the
 *      L2ToL1MessagePasser AppendMessage event (leaf index, message hash, tree root).
 *   3. It waits until the L1 Rollup lists that root as a finalized withdrawal root (i.e. the batch holding the withdrawal
 *      was committed and finalized on L1).
 *   4. It rebuilds the depth-32 sparse Merkle tree from ALL AppendMessage events (zero hash 0x00..00, node = keccak(l||r),
 *      empty subtree of height h = z[h], z[h+1] = keccak(z[h]||z[h])), takes the sibling path of the leaf, and checks that it
 *      reproduces the on-chain root. The proof therefore does not depend on any off-chain proof service.
 *   5. On L1 it calls L1CrossDomainMessenger.proveAndRelayMessage(...) and checks that the new address received the ETH.
 *
 * Env: KEY (deployer private key), L1_RPC, L2_RPC, L1_RUN (dir with deployment/hoodidev.json), MORPH_CONTRACTS.
 * Exit 0 only if the L1 recipient's balance rose by exactly the amount.
 */
const fs = require("fs");
const path = require("path");
const ethers = require(path.join(process.env.MORPH_CONTRACTS, "node_modules", "ethers"));
const art = (p) => require(path.join(process.env.MORPH_CONTRACTS, "artifacts/contracts", p)).abi;

const L1_RPC = process.env.L1_RPC || "http://127.0.0.1:18546";
const L2_RPC = process.env.L2_RPC || "http://127.0.0.1:8545";
const L1_RUN = process.env.L1_RUN;
const AMOUNT = ethers.utils.parseEther(process.env.AMOUNT_ETH || "0.005");
const L1_GAS_LIMIT = Number(process.env.WITHDRAW_GAS_LIMIT || 400000);   // gas for the L1 side of the relay
const TIMEOUT_MIN = Number(process.env.FINALIZE_TIMEOUT_MIN || 25);
const L2_ROUTER = "0x5300000000000000000000000000000000000002";
const L2_MESSENGER = "0x5300000000000000000000000000000000000007";
const L2_PASSER = "0x5300000000000000000000000000000000000001";

const say = (...a) => console.log(new Date().toISOString().slice(11, 19), ...a);
const sleep = (ms) => new Promise((r) => setTimeout(r, ms));
const keccak2 = (a, b) => ethers.utils.keccak256(ethers.utils.concat([a, b]));

function sparseProof(leaves, index) {
  const ZERO = "0x" + "00".repeat(32);
  const z = [ZERO];
  for (let h = 0; h < 32; h++) z.push(keccak2(z[h], z[h]));
  let level = leaves.slice();
  let idx = index;
  const proof = [];
  for (let h = 0; h < 32; h++) {
    const sib = idx ^ 1;
    proof.push(sib < level.length ? level[sib] : z[h]);
    const next = [];
    for (let k = 0; k < Math.ceil(level.length / 2); k++) next.push(keccak2(level[2 * k], 2 * k + 1 < level.length ? level[2 * k + 1] : z[h]));
    level = next.length ? next : [keccak2(z[h], z[h])];
    idx >>= 1;
  }
  return { proof, root: level[0] };
}

(async () => {
  const deployment = JSON.parse(fs.readFileSync(path.join(L1_RUN, "deployment/hoodidev.json"), "utf8"));
  const addr = (n) => deployment.find((x) => x.name === n).address;
  const l1 = new ethers.providers.JsonRpcProvider(L1_RPC);
  const l2 = new ethers.providers.JsonRpcProvider(L2_RPC);
  const wallet2 = new ethers.Wallet(process.env.KEY, l2);
  const wallet1 = new ethers.Wallet(process.env.KEY, l1);

  const router = new ethers.Contract(L2_ROUTER, art("l2/gateways/L2GatewayRouter.sol/L2GatewayRouter.json"), wallet2);
  const messengerIface = new ethers.utils.Interface(art("l2/L2CrossDomainMessenger.sol/L2CrossDomainMessenger.json"));
  const passerIface = new ethers.utils.Interface(art("l2/system/L2ToL1MessagePasser.sol/L2ToL1MessagePasser.json"));
  const rollup = new ethers.Contract(addr("Proxy__Rollup"), art("l1/rollup/Rollup.sol/Rollup.json"), l1);
  const l1Messenger = new ethers.Contract(addr("Proxy__L1CrossDomainMessenger"), art("l1/L1CrossDomainMessenger.sol/L1CrossDomainMessenger.json"), wallet1);

  const recipient = ethers.Wallet.createRandom().address;   // address only: nobody holds its key
  say("L1 recipient", recipient, "(balance", ethers.utils.formatEther(await l1.getBalance(recipient)), "ETH)");

  // 1) withdraw on L2
  const tx = await router["withdrawETH(address,uint256,uint256)"](recipient, AMOUNT, L1_GAS_LIMIT, { value: AMOUNT });
  say("L2 withdraw tx", tx.hash);
  const rc = await tx.wait();
  say("L2 block", rc.blockNumber, "status", rc.status, "gasUsed", rc.gasUsed.toString());
  if (rc.status !== 1) throw new Error("the L2 withdrawal transaction reverted");

  // 2) events
  const sent = rc.logs.filter((l) => l.address.toLowerCase() === L2_MESSENGER.toLowerCase())
    .map((l) => { try { return messengerIface.parseLog(l); } catch (_) { return null; } }).find((e) => e && e.name === "SentMessage");
  const appended = rc.logs.filter((l) => l.address.toLowerCase() === L2_PASSER.toLowerCase())
    .map((l) => { try { return passerIface.parseLog(l); } catch (_) { return null; } }).find((e) => e && e.name === "AppendMessage");
  if (!sent || !appended) throw new Error("SentMessage / AppendMessage not found in the receipt");
  const index = appended.args[0].toNumber();
  say("SentMessage: from", sent.args.sender, "to", sent.args.target, "value", ethers.utils.formatEther(sent.args.value), "nonce", sent.args.messageNonce.toString());
  say("AppendMessage: leaf index", index, "hash", appended.args[1].slice(0, 18) + "…", "root", appended.args[2].slice(0, 18) + "…");

  // 3) wait until a finalized batch lists a withdrawal root that contains our leaf
  //    (the root grows with every later withdrawal, so try the roots of the leaf counts index+1 .. current)
  const allLogs = await l2.getLogs({ address: L2_PASSER, fromBlock: 0, toBlock: "latest", topics: [passerIface.getEventTopic("AppendMessage")] });
  const parsed = allLogs.map((l) => passerIface.parseLog(l).args);
  const leaves = parsed.map((a) => a[1]);
  say("leaves in the L2 withdraw tree:", leaves.length);
  const t0 = Date.now();
  let usable = null;
  while (!usable) {
    for (let n = index + 1; n <= parsed.length; n++) {
      const root = parsed[n - 1][2];                                       // root after leaf n-1 was appended
      if (await rollup.withdrawalRoots(root)) { usable = { n, root }; }    // keep the newest finalized one
    }
    if (usable) break;
    if (Date.now() - t0 > TIMEOUT_MIN * 60000) throw new Error("timed out waiting for the batch with the withdrawal to be finalized on L1");
    say("waiting for the batch to be finalized on L1 (Rollup.lastFinalizedBatchIndex =", (await rollup.lastFinalizedBatchIndex()).toString() + ")");
    await sleep(30000);
  }
  say("finalized withdrawal root found: covers", usable.n, "leaves,", usable.root.slice(0, 18) + "…", "after", Math.round((Date.now() - t0) / 1000), "s");

  // 4) local proof, checked against the root
  const { proof, root } = sparseProof(leaves.slice(0, usable.n), index);
  if (root.toLowerCase() !== usable.root.toLowerCase()) throw new Error(`locally rebuilt root ${root} != on-chain root ${usable.root}`);
  say("proof rebuilt locally; root matches the on-chain root");

  // 5) prove and relay on L1
  const before = await l1.getBalance(recipient);
  const relay = await l1Messenger.proveAndRelayMessage(sent.args.sender, sent.args.target, sent.args.value, sent.args.messageNonce, sent.args.message, proof, usable.root);
  say("L1 proveAndRelayMessage tx", relay.hash);
  const rr = await relay.wait();
  say("L1 block", rr.blockNumber, "status", rr.status, "gasUsed", rr.gasUsed.toString());
  const events = rr.logs.map((l) => { try { return l1Messenger.interface.parseLog(l); } catch (_) { return null; } }).filter(Boolean).map((e) => e.name);
  say("L1 messenger events:", events.join(", ") || "(none)");
  await sleep(15000);
  const after = await l1.getBalance(recipient);
  say("L1 recipient balance:", ethers.utils.formatEther(before), "->", ethers.utils.formatEther(after), "ETH");
  const ok = after.sub(before).eq(AMOUNT) && events.includes("RelayedMessage");
  say(ok ? `OK: ${ethers.utils.formatEther(AMOUNT)} ETH withdrawn from L2 and released on L1 to ${recipient}` : "FAIL: the L1 recipient did not receive exactly the withdrawn amount");
  process.exit(ok ? 0 : 1);
})().catch((e) => { console.error("withdraw-test failed:", e.message || e); process.exit(1); });
