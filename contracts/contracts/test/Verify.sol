// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {stdJson} from "forge-std/StdJson.sol";

import {BatchHeaderCodecV0} from "../libraries/codec/BatchHeaderCodecV0.sol";
import {BatchHeaderCodecV1} from "../libraries/codec/BatchHeaderCodecV1.sol";

struct ProofFixture {
    bytes proof;
    bytes publicValues;
    bytes32 vkey;
}

contract VerifyTest is Test {
    using stdJson for string;

    uint64 public constant LAYER_2_CHAIN_ID = 1;

    function setUp() public {}

    // 保持这个函数接受 memory，但内部处理
    function proveState(bytes memory _batchHeader, bytes memory _batchProof) public returns (bytes32) {
        (uint256 memPtr, bytes32 _batchHash) = _loadBatchHeaderFromMemory(_batchHeader);
        uint256 _batchIndex = BatchHeaderCodecV0.getBatchIndex(memPtr);
        bytes32 _blobVersionedHash = BatchHeaderCodecV0.getBlobVersionedHash(memPtr);

        bytes32 _publicInputHash = keccak256(
            abi.encodePacked(
                LAYER_2_CHAIN_ID,
                BatchHeaderCodecV0.getPrevStateHash(memPtr),
                BatchHeaderCodecV0.getPostStateHash(memPtr),
                BatchHeaderCodecV0.getWithdrawRootHash(memPtr),
                BatchHeaderCodecV0.getSequencerSetVerifyHash(memPtr),
                BatchHeaderCodecV0.getDataHash(memPtr),
                _blobVersionedHash
            )
        );

        return _publicInputHash;
    }

    // 新的辅助函数：从 memory 加载 batch header
    function _loadBatchHeaderFromMemory(bytes memory _batchHeader) 
        internal 
        pure 
        returns (uint256 _memPtr, bytes32 _batchHash) 
    {
        uint8 _version = _getBatchVersionFromMemory(_batchHeader);

        // 直接在内存中处理，不使用 calldata 版本
        uint256 _length = _batchHeader.length;
        
        if (_version == 0) {
            require(_length >= 249, "batch header length too small");
        } else if (_version == 1) {
            require(_length == 257, "batch header length is incorrect");
        } else {
            revert("Unsupported batch version");
        }

        // 获取内存指针 - batch header 已经在内存中
        assembly {
            _memPtr := add(_batchHeader, 0x20) // 跳过长度字段
        }

        // compute batch hash
        _batchHash = BatchHeaderCodecV0.computeBatchHash(_memPtr, _length);
    }

    function _getBatchVersionFromMemory(bytes memory batchHeader) internal pure returns (uint8 version) {
        require(batchHeader.length > 0, "Empty batch header");
        version = uint8(batchHeader[0]);
    }

    // 如果你想要一个接受 calldata 的版本（用于外部调用）
    function _loadBatchHeader(bytes calldata _batchHeader) 
        internal 
        pure 
        returns (uint256 _memPtr, bytes32 _batchHash) 
    {
        uint8 _version = _getBatchVersion(_batchHeader);

        // load to memory
        uint256 _length;
        if (_version == 0) {
            (_memPtr, _length) = BatchHeaderCodecV0.loadAndValidate(_batchHeader);
        } else if (_version == 1) {
            (_memPtr, _length) = BatchHeaderCodecV1.loadAndValidate(_batchHeader);
        } else {
            revert("Unsupported batch version");
        }

        // compute batch hash
        _batchHash = BatchHeaderCodecV0.computeBatchHash(_memPtr, _length);
    }

    function _getBatchVersion(bytes calldata batchHeader) internal pure returns (uint8 version) {
        require(batchHeader.length > 0, "Empty batch header");
        version = uint8(batchHeader[0]);
    }

    // Prove state success.
    function test_ProveStateSuccess() public {
        bytes
            memory _batchHeader = hex"01000000000000b33b0000000000000000000000000000635c00bd1546a9557a270495c743a00dfbdf4c3b57873740eb036aae71d13041b6cf0124968ca5cf42b2db2c57542568ddc2dff098a9ffe0e120df4b24322e9985b00463265e4b6dd008936a4a0adc38c5a7e786b7de27a4df951e9115e754cf2f881000d6108f7b74a96d688314104222624a0ab21ebef383c944fa7833bf2ec0aa885d6ea9187fb9d6139f767631ac0aad4cf94d73af813bd069b4c7e45a9867cae687203afe8c37fdaf5f2b4d01ac53992d0af040e2b5913eb37b0dcf3e927046cd06ce2b9123de9ede20c4d5a98c86f559cffc32d30d13471fdded585a7c305100000000011c0d73";
        bytes
            memory _batchProof = hex"ffea2d2e151c9300762295d9b0dde28c97bda5ed68afe764b65c26fcc0559cf128c88ba529016be93f3dded791b04e85d0c8af65224175287fba2bdff17f4232065c29371642f0e12779ea298059a02dbdcf21e4428335acfa692d30ad0e0a9364de6eb120f48ad2eac1256e86f2e5a3e581122ec276ae1f1482697cd9995745e0ba84371dfa96f820899a465b0cd61623bf5465c5b1c0d772f1597c82521dbcf24a97f21aad3f1da890e9ea7cd004664067dffb062a6c0dad465494888228696c41876826968671c145b89dc7878885ce62f8979bce9bd7b44480f0ab1a3e20539fdaf118e13deabcc4c7786fefe0a197331f70f061fede5e98301eaca76e16224b1be002f900636da075c3ed17894d25099f2072ba6efce7028e78ae3097a25bd12c5113c29b1b4fc9cc6caf138095ddfb7ec21493633b28b404d50e19a1642eae63162ea4d5385373d431dd81646eb3a34e03a5641dafb145a813a51bae0ad010933f0b110d9d6f9e337a0f6cee71b97877164ed2437de7bf7f3b0db7a74c406f06d20bb82c2ea6aab0e876f0041c11b885b7dd14b4e1307ab7da5bb63fa52ca3d4760aa9d4899c50b59ac8f495d46acba97f3f03b59e09c0167ad98da039be7076420a9c799dee7c3eb0e3696787928c79df4df83a8920dbaa6c9fa856b9aebdfee9016a668bb8e258b63dffc14d4204d9d0b174dc519988445231c494bca0e349292edef25cc0ed9c76aed0e6def99fee173c3bfd2742d0566e3d1970171d35f0bb20c7adf0a5b596b9c19b5a16815e2628a3c099a2b56a9edbc546692a23ba09b2154f581deadb812bd84f6f6895521eef851c586dc7a34cf66ec4939f4fd139dd02d4620d51a15d0fd42c74de4a5ec59a31d340318811f62a91adb0af6946a8ba1e50e20f520536078b45173f150ac68344dcf474d929ef106356e6b6b32a4a4b10479ab02d01ff40d21008307ef5cdd853ee164862418e7d0e3640a01825e2c723abdb6712cf01116158233be29e4af3d3f21fa019be93313a43350a0b8c61b6209f211ff7cee5bf5d6307fef7e47c3a211730d83cff1b00918fe9515464f1ae010ea81bf21d0b9ddcc345bcc384bf2d9f5b6f61ebd0b97249e038572fcac8150d509cd3671f7736a7dd8fcfd4dab9adfe6118881ca2bb5a2441d492c050763524e470ce2dd2dec4bef2979857d0fb9def8a5e921a37a10165ad577badbf6de1";
        bytes32 _publicInputHash = proveState(_batchHeader, _batchProof);
        console.logBytes32(_publicInputHash);
    }
}

