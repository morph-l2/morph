use ethers::prelude::*;
use ethers::providers::{Http, Provider};
use ethers::signers::Wallet;
use ethers::types::Address;
use std::thread;
use std::time::Duration;
use std::{error::Error, str::FromStr, sync::Arc};

const CONTRACT_ADDRESS: &str = "0x";
const PRIVATE_KEY: &str = "0x";

abigen!(TestZkEVM, "./tests/abi/TestZkEVM.json");

// Generate tx for traces.
#[tokio::test]
async fn test() -> Result<(), Box<dyn Error>> {
    let result = call().await;
    match result {
        Ok(()) => Ok(()),
        Err(e) => {
            println!("call error:");
            Err(e)
        }
    }
}

// Call a contract.
pub async fn call() -> Result<(), Box<dyn Error>> {
    let provider: Provider<Http> = Provider::<Http>::try_from("http://127.0.0.1:8569")?;
    let wallet: LocalWallet = Wallet::from_str(PRIVATE_KEY)?;

    let signer = Arc::new(SignerMiddleware::new(
        provider.clone(),
        wallet.with_chain_id(1337_u64),
    ));

    let contract: TestZkEVM<SignerMiddleware<Provider<Http>, _>> =
        TestZkEVM::new(Address::from_str(CONTRACT_ADDRESS)?, signer);

    let tx = contract.transfer(Address::from_str("0x").unwrap(), 10.into()).legacy();
    let pending_tx = tx.send().await.unwrap();

    let block_num = get_tx_blocknum(provider.clone(), pending_tx.tx_hash())
        .await
        .expect("get tx blocknum err");
    println!("block_num: {:?}", block_num);
    Ok(())
}

// Deploy a contract.
pub async fn deploy() -> Result<(), Box<dyn Error>> {
    let provider: Provider<Http> = Provider::<Http>::try_from("http://127.0.0.1:8569")?;
    let wallet: LocalWallet = Wallet::from_str(PRIVATE_KEY)?;

    let signer = Arc::new(SignerMiddleware::new(
        provider.clone(),
        wallet.with_chain_id(1337_u64),
    ));
    // let factory = ContractFactory::new(abi, bytecode, client.clone());

    let initial_supply: u64 = 10;

    let tx = TestZkEVM::deploy(signer, initial_supply.pow(18))?.legacy();
    let contract = tx.send().await;

    match contract {
        Ok(sent_tx) => println!("====testZkEVM: {:?}", sent_tx),
        Err(e) => println!("deploy exception: {:?}", e),
    }

    Ok(())
}

async fn get_tx_blocknum(provider: Provider<Http>, tx_hash: H256) -> Option<U64> {
    println!("tx_hash: {:?}", tx_hash);
    for _ in 0..10 {
        thread::sleep(Duration::from_millis(2000));
        let receipt = provider.get_transaction_receipt(tx_hash).await.unwrap();
        match receipt {
            Some(receipt) => {
                return match receipt.status.unwrap().as_u64() {
                    1 => receipt.block_number,
                    _ => None,
                };
            }
            // Maybe still pending
            None => continue,
        }
    }
    println!("nothing");
    None
}
