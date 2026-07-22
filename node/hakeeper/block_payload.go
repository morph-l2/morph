package hakeeper

import (
	"fmt"

	tmseq "github.com/tendermint/tendermint/proto/tendermint/sequencer"
	"github.com/tendermint/tendermint/types"
)

// encodeBlock serializes a BlockV2 into bytes for writing into the Raft log.
// Uses the existing tendermint proto path: BlockV2ToProto / proto.Marshal.
func encodeBlock(block *types.BlockV2) ([]byte, error) {
	pb := types.BlockV2ToProto(block)
	data, err := pb.Marshal()
	if err != nil {
		return nil, fmt.Errorf("encodeBlock: marshal failed: %w", err)
	}
	return data, nil
}

// decodeBlock deserializes a BlockV2 from bytes previously written to the Raft log.
func decodeBlock(data []byte) (*types.BlockV2, error) {
	var pb tmseq.BlockV2
	if err := pb.Unmarshal(data); err != nil {
		return nil, fmt.Errorf("decodeBlock: unmarshal failed: %w", err)
	}
	block, err := types.BlockV2FromProto(&pb)
	if err != nil {
		return nil, fmt.Errorf("decodeBlock: from proto failed: %w", err)
	}
	return block, nil
}
