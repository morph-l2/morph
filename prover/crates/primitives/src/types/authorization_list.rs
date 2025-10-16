use alloy::{
    eips::eip7702::SignedAuthorization,
    primitives::{Address, Signature, U256, U8},
};
use serde::{Deserialize, Deserializer, Serialize, Serializer};

/// A wrapper around SignedAuthorization that implements Archive trait
#[derive(
    Clone, Debug, Default, PartialEq, Eq, Hash, rkyv::Archive, rkyv::Serialize, rkyv::Deserialize,
)]
#[archive(check_bytes)]
#[archive_attr(derive(Debug, PartialEq, Eq, Hash))]
pub struct ArchivedSignedAuthorization {
    /// The chain ID of the authorization
    pub chain_id: U256,
    /// The address of the authorization
    pub address: Address,
    /// The nonce for the authorization
    pub nonce: u64,
    /// Signature r value
    pub r: U256,
    /// Signature s value
    pub s: U256,
    /// Signature v value
    pub v: U8,
}

// Custom serialization for JSON compatibility
impl Serialize for ArchivedSignedAuthorization {
    fn serialize<S>(&self, serializer: S) -> Result<S::Ok, S::Error>
    where
        S: Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut state = serializer.serialize_struct("ArchivedSignedAuthorization", 6)?;
        state.serialize_field("chainId", &self.chain_id)?;
        state.serialize_field("address", &self.address)?;
        state.serialize_field("nonce", &self.nonce)?;
        state.serialize_field("r", &self.r)?;
        state.serialize_field("s", &self.s)?;
        state.serialize_field("yParity", &self.v)?;
        state.end()
    }
}

// Custom deserialization for JSON compatibility
impl<'de> Deserialize<'de> for ArchivedSignedAuthorization {
    fn deserialize<D>(deserializer: D) -> Result<Self, D::Error>
    where
        D: Deserializer<'de>,
    {
        #[derive(Deserialize)]
        struct ArchivedSignedAuthorizationHelper {
            #[serde(rename = "chainId")]
            chain_id: U256,
            address: Address,
            nonce: u64,
            r: U256,
            s: U256,
            #[serde(rename = "yParity")]
            y_parity: U8,
        }

        let helper = ArchivedSignedAuthorizationHelper::deserialize(deserializer)?;
        Ok(ArchivedSignedAuthorization {
            chain_id: helper.chain_id,
            address: helper.address,
            nonce: helper.nonce,
            r: helper.r,
            s: helper.s,
            v: helper.y_parity,
        })
    }
}

impl From<SignedAuthorization> for ArchivedSignedAuthorization {
    fn from(auth: SignedAuthorization) -> Self {
        let (inner, signature) = auth.into_parts();
        Self {
            chain_id: inner.chain_id,
            address: inner.address,
            nonce: inner.nonce,
            r: signature.r(),
            s: signature.s(),
            v: U8::from(signature.v().y_parity_byte()),
        }
    }
}

impl From<ArchivedSignedAuthorization> for SignedAuthorization {
    fn from(auth: ArchivedSignedAuthorization) -> Self {
        use alloy::eips::eip7702::Authorization;
        let inner =
            Authorization { chain_id: auth.chain_id, address: auth.address, nonce: auth.nonce };
        let v: u8 = auth.v.to();
        let parity = alloy::primitives::Parity::try_from(v as u64).unwrap();
        let signature = Signature::from_rs_and_parity(auth.r, auth.s, parity).unwrap();
        inner.into_signed(signature)
    }
}

/// A wrapper around Vec<SignedAuthorization> that implements Archive trait
#[derive(
    Clone,
    Debug,
    Default,
    PartialEq,
    Eq,
    Hash,
    rkyv::Archive,
    rkyv::Serialize,
    rkyv::Deserialize,
    Serialize,
    Deserialize,
)]
#[archive(check_bytes)]
#[archive_attr(derive(Debug, PartialEq, Eq, Hash))]
pub struct AuthorizationList(pub Vec<ArchivedSignedAuthorization>);

impl From<Vec<SignedAuthorization>> for AuthorizationList {
    fn from(auths: Vec<SignedAuthorization>) -> Self {
        Self(auths.into_iter().map(ArchivedSignedAuthorization::from).collect())
    }
}

impl From<AuthorizationList> for Vec<SignedAuthorization> {
    fn from(auths: AuthorizationList) -> Self {
        auths.0.into_iter().map(SignedAuthorization::from).collect()
    }
}
