package node

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	tmconfig "github.com/tendermint/tendermint/config"
	tmlog "github.com/tendermint/tendermint/libs/log"
	"github.com/urfave/cli"
	"gopkg.in/natefinch/lumberjack.v2"

	"morph-l2/bindings/predeploys"
	"morph-l2/node/flags"
	"morph-l2/node/types"
)

type Config struct {
	L2                            *types.L2Config `json:"l2"`
	L2CrossDomainMessengerAddress common.Address  `json:"cross_domain_messenger_address"`
	SequencerAddress              common.Address  `json:"sequencer_address"`
	GovAddress                    common.Address  `json:"gov_address"`
	L2StakingAddress              common.Address  `json:"l2staking_address"`
	MaxL1MessageNumPerBlock       uint64          `json:"max_l1_message_num_per_block"`
	DevSequencer                  bool            `json:"dev_sequencer"`
	Logger                        tmlog.Logger    `json:"logger"`
}

func DefaultConfig() *Config {
	return &Config{
		L2:                            new(types.L2Config),
		Logger:                        tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout)),
		MaxL1MessageNumPerBlock:       100,
		L2CrossDomainMessengerAddress: predeploys.L2CrossDomainMessengerAddr,
		SequencerAddress:              predeploys.SequencerAddr,
		GovAddress:                    predeploys.GovAddr,
		L2StakingAddress:              predeploys.L2StakingAddr,
	}
}

func (c *Config) SetCliContext(ctx *cli.Context) error {
	// logger setting
	output := io.Writer(os.Stderr)
	if ctx.GlobalIsSet(flags.LogFilename.Name) {
		logFilename := ctx.GlobalString(flags.LogFilename.Name)
		f, err := os.OpenFile(filepath.Clean(logFilename), os.O_CREATE|os.O_RDWR, os.FileMode(0600))
		if err != nil {
			return fmt.Errorf("wrong log.filename set: %d", err)
		}
		_ = f.Close()
		maxSize := ctx.GlobalInt(flags.LogFileMaxSize.Name)
		if maxSize < 1 {
			return fmt.Errorf("wrong log.maxsize set: %d", maxSize)
		}
		maxAge := ctx.GlobalInt(flags.LogFileMaxAge.Name)
		if maxAge < 1 {
			return fmt.Errorf("wrong log.maxage set: %d", maxAge)
		}
		logFile := &lumberjack.Logger{
			Filename: logFilename,
			MaxSize:  maxSize, // megabytes
			MaxAge:   maxAge,  // days
			Compress: ctx.GlobalBool(flags.LogCompress.Name),
		}
		output = io.MultiWriter(output, logFile)
	}

	logger := tmlog.NewTMLogger(tmlog.NewSyncWriter(output))
	if format := ctx.GlobalString(flags.LogFormat.Name); len(format) > 0 && format == tmconfig.LogFormatJSON {
		logger = tmlog.NewTMJSONLogger(tmlog.NewSyncWriter(output))
	}

	logLevel := "info"
	if ctx.GlobalIsSet(flags.LogLevel.Name) {
		logLevel = ctx.GlobalString(flags.LogLevel.Name)
	}
	option, err := tmlog.AllowLevel(logLevel)
	if err != nil {
		return err
	}
	logger = tmlog.NewFilter(logger, option)
	c.Logger = logger

	l2EthAddr := ctx.GlobalString(flags.L2EthAddr.Name)
	l2EngineAddr := ctx.GlobalString(flags.L2EngineAddr.Name)
	fileName := ctx.GlobalString(flags.L2EngineJWTSecret.Name)
	var secret [32]byte
	fileName = strings.TrimSpace(fileName)
	if fileName == "" {
		return fmt.Errorf("file-name of jwt secret is empty")
	}
	if data, err := os.ReadFile(filepath.Clean(fileName)); err == nil {
		jwtSecret := common.FromHex(strings.TrimSpace(string(data)))
		if len(jwtSecret) != 32 {
			return fmt.Errorf("invalid jwt secret in path %s, not 32 hex-formatted bytes", fileName)
		}
		copy(secret[:], jwtSecret)
	} else {
		logger.Info("Failed to read JWT secret from file, generating a new one now. Configure L2 geth with --authrpc.jwt-secret=" + fmt.Sprintf("%q", fileName))
		if _, err := io.ReadFull(rand.Reader, secret[:]); err != nil {
			return fmt.Errorf("failed to generate jwt secret: %w", err)
		}
		if err := os.WriteFile(fileName, []byte(hexutil.Encode(secret[:])), 0600); err != nil {
			return err
		}
	}
	c.L2.EthAddr = l2EthAddr
	c.L2.EngineAddr = l2EngineAddr
	c.L2.JwtSecret = secret

	if ctx.GlobalIsSet(flags.MaxL1MessageNumPerBlock.Name) {
		c.MaxL1MessageNumPerBlock = ctx.GlobalUint64(flags.MaxL1MessageNumPerBlock.Name)
		if c.MaxL1MessageNumPerBlock == 0 {
			return fmt.Errorf("MaxL1MessageNumPerBlock must be above 0")
		}
	}

	if ctx.GlobalIsSet(flags.L2CrossDomainMessengerContractAddr.Name) {
		addr := common.HexToAddress(ctx.GlobalString(flags.L2CrossDomainMessengerContractAddr.Name))
		c.L2CrossDomainMessengerAddress = addr
		if len(c.L2CrossDomainMessengerAddress.Bytes()) == 0 {
			return errors.New("invalid L2CrossDomainMessengerContractAddr")
		}
	}

	if ctx.GlobalIsSet(flags.L2SequencerAddr.Name) {
		addr := common.HexToAddress(ctx.GlobalString(flags.L2SequencerAddr.Name))
		c.SequencerAddress = addr
		if len(c.SequencerAddress.Bytes()) == 0 {
			return errors.New("invalid L2SequencerAddr")
		}
	}

	if ctx.GlobalIsSet(flags.GovAddr.Name) {
		addr := common.HexToAddress(ctx.GlobalString(flags.GovAddr.Name))
		c.GovAddress = addr
		if len(c.GovAddress.Bytes()) == 0 {
			return errors.New("invalid GovAddr")
		}
	}

	if ctx.GlobalIsSet(flags.DevSequencer.Name) {
		c.DevSequencer = ctx.GlobalBool(flags.DevSequencer.Name)
	}

	return nil
}