// cast call verifyBatch
// cast call 0x045d4BC73Bd1918192f34e98532A5272Ef620423 \
//   "verifyBatch(bytes,bytes32)" \
//   0xffea2d2e151c9300762295d9b0dde28c97bda5ed68afe764b65c26fcc0559cf128c88ba529016be93f3dded791b04e85d0c8af65224175287fba2bdff17f4232065c29371642f0e12779ea298059a02dbdcf21e4428335acfa692d30ad0e0a9364de6eb120f48ad2eac1256e86f2e5a3e581122ec276ae1f1482697cd9995745e0ba84371dfa96f820899a465b0cd61623bf5465c5b1c0d772f1597c82521dbcf24a97f21aad3f1da890e9ea7cd004664067dffb062a6c0dad465494888228696c41876826968671c145b89dc7878885ce62f8979bce9bd7b44480f0ab1a3e20539fdaf118e13deabcc4c7786fefe0a197331f70f061fede5e98301eaca76e16224b1be002f900636da075c3ed17894d25099f2072ba6efce7028e78ae3097a25bd12c5113c29b1b4fc9cc6caf138095ddfb7ec21493633b28b404d50e19a1642eae63162ea4d5385373d431dd81646eb3a34e03a5641dafb145a813a51bae0ad010933f0b110d9d6f9e337a0f6cee71b97877164ed2437de7bf7f3b0db7a74c406f06d20bb82c2ea6aab0e876f0041c11b885b7dd14b4e1307ab7da5bb63fa52ca3d4760aa9d4899c50b59ac8f495d46acba97f3f03b59e09c0167ad98da039be7076420a9c799dee7c3eb0e3696787928c79df4df83a8920dbaa6c9fa856b9aebdfee9016a668bb8e258b63dffc14d4204d9d0b174dc519988445231c494bca0e349292edef25cc0ed9c76aed0e6def99fee173c3bfd2742d0566e3d1970171d35f0bb20c7adf0a5b596b9c19b5a16815e2628a3c099a2b56a9edbc546692a23ba09b2154f581deadb812bd84f6f6895521eef851c586dc7a34cf66ec4939f4fd139dd02d4620d51a15d0fd42c74de4a5ec59a31d340318811f62a91adb0af6946a8ba1e50e20f520536078b45173f150ac68344dcf474d929ef106356e6b6b32a4a4b10479ab02d01ff40d21008307ef5cdd853ee164862418e7d0e3640a01825e2c723abdb6712cf01116158233be29e4af3d3f21fa019be93313a43350a0b8c61b6209f211ff7cee5bf5d6307fef7e47c3a211730d83cff1b00918fe9515464f1ae010ea81bf21d0b9ddcc345bcc384bf2d9f5b6f61ebd0b97249e038572fcac8150d509cd3671f7736a7dd8fcfd4dab9adfe6118881ca2bb5a2441d492c050763524e470ce2dd2dec4bef2979857d0fb9def8a5e921a37a10165ad577badbf6de1 \
//   0xa70157e10d8ef5a038bf7ed5c7cdac15b8b4e688327e71309398a32852d84217 \
//   --rpc-url https://cloudflare-eth.com