package sync

import (
	"context"
	"errors"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/ethclient"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"morph-l2/node/types"
)

type Syncer struct {
	ctx          context.Context
	cancel       context.CancelFunc
	bridgeClient *BridgeClient
	latestSynced uint64
	db           Database
	logger       tmlog.Logger
	metrics      *Metrics

	fetchBlockRange     uint64
	pollInterval        time.Duration
	logProgressInterval time.Duration
	stop                chan struct{}
	isFake              bool
}

func NewSyncer(ctx context.Context, db Database, config *Config, logger tmlog.Logger) (*Syncer, error) {
	l1Client, err := ethclient.Dial(config.L1.Addr)
	if err != nil {
		return nil, err
	}

	if config.L1MessageQueueAddress == nil {
		return nil, errors.New("deposit contract address cannot be nil")
	}

	bridgeClient, err := NewBridgeClient(l1Client, *config.L1MessageQueueAddress, config.L1.Confirmations, logger)
	if err != nil {
		return nil, err
	}

	logger = logger.With("module", "syncer")
	latestSynced := db.ReadLatestSyncedL1Height()
	if latestSynced == nil {
		if config.StartHeight == 0 {
			logger.Info("syncing warning", "msg", "Missing `sync.startHeight` configured. Detected that it is your first time to start the node as a sequencer, it is dangerous not setting `sync.startHeight`, as it may lost some previous L1Messages")
			if config.StartHeight, err = l1Client.BlockNumber(context.Background()); err != nil {
				return nil, err
			}
		}
		h := config.StartHeight - 1
		latestSynced = &h
	}
	metrics := PrometheusMetrics("morphnode")
	metrics.SyncedL1Height.Set(float64(*latestSynced))

	ctx, cancel := context.WithCancel(ctx)
	return &Syncer{
		ctx:          ctx,
		cancel:       cancel,
		bridgeClient: bridgeClient,
		latestSynced: *latestSynced,
		db:           db,
		stop:         make(chan struct{}),
		logger:       logger,
		metrics:      metrics,

		fetchBlockRange:     config.FetchBlockRange,
		pollInterval:        config.PollInterval,
		logProgressInterval: config.LogProgressInterval,
	}, nil
}

func (s *Syncer) Start() {
	if s.isFake {
		return
	}
	// block node startup during initial sync and print some helpful logs
	s.logger.Info("initial sync start", "msg", "Running initial sync of L1 messages before starting sequencer, this might take a while...")
	s.fetchL1Messages()
	s.logger.Info("initial sync completed", "latestSyncedBlock", s.latestSynced)

	go func() {
		t := time.NewTicker(s.pollInterval)
		defer t.Stop()

		for {
			// don't wait for ticker during startup
			s.fetchL1Messages()

			select {
			case <-s.ctx.Done():
				close(s.stop)
				return
			case <-t.C:
				continue
			}
		}
	}()
}

func (s *Syncer) Stop() {
	if s == nil {
		return
	}

	s.logger.Info("Stopping sync service")

	if s.cancel != nil {
		s.cancel()
	}
	<-s.stop
	s.logger.Info("Sync service is stopped")
}

func (s *Syncer) fetchL1Messages() {
	latestConfirmed, err := s.bridgeClient.getLatestConfirmedBlockNumber(s.ctx)
	if err != nil {
		s.logger.Error("failed to get latest confirmed block number", "err", err)
		return
	}

	// ticker for logging progress
	t := time.NewTicker(s.logProgressInterval)
	numMessagesCollected := 0
	// query in batches
	for from := s.latestSynced + 1; from <= latestConfirmed; from += s.fetchBlockRange {
		select {
		case <-s.ctx.Done():
			return
		case <-t.C:
			progress := 100 * float64(s.latestSynced) / float64(latestConfirmed)
			s.logger.Info("Syncing L1 messages", "synced", s.latestSynced, "confirmed", latestConfirmed, "collected", numMessagesCollected, "progress(%)", progress)
		default:
		}

		to := from + s.fetchBlockRange - 1

		if to > latestConfirmed {
			to = latestConfirmed
		}

		l1Messages, err := s.bridgeClient.L1Messages(s.ctx, from, to)
		if err != nil {
			s.logger.Error("failed to fetch L1 messages", "fromBlock", from, "toBlock", to, "err", err)
			return
		}

		if len(l1Messages) > 0 {
			s.logger.Debug("Received new L1 events", "fromBlock", from, "toBlock", to, "count", len(l1Messages))
			if err = s.db.WriteSyncedL1Messages(l1Messages, to); err != nil {
				// crash on database error
				s.logger.Error("failed to write L1 messages to database", "err", err)
				return
			}
			numMessagesCollected += len(l1Messages)

			s.metrics.SyncedL1MessageCount.Add(float64(len(l1Messages)))
			s.metrics.SyncedL1MessageNonce.Set(float64(l1Messages[len(l1Messages)-1].QueueIndex))
		} else {
			s.db.WriteLatestSyncedL1Height(to)
		}
		s.latestSynced = to

		s.metrics.SyncedL1Height.Set(float64(to))
	}
}

func (s *Syncer) GetL1Message(index uint64, txHash common.Hash) (*types.L1Message, error) {
	l1Message := s.db.ReadL1MessageByIndex(index)
	if l1Message != nil {
		return l1Message, nil
	}

	l1Messages, err := s.bridgeClient.L1MessagesFromTxHash(s.ctx, txHash)
	if err != nil {
		return nil, err
	}

	for _, msg := range l1Messages {
		if msg.QueueIndex == index {
			return &msg, nil
		}
	}
	return nil, nil
}

func (s *Syncer) ReadL1MessagesInRange(start, end uint64) []types.L1Message {
	return s.db.ReadL1MessagesInRange(start, end)
}

func (s *Syncer) LatestSynced() uint64 {
	return s.latestSynced
}
