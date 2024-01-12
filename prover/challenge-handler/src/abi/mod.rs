//! ABIs
//!
//! Contract ABIs are refactored into their own module to gracefully deal with allowing missing docs on the abigen macro.
#![allow(missing_docs)]


pub mod rollup_abi {
    use ethers::prelude::abigen;
    abigen!(Rollup, "src/abi/Rollup.json");
}
