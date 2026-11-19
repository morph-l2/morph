const path = require("path");
const ethers = require(path.join(__dirname, "..", "..", "..", "contracts", "node_modules", "ethers"));
const keccak2 = (a, b) => ethers.utils.keccak256(ethers.utils.concat([a, b]));
// --- copy of sparseProof from withdraw-test.js (extracted by regex so the test runs the real code)
const src = require("fs").readFileSync(require("path").join(__dirname, "withdraw-test.js"), "utf8");
const fn = src.match(/function sparseProof[\s\S]*?\n}\n/)[0];
const sparseProof = new Function("ethers", "keccak2", fn + "; return sparseProof;")(ethers, keccak2);
// --- the contract's algorithm, line by line (Tree.sol)
const ZERO = "0x" + "00".repeat(32);
class Tree { constructor(){ this.branch = Array(32).fill(ZERO); this.count = 0; }
  append(leaf){ let node = leaf; const size = ++this.count; for (let h = 0; h < 32; h++) { if (((size >> h) & 1) === 1) { this.branch[h] = node; return; } node = keccak2(this.branch[h], node); } }
  root(){ let node = ZERO; const size = this.count; let zh = ZERO; for (let h = 0; h < 32; h++) { node = ((size >> h) & 1) === 1 ? keccak2(this.branch[h], node) : keccak2(node, zh); zh = keccak2(zh, zh); } return node; } }
const verify = (leaf, proof, index, root) => { let node = leaf; for (let h = 0; h < 32; h++) node = ((index >> h) & 1) === 1 ? keccak2(proof[h], node) : keccak2(node, proof[h]); return node === root; };
let bad = 0, checks = 0;
const leaves = [];
const t = new Tree();
for (let n = 1; n <= 20; n++) {
  const leaf = ethers.utils.keccak256(ethers.utils.toUtf8Bytes("leaf" + n)); leaves.push(leaf); t.append(leaf);
  const contractRoot = t.root();
  for (let i = 0; i < n; i++) { const { proof, root } = sparseProof(leaves.slice(0, n), i); checks++;
    if (root !== contractRoot || !verify(leaves[i], proof, i, contractRoot)) { bad++; console.log("MISMATCH n=" + n + " i=" + i); } }
}
console.log("checked", checks, "(leaf-count, index) pairs;", bad, "mismatches");
process.exit(bad ? 1 : 0);
