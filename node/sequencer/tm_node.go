package sequencer

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	tmtypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/blssignatures"
	"github.com/tendermint/tendermint/config"
	tmflags "github.com/tendermint/tendermint/libs/cli/flags"
	tmlog "github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmnode "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/proxy"
	"github.com/tendermint/tendermint/types"
	"github.com/urfave/cli"

	node "morph-l2/node/core"
	"morph-l2/node/flags"
	nodetypes "morph-l2/node/types"
)

func LoadTmConfig(ctx *cli.Context, home string) (*config.Config, error) {
	var (
		tmCfg      *config.Config
		configPath = ctx.GlobalString(flags.TendermintConfigPath.Name)
	)
	if configPath == "" {
		if home == "" {
			return nil, fmt.Errorf("either Home or Config Path has to be provided")
		}
		configPath = filepath.Join(home, "config")
	}
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	tmCfg = config.DefaultConfig()

	if err := viper.Unmarshal(tmCfg); err != nil {
		return nil, err
	}

	tmCfg.SetRoot(home)
	if err := tmCfg.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("error in config file: %w", err)
	}
	return tmCfg, nil
}

func SetupNode(tmCfg *config.Config, privValidator types.PrivValidator, executor *node.Executor, logger tmlog.Logger) (*tmnode.Node, error) {
	if tmCfg.LogFormat == config.LogFormatJSON {
		logger = tmlog.NewTMJSONLogger(tmlog.NewSyncWriter(os.Stdout))
	}
	nodeLogger, err := tmflags.ParseLogLevel(tmCfg.LogLevel, logger, config.DefaultLogLevel)
	if err != nil {
		return nil, err
	}
	nodeLogger = nodeLogger.With("module", "main")

	nodeKey, err := p2p.LoadOrGenNodeKey(tmCfg.NodeKeyFile())
	if err != nil {
		return nil, err
	}

	if !tmos.FileExists(tmCfg.BLSKeyFile()) {
		blssignatures.GenFileBLSKey().Save(tmCfg.BLSKeyFile())
	}
	blsPrivKey, err := blssignatures.PrivateKeyFromBytes(blssignatures.LoadBLSKey(tmCfg.BLSKeyFile()).PrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed to load bls priv key")
	}

	//var app types.Application
	n, err := tmnode.NewNode(
		tmCfg,
		executor,
		privValidator,
		&blsPrivKey,
		nodeKey,
		proxy.NewLocalClientCreator(NewApplication(tmtypes.NewBaseApplication(), executor.L2Client())),
		tmnode.DefaultGenesisDocProviderFunc(tmCfg),
		tmnode.DefaultDBProvider,
		tmnode.DefaultMetricsProvider(tmCfg.Instrumentation),
		nodeLogger,
	)
	return n, err
}

type Application struct {
	*tmtypes.BaseApplication
	l2Client *nodetypes.RetryableClient
}

func NewApplication(baseApp *tmtypes.BaseApplication, l2Client *nodetypes.RetryableClient) Application {
	return Application{
		BaseApplication: baseApp,
		l2Client:        l2Client,
	}
}

func (a Application) Info(req tmtypes.RequestInfo) tmtypes.ResponseInfo {
	blockNumber, err := a.l2Client.BlockNumber(context.Background())
	if err != nil {
		return tmtypes.ResponseInfo{}
	}
	return tmtypes.ResponseInfo{
		LastBlockHeight: int64(blockNumber),
	}
}
