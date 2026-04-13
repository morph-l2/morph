package l1sequencer

import (
	"context"
	"time"

	"github.com/morph-l2/go-ethereum/ethclient"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

// L1Tracker monitors L1 RPC sync status and logs warnings if behind.
// It runs as an independent service.
type L1Tracker struct {
	ctx          context.Context
	cancel       context.CancelFunc
	l1Client     *ethclient.Client
	lagThreshold time.Duration
	logger       tmlog.Logger
	stop         chan struct{}
}

// NewL1Tracker creates a new L1Tracker
func NewL1Tracker(
	ctx context.Context,
	l1Client *ethclient.Client,
	lagThreshold time.Duration,
	logger tmlog.Logger,
) *L1Tracker {
	ctx, cancel := context.WithCancel(ctx)
	return &L1Tracker{
		ctx:          ctx,
		cancel:       cancel,
		l1Client:     l1Client,
		lagThreshold: lagThreshold,
		logger:       logger.With("module", "l1tracker"),
		stop:         make(chan struct{}),
	}
}

// Start starts the L1Tracker
func (t *L1Tracker) Start() error {
	t.logger.Info("Starting L1Tracker", "lagThreshold", t.lagThreshold)
	go t.loop()
	return nil
}

// Stop stops the L1Tracker
func (t *L1Tracker) Stop() {
	t.logger.Info("Stopping L1Tracker")
	t.cancel()
	<-t.stop
}

func (t *L1Tracker) loop() {
	defer close(t.stop)

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-t.ctx.Done():
			return
		case <-ticker.C:
			t.checkL1SyncLag()
		}
	}
}

func (t *L1Tracker) checkL1SyncLag() {
	header, err := t.l1Client.HeaderByNumber(t.ctx, nil)
	if err != nil {
		t.logger.Error("Failed to get L1 header", "error", err)
		return
	}

	blockTime := time.Unix(int64(header.Time), 0)
	lag := time.Since(blockTime)
	if lag > t.lagThreshold {
		t.logger.Error("L1 RPC is behind",
			"latestBlock", header.Number.Uint64(),
			"blockTime", blockTime.Format(time.RFC3339),
			"lag", lag.Round(time.Second),
		)
	}
}
