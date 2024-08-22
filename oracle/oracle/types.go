package oracle

import "errors"

const maxBatchSize = 72

var (
	ErrRewardNotStart = errors.New("reward has not start")
)
