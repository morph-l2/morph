# Staking

## Stake ETH

- Staking amount is immutable

1. `L1` Add to whitelist
2. `L1` Register and stake ETH
3. `L1 -> L2` Bridge Staking message

## Withdraw

- The exit lock period should be long enough to ensure that stakers and sequencers in L2 have been updated
- TBD: When the removal causes the number of stakers to be 0, it will not be successful because L2 Unable to produce new block

1. `L1` Apply for Withdrawal, enter lock peroid
2. `L1 -> L2` Bridge withdrawal message
3. `L2` Update stakers sequencer set
4. `L1` Reach unlock height and staker not in sequencer set，claim allowed，remove staker after claiming

## Delegate stake Morph & sequencer selection

- The number of Sequencers is fixed at 1-X. The number of Staker in L1 should be greater than or equal to X before L2 is started
- When L2 is fisrt started, staker should be initialized to be consistent with L1, and the first X ones will be sequencers
- All users can delegate stake Morph token to all Staker in L2
- Staker sorts according to the staked Morph token amount, selects the first X as sequencers
- A lock-in period is required for users to undelegate or redelegate
- If staker was removed, all delegation staking will be released and need claim manually
- TBD: All stakers were removed

## Rollup

- Before BLS implementation, only staker (not withdrawing, not slashing) can rollup. After the implementation, everyone can rollup
- TBD: When the slash causes the number of stakers or sequencer to be 0, it will not be successful because L2 Unable to produce new block

1. Rollup contract request verifying BLS signature to staking contract. Parameters:
   - sign(batch_hash + sequencer_set_change)
   - signed_sequencers_addresses
   - sequencer_set_addresses
2. Staking contract verifies signature and update latest_sequencer_set
   - If latest_sequencer_set is null (first rollup), sequencer_set_change should be null

## Record

- Block record: Oracle
  - sequencer: `array(block_number, block_proposer, (block_size ? tx_number ?))`
- Delegate staking revord: Contract
  - staker: `array(block, amount)`,
  - delegate_staker: `array(staker, block, amount)`
- Rollup record: Oracle
  - submitter: `array(submitter, BatchInfo(index, startBlock, endBlock, rollupTime, size))`

## Rollup reward

- Reward source: gasfee
- Reward calculation: Ergodic the blocks to claim, statistics on the number of Rollups and data volume, minus the timeout penalty
- Distribute: distribute to submitter by EOA account
- Penalty for rollup timeout: timeout times reaches the threshold

## Staking reward

- Reward source: Morph inflation, claim by `Distribute` contract
- Claim: Manual claim
- Reward calculation: Calculate rewards according to the proportion of blocks produced, record the total reward amount of each block, and record the height of each claim
  - Sequencer: commission (+ block reward?)
  - Delegate Staker: Proportion minus commission

## Penalty for defense challenge fails

- Sequencer will be confiscated all stake amount and removed from sequencer set if defense challenge fails
- Even if challenged repeatedly, each sequencer will only be slash once
- TBD: The reward for a successful challenge is a fixed proportion of the staking amount, and is greater than the challenge deposit

1. `L1` Challenger start a challenge
2. `L1` Challenger win, confiscate all staking amount of sequencers
3. `L1 -> L2` Bridge stakers removed message
4. `L2` Update stakers and sequencer set
5. `L1` Confirm stakers were removed on L2, challenger claim reward and remove stakers on L1

## Penalty for not producing blocks

- If a Sequencer does not produce blocks for a long time, other Sequencers can kick it out
