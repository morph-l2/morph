# Lido's Morph Bridge

The document outlines the process of bridging ERC20-compatible tokens between the Ethereum and Morph chains.

This marks the initial phase of Lido's integration into the Morph protocol. The primary objective of this implementation is to establish a solid foundation for Lido's long-term expansion goals on the Morph chain. The broader vision for Lido's integration into Layer 2 solutions includes:

- Bridging Lido's tokens from Layer 1 to Layer 2 chains
- Enabling instant ETH staking on Layer 2 chains, with users receiving stETH/wstETH immediately on the corresponding Layer 2
- Maintaining a user experience on Layer 2 that closely resembles that of the Ethereum mainnet
At this stage, the implementation aims to deliver a scalable and reliable solution for Lido to bridge ERC20-compatible tokens between Morph and the Ethereum chain.

## Security surface overview

| Statement                                                                                                                                    | Answer                                                                                                                                                                                                                                                                             |
| -------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| It is possible to bridge wstETH forth and back using this bridge                                                                             | Yes                                                                                                                                                                                                                                                                                |
| The bridge using a canonical mechanism for message/value passing                                                                             | Yes                                                                                                                                                                                                                                                                                |
| The bridge is upgradeable                                                                                                                    | Yes                                                                                                                                                                                                                                                                                |
| Upgrade authority for the bridge                                                                                                             | Yes                                                                                                                                                                                                                                                                                |
| Emergency pause/cancel mechanisms and their authorities                                                                                      | TBA                                                                                                                                                                                                                                                                                |
| The bridged token support permits and ERC-1271                                                                                               | Yes                                                                                                                                                                                                                                                                                |
| Are the following things in the scope of this bridge deployment:                                                                             |                                                                                                                                                                                                                                                                                    |
| - Passing the (w)stETH/USD price feed                                                                                                        | No                                                                                                                                                                                                                                                                                 |
| - Passing Lido DAO governance decisions                                                                                                      | [Lido DAO Agent](https://etherscan.io/address/0x3e40D73EB977Dc6a537aF587D48316feE66E9C8c) representation via [MorphBridgeExecutor](TBD) |
| Bridges are complicated in that the transaction can succeed on one side and fail on the other. What's the handling mechanism for this issue? | TBA                                                                                                                                                                                                                                                                                |
| Is there a deployment script that sets all the parameters and authorities correctly?                                                         | No, We use hardhat tasks to deploy and perform related permission operations after the test is completed.                                                                                                                                                                             |
| Is there a post-deploy check script that, given a deployment, checks that all parameters and authorities are set correctly?                  | No                                                                                                                                                                                                                                                                                 |

## Morph's Bridging Flow

The default Morph bridging solution consists of two components: `L1StandardERC20Gateway` and `L2StandardERC20Gateway`. These contracts facilitate the bridging of ERC20 tokens between the Ethereum and Morph chains.

In the standard bridge, when an ERC20 token is deposited on Layer 1 and sent to the bridge contract, it is "locked" there while an equivalent amount of the L2 token is minted. For withdrawals, the process is reversed: the L2 token amount is burned, and the same amount of L1 tokens is sent to the recipient.

While the default Morph bridge meets Lido's short-term goal of bridging the wstETH token into Morph, it poses challenges for achieving long-term objectives. For instance, implementing staking from Layer 2 will likely require modifications to both the token and gateway implementations.

Moreover, Morph offers the flexibility to create a custom bridge solution using the same cross-domain infrastructure as the standard bridge. The only requirement for a custom bridge to be compatible with the default Morph Gateway is the implementation of the `IL1ERC20Gateway` and `IL2ERC20Gateway` interfaces.

The remainder of the document includes technical specifications for the bridge that Lido will use to transfer tokens between the Ethereum and Morph chains.


## Lido's Bridge Implementation

The current implementation of the token bridge enables the transfer of ERC20-compatible tokens between Ethereum and Morph chains. It also includes administrative features, such as the ability to temporarily disable deposits and withdrawals. This functionality is crucial for quickly addressing potential malicious activities or vulnerabilities within the contracts, as well as facilitating the upgrade process.

The technical implementation prioritizes the following requirements for the contracts:

- **Scalability**: The current design must allow for future extensions and new functionalities.
- **Simplicity**: The contracts should be clear, straightforward, and easy for future developers to understand and work with.
- **Gas Efficiency**: The solution should minimize gas costs for users while maintaining clarity and simplicity.

A high-level overview of the proposed solution can be found in the diagram below:


![](https://i.imgur.com/rbIvCr6.png)

- [**`LidoGatewayManager`**](./LidoGatewayManager.sol): This contract includes administrative methods to manage and monitor the state of the bridging process.
- [**`LidoBridgeableTokens`**](./LidoBridgeableTokens.sol) : This contract implements the logic for validating tokens involved in the bridging process.
- [**`L1LidoGateway`**](./L1LidoGateway.sol): This contract serves as Ethereum's counterpart for bridging registered ERC20-compatible tokens between Ethereum and Morph chains.
- [**`L2LidoGateway`**](./L2LidoGateway.sol): This contract acts as Morph's counterpart for bridging registered ERC20-compatible tokens between Ethereum and Morph chains.
- [**`MorphStandardERC20`**](../libraries/token/MorphStandardERC20.sol): This is an implementation of the `ERC20` token, equipped with administrative methods for minting and burning tokens.
- [**`TransparentUpgradeableProxy`**](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) : This is an ERC1967 proxy that includes additional administrative functionalities.

## Morph's Bridging Flow

You can find the general process for bridging tokens through Morph's Lido bridge here: [ETH and ERC20 Token Bridge](https://docs.morphl2.io/docs/build-on-morph/build-on-morph/bridge-between-morph-and-ethereum).

The goal of cross-chain governance is bridging the Lido DAO governance decisions, voted by the LDO holders on Ethereum, to Morph network. The main component of this bridge is MorphBridgeExecutor contract on L2, which queues the action sets sent from L1 Lido Agent through L1Executor.

![](https://i.imgur.com/4cgEo63.png)


## Deployment Process

To minimize gas costs for users, the `L1LidoGateway`, `L2LidoGateway`, and `MorphStandardERC20` contracts utilize immutable variables wherever possible. However, some of these variables reference each other; for instance, `L1LidoGateway` refers to `L2LidoGateway` and vice versa. By using proxies, we can initially deploy them without invoking the `initialize` function for each gateway, and then call the `initialize` function with the correct contract addresses.

Alternatively, we could pre-calculate the future addresses of the deployed contracts off-chain and deploy the implementation using these pre-calculated addresses, but this approach is less fault-tolerant than the first option.

## Integration Risks

As an additional component in the token flow chain, the Morph protocol and bridges introduce potential points of failure. The main risks associated with the current integration are outlined below:

### Minting of Uncollateralized L2 Tokens

An attack could occur if an attacker gains access to call `L2LidoGateway.finalizeDepositERC20()` directly. In this scenario, they could mint uncollateralized tokens on L2 and later initiate a withdrawal.

To detect such an attack, off-chain monitoring of minting and deposit/withdrawal events is essential. The following statistics can be tracked based on these events:

- `l1ERC20TokenBridgedAmount`: Total number of tokens bridged on the L1 bridge contract.
- `l2TokenTotalSupply`: Total number of minted L2 tokens.
- `l2TokenNotFinalizedDeposit`: Total number of locked L1 tokens that have not been finalized and relayed from the L2 bridge.
- `l2TokenNotWithdrawn`: Total number of burned L2 tokens that have not been withdrawn from the L1 bridge.

The following invariant must always hold true: `l1ERC20TokenBridgedAmount == l2TokenTotalSupply + l2TokenNotWithdrawn + l2TokenNotFinalizedDeposit`.

If this invariant is violated, Lido will enter a dispute period to suspend both the L1 and L2 bridges. During this time, the bridges will be disabled, preventing the minting of L2 tokens and the withdrawal of minted tokens until the issue is resolved.

### Attack on L1CrossDomainMessenger

According to the Morph documentation, the `L1CrossDomainMessenger` contract sends messages from L1 to L2 and executes the challenged L2 messages in the `Rollup`.

This contract plays a crucial role in L2-to-L1 communication, as all messages from L2 that are finalized by a challenger and verified by Merkle proof are executed on behalf of this contract.

If there is a vulnerability in the `L1CrossDomainMessenger` or `Rollup` that allows an attacker to send arbitrary messages while bypassing the challenge and Merkle proof, they could potentially drain tokens from the L1 bridge.

Additionally, the upgradeability of the `L1CrossDomainMessenger` and `Rollup` poses risks. An attacker could exploit this by replacing the implementation with malicious functionality, which could lead to the theft of all locked tokens on the L1 bridge.

To respond swiftly to such an attack, Lido can implement monitoring of the Proxy contract to raise an alert in the event of an implementation upgrade.

--- 

Feel free to ask if you need further adjustments or additional information!
