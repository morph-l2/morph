# Staking

## Stake ETH

- Staking 金额固定且不可修改

1. `L1` 加入白名单
2. `L1` 注册并质押 ETH
3. `L1` Restaking
4. `L1 -> L2` Bridge Staking 消息

## Withdraw

- 退出锁定期应足够长，以保证 L2 中 Stakers 和 Sequencer 已完成更新
- TBD: Staker 全部退出后如何处理？

1. `L1` 申请 Withdraw，进入锁定期
2. `L1` UnRestaking
3. `L1 -> L2` Bridge 退出消息
4. `L2` 更新 Stakers 和 SequencerSet
5. `L1` 到达锁定期且 Staker 不在 SequencerSet 中，允许 Claim，Claim 后从 Stakers 中移除

## Delegate Stake Morph & Sequencer Selection

- Sequencer 数量固定为 1-X 个，L2 启动前 L1 中 Staker 数量应大于等于 X 个
- L2 初始启动时应将 Staker 初始化为与 L1 一致，且前 X 个作为 SequencerSet
- 所有用户可向 L2 中所有 Staker 委托质押 Morph Token
- Staker 按照 Morph 质押数量排序，选取前 X 个作为 SequencerSet，合约内实时更新，共识层延迟一块更新完成
- 用户解除委托或转委托需要锁定期 (具体时间 TBD)
- 若 Staker 被移除，则用户委托质押的金额立即释放
- 若 Staker 全部移除，则链停止，重启方案 TBD

## Rollup

- BLS 功能未实现前，只有 Staker（not withdrawing, not slashing）可以进行 Rollup
  - TBD：不经过验签存在风险，某个 Staker 上传 Batch 导致挑战成功，所有签名者全部被 Slash
- BLS 功能实现后，任何人都可以进行 Rollup（须验签通过）
- TBD: Sequencer 使用 Index 还是 Address？优化减小数据
- TBD: 挑战导致 Sequencer 或 Staker 全部退出后如何处理？

1. Rollup 合约请求 Staking 合约验证 BLS 签名。参数：
   - sign(batch_hash + sequencer_set_change)
   - signed_sequencers_addresses
   - sequencer_set_addresses
2. Staking 合约验证签名并更新 latest_sequencer_set
   - latest_sequencer_set 为空（首个 Rollup）时 sequencer_set_change 为空

## Slash Defense Challenge Failed

- Sequencer 挑战失败后罚没质押金额
- 每个 Sequencer 只会被惩罚一次，罚没全部质押金额
- TBD：挑战成功奖励金额 = 所有 Sequencer 被罚没的金额？

1. `L1` 挑战者发起挑战
2. `L1` 挑战者挑战成功，记录奖励信息及解锁区块
3. `L1` UnRestaking 受罚者
4. `L1 -> L2` Bridge 受罚者退出 Stakers 消息
5. `L2` 更新 Stakers 和 SequencerSet
6. `L1` 达到解锁区块后（UnRestaking 完成），挑战者 Claim 奖励，将受罚者移除 Stakers

## Record

- 出块记录: Oracle
  - sequencer: `array(block_number, block_proposer, (block_size ? tx_number ?))`
- Delegate Staking 记录: 合约存储
  - staker: `array(block, amount)`,
  - delegate_staker: `array(staker, block, amount)`
- Rollup 记录: Oracle
  - submitter: `array(submitter, BatchInfo(index, startBlock, endBlock, rollupTime, size))`

## Pay Reward

- TBD: 奖励来源
  - 方案一：自动增发到 Distribute 合约（需要修改电路？）
  - 方案二：外部定期向 Distribute 合约转账

## Sequencer Claim Reward

- Claim 奖励时自动扣除 Slash 金额
- Slash:
  - TBD: Sequencer 超过一定区块数未出块
  - TBD: 作为 Submitter 超时达到一定数量
- 每次 Claim 奖励区间为上次 Claim 高度至当前区块
  - Calculate block reward：遍历 Claim 区间区块，记录每个区块的奖励金额 (sequencer 奖励 & delegate_staker 奖励)，累加 Sequencer 奖励后减去惩罚，计算总奖励金额
  - Calculate rollup reward：遍历 Claim 区间区块，统计 Rollup 数量及数据量，减去超时惩罚，计算总奖励金额
- TBD: 具体金额及分配比例
- TBD: 若奖励来源为外部向 Distribute 合约转账则需确定可以 Claim 的区间

## Delegate Staker Claim Reward

- 遍历 Claim 区间区块，逐块计算委托的 Staker 奖励金额并记录 (Sequencer 和所有 delegate_staker 只需计算一次)，按比例计算所得份额，累加后获得总奖励
- TBD: 具体金额及分配比例

---

TBD: Sequencer 长期不出块是否踢出 SequencerSet, 若踢出则自动触发还是手动触发
