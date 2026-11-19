#!/usr/bin/env node
/*
 * send-tx.js — send one transaction and wait for it. The private key comes from the environment (KEY), never from the
 * command line, so it does not show up in `ps`. (This cast version ignores ETH_PRIVATE_KEY for `cast send`.)
 *
 * Usage:  KEY=0x... send-tx.js RPC_URL TO [--value-ether N] [--data 0x...]
 * Needs ethers v5: it is loaded from $MORPH_CONTRACTS/node_modules (the contracts workspace of this repo).
 * Exit code 0 only if the receipt has status 1.
 */
const path = require("path");
const contracts = process.env.MORPH_CONTRACTS;
if (!contracts) { console.error("MORPH_CONTRACTS is not set"); process.exit(2); }
const ethers = require(path.join(contracts, "node_modules", "ethers"));

const [rpc, to, ...rest] = process.argv.slice(2);
if (!rpc || !to || !process.env.KEY) { console.error("usage: KEY=0x... send-tx.js RPC_URL TO [--value-ether N] [--data 0x...]"); process.exit(2); }
const opt = (name) => { const i = rest.indexOf(name); return i >= 0 ? rest[i + 1] : undefined; };

(async () => {
  const provider = new ethers.providers.JsonRpcProvider(rpc);
  const wallet = new ethers.Wallet(process.env.KEY, provider);
  const tx = { to };
  if (opt("--value-ether")) tx.value = ethers.utils.parseEther(opt("--value-ether"));
  if (opt("--data")) tx.data = opt("--data");
  const sent = await wallet.sendTransaction(tx);
  console.log("from", wallet.address, "to", to, "hash", sent.hash);
  const receipt = await sent.wait(1);
  console.log("block", receipt.blockNumber, "status", receipt.status, "gasUsed", receipt.gasUsed.toString());
  process.exit(receipt.status === 1 ? 0 : 1);
})().catch((e) => { console.error("send-tx failed:", e.message || e); process.exit(1); });
