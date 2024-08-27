package oracle

import "errors"

// empirical value
const maxBatchSize = 72

var (
	ErrRewardNotStart = errors.New("reward has not start")
)
