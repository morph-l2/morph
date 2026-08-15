# Staking

---

## Core Process

**Staking & Sequencer Selection**

Morph token staking will be divided into three stages:

- MorphToken has not yet issued yet
- MorphToken has been issued, but inflation and staking rewards have not started yet.
- Inflation starts and distribute staking rewards

1. `L1` Update whitelist
2. `L1` Register and stake eth to become staker
3. `L1->L2` Send {add staker} message
4. `L2` Update stakers
5. `L2` Delegate stake MorphToken to staker
6. `L2` Update sequencer set by amount of MorphToken delegation

**Rollp Verify**

- The submitted batch requires the BLS signature of more than 2/3 of the sequencers
- Before BLS implementation, only staker (not withdrawing, not slashing) can rollup

1. `L1` Submit batch and sequencer sets
2. `L1` Verify sequencer set
3. `L1` Verify batch signature

**Slash Sequencers After Challenger Win**

- Sequencer will be confiscated all stake amount and removed from sequencer set if defense challenge fails
- Even if challenged repeatedly, each sequencer will only be slashed once
- The reward for a successful challenge is a fixed proportion of the staking amount, and is greater than the challenge deposit
- If the slash causes the number of sequences to be 0, then the layer2 will stop running. We can restart by upgrading, reset stakers and sequencer set. This does not affect the security of the evm state

1. `L1` Slash staking value of signed sequencers and remove from stakers
2. `L1` Distribute challenger rewards
3. `L1->L2` Send {remove stakers} message
4. `L2` Update sequencer set by amount of MorphToken delegation

**Delegation Rewards**

- MorphToken is inflated at a fixed ratio every day, and the additional tokens are used as rewards for sequencers and delegators
- The sequencer and his delegators get MorphToken inflation reward according to the block production ratio
- Sequencer charges commission as its own income
- Users get remaining reward according to their delegation amount

1. `L2` Staker set delegation commission rate
2. `L2` Upload sequencers work records in epoch (an epoch is a day)
3. `L2` Mint MorphToken inflation as delegation reward
4. `L2` Claim delegation reward or claim commission

**Quit Staker**

The exit lock period should be long enough to ensure that stakers and sequencers in L2 have been updated

1. `L1` Apply to withdraw, remove from staker list, enter lock period
2. `L1->L2` Send {remove staker} message
3. `L2` Remove from stakers and sequencers if needed
4. `L1` Claim allowed until reach unlock height，remove staker info after claiming

---

## Contracts

### L1 Contract

#### L1Staking Contract

**Main functions**

- Update whitelist
- Register as staker
- Staker quit
- Slash signed sequencers if challenger win, call by rollup contract

#### Rollup Contract

**Main functions**

- Commit batch & verify
- Challenge batch

### L2 Contract

#### MorphToken Contract

**Main functions**

- Mint inflation

#### Record Contract

**Main functions**

- Record finalized batch submissions, call by oracle
- Record rollup epochs, call by oracle
- Record reward epochs, call by oracle

#### L2Staking Contract

**Main functions**

- Add staker, sync from L1 staking contract
- Remove stakers, sync from L1 staking contract
- Staker set commission rate
- User delegate stake morph token
- User undelegate stake morph token
- User claim delegation reward
- Sequencer claim commission

#### Sequencer Contract

**Main functions**

- Update sequencer set, call form l2 staking contract

#### Distribute Contract

**Main functions**

- Record delegation info by l2 staking contract
- Record undelegation info by l2 staking contract
- Compute delegator's delegation reward, call form l2 staking contract
- Compute sequencer's commission, call form l2 staking contract

#### Gov Contract

**Main functions**

- Submit proposal
- Vote proposal
