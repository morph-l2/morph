package log

import (
	"fmt"
	"strings"

	"github.com/morph-l2/go-ethereum/log"
	"github.com/urfave/cli"
)

const (
	LevelFlagName  = "log.level"
	FormatFlagName = "log.format"
	ColorFlagName  = "log.color"
)

func CLIFlags(envPrefix string) []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   LevelFlagName,
			Usage:  "The lowest log level that will be output",
			Value:  "info",
			EnvVar: PrefixEnvVar(envPrefix, "LOG_LEVEL"),
		},
		cli.StringFlag{
			Name:   FormatFlagName,
			Usage:  "Format the log output. Supported formats: 'text', 'terminal', 'logfmt', 'json', 'json-pretty',",
			Value:  "text",
			EnvVar: PrefixEnvVar(envPrefix, "LOG_FORMAT"),
		},
		cli.BoolFlag{
			Name:   ColorFlagName,
			Usage:  "Color the log output if in terminal mode",
			EnvVar: PrefixEnvVar(envPrefix, "LOG_COLOR"),
		},
	}
}

func PrefixEnvVar(prefix, suffix string) string {
	return prefix + "_" + suffix
}

type CLIConfig struct {
	Level  string // Log level: trace, debug, info, warn, error, crit. Capitals are accepted too.
	Color  bool   // Color the log output. Defaults to true if terminal is detected.
	Format string // Format the log output. Supported formats: 'text', 'terminal', 'logfmt', 'json', 'json-pretty'
}

func (cfg CLIConfig) Check() error {
	switch cfg.Format {
	case "json", "json-pretty", "terminal", "text", "logfmt":
	default:
		return fmt.Errorf("unrecognized log format: %s", cfg.Format)
	}

	level := strings.ToLower(cfg.Level)
	_, err := log.LvlFromString(level)
	if err != nil {
		return fmt.Errorf("unrecognized log level: %w", err)
	}
	return nil
}
