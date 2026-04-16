package hakeeper

import (
	"fmt"
	"math"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

// Config defines the configuration for hakeeper.
type Config struct {
	Enabled    bool     `mapstructure:"enabled"`
	ServerID   string   `mapstructure:"server_id"`
	StorageDir string   `mapstructure:"storage_dir"`
	Bootstrap  bool     `mapstructure:"bootstrap"`
	JoinAddrs  []string `mapstructure:"join_addrs"`

	// Debug enables verbose Raft internal logging. Set automatically when
	// the node's log level is "debug". Not a config file / env option.
	Debug bool `mapstructure:"-"`

	Consensus ConsensusConfig `mapstructure:"consensus"`
	Snapshot  SnapshotConfig  `mapstructure:"snapshot"`
	Timeout   TimeoutConfig   `mapstructure:"timeout"`
	RPC       RPCConfig       `mapstructure:"rpc"`
}

type ConsensusConfig struct {
	ListenAddr     string `mapstructure:"listen_addr"`
	ListenPort     int    `mapstructure:"listen_port"`
	AdvertisedAddr string `mapstructure:"advertised_addr"`
}

type SnapshotConfig struct {
	Interval     time.Duration `mapstructure:"interval"`
	Threshold    uint64        `mapstructure:"threshold"`
	TrailingLogs uint64        `mapstructure:"trailing_logs"`
}

type TimeoutConfig struct {
	Heartbeat   time.Duration `mapstructure:"heartbeat"`
	LeaderLease time.Duration `mapstructure:"leader_lease"`
}

type RPCConfig struct {
	ListenAddr string `mapstructure:"listen_addr"`
	ListenPort int    `mapstructure:"listen_port"`
	Token      string `mapstructure:"token"`
}

// ── Step 1: Defaults ─────────────────────────────────────────────────────────

// DefaultConfig returns the default configuration with sensible values
// for all common/generic settings. Node-specific fields (ServerID, StorageDir,
// AdvertisedAddr) are left empty for Resolve() to auto-detect.
func DefaultConfig() *Config {
	return &Config{
		Consensus: ConsensusConfig{
			ListenAddr: "0.0.0.0",
			ListenPort: 9400,
		},
		Snapshot: SnapshotConfig{
			Interval:     120 * time.Second,
			Threshold:    8192,
			TrailingLogs: 1200,
		},
		Timeout: TimeoutConfig{
			Heartbeat:   1 * time.Second,
			LeaderLease: 500 * time.Millisecond,
		},
		RPC: RPCConfig{
			ListenAddr: "0.0.0.0",
			ListenPort: 9401,
		},
	}
}

// ── Step 2: Config file overlay (optional) ───────────────────────────────────

// LoadFile reads a TOML config file and overlays values onto c.
// Only fields present in the file are overwritten; others keep their current value.
func (c *Config) LoadFile(path string) error {
	dir := filepath.Dir(path)
	filename := filepath.Base(path)
	ext := filepath.Ext(filename)
	name := filename[:len(filename)-len(ext)]

	v := viper.New()
	v.AddConfigPath(dir)
	v.SetConfigName(name)
	v.SetConfigType("toml")

	if err := v.ReadInConfig(); err != nil {
		return errors.Wrap(err, "failed to read HA config file")
	}
	if err := v.Unmarshal(c); err != nil {
		return errors.Wrap(err, "failed to parse HA config file")
	}
	return nil
}

// ── Step 3: Auto-resolve node-specific fields ────────────────────────────────

// Resolve fills in empty node-specific fields with auto-detected values:
//   - ServerID  → os.Hostname()
//   - StorageDir → <homeDir>/raft
//   - AdvertisedAddr → local non-loopback IP (if ListenAddr is 0.0.0.0)
//
// Call this AFTER flag overrides have been applied and BEFORE Validate().
func (c *Config) Resolve(homeDir string) error {
	// ServerID
	if c.ServerID == "" {
		hostname, err := os.Hostname()
		if err != nil {
			return fmt.Errorf("server_id not set and hostname detection failed: %w", err)
		}
		if hostname == "" {
			return fmt.Errorf("server_id not set and hostname is empty")
		}
		c.ServerID = hostname
	}

	// StorageDir
	if c.StorageDir == "" {
		c.StorageDir = filepath.Join(homeDir, "raft")
	}

	// AdvertisedAddr
	if c.Consensus.AdvertisedAddr == "" {
		addr, err := resolveAdvertisedAddr(c.Consensus.ListenAddr, c.Consensus.ListenPort)
		if err != nil {
			return err
		}
		c.Consensus.AdvertisedAddr = addr
	}

	return nil
}

// resolveAdvertisedAddr derives the advertised address when not explicitly set.
func resolveAdvertisedAddr(listenAddr string, listenPort int) (string, error) {
	port := fmt.Sprintf("%d", listenPort)

	// If ListenAddr is a specific IP, use it directly.
	if listenAddr != "0.0.0.0" && listenAddr != "" {
		return net.JoinHostPort(listenAddr, port), nil
	}

	// Auto-detect: first non-loopback IPv4 on any active interface.
	ip, err := localNonLoopbackIP()
	if err != nil {
		return "", fmt.Errorf("advertised_addr not set and auto-detect failed: %w", err)
	}
	return net.JoinHostPort(ip, port), nil
}

func localNonLoopbackIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip4 := ip.To4(); ip4 != nil && !ip4.IsLoopback() {
				return ip4.String(), nil
			}
		}
	}
	return "", fmt.Errorf("no non-loopback IPv4 address found")
}

// ── Step 4: Validate ─────────────────────────────────────────────────────────

// Validate checks that all required fields are present. Call AFTER Resolve().
func (c *Config) Validate() error {
	if c.ServerID == "" {
		return fmt.Errorf("server_id is required (set via config, --ha.server-id, or ensure hostname is available)")
	}
	if c.StorageDir == "" {
		return fmt.Errorf("storage_dir is required")
	}
	if c.Consensus.ListenPort < 0 || c.Consensus.ListenPort > math.MaxUint16 {
		return fmt.Errorf("invalid consensus.listen_port: %d", c.Consensus.ListenPort)
	}
	if c.RPC.ListenPort < 0 || c.RPC.ListenPort > math.MaxUint16 {
		return fmt.Errorf("invalid rpc.listen_port: %d", c.RPC.ListenPort)
	}

	// AdvertisedAddr must be a routable address (IP or hostname) after Resolve().
	if c.Consensus.AdvertisedAddr != "" {
		host, _, err := net.SplitHostPort(c.Consensus.AdvertisedAddr)
		if err != nil {
			return fmt.Errorf("invalid consensus.advertised_addr %q: %w", c.Consensus.AdvertisedAddr, err)
		}
		if host == "0.0.0.0" || host == "" {
			return fmt.Errorf("consensus.advertised_addr must be a specific address, not %q", host)
		}
	}

	// Follower must have at least one address to join.
	if !c.Bootstrap && len(c.JoinAddrs) == 0 {
		return fmt.Errorf("join_addrs is required when bootstrap=false (set via config or --ha.join)")
	}

	return nil
}

// ── Print effective config ───────────────────────────────────────────────────

// LogEffectiveConfig prints the resolved HA configuration for operator visibility.
func (c *Config) LogEffectiveConfig(logger tmlog.Logger) {
	role := "follower"
	if c.Bootstrap {
		role = "bootstrap-leader"
	}
	joinAddrs := "(none)"
	if len(c.JoinAddrs) > 0 {
		joinAddrs = strings.Join(c.JoinAddrs, ", ")
	}

	logger.Info("========== HA Effective Config ==========")
	logger.Info("ha config",
		"role", role,
		"server_id", c.ServerID,
		"advertised_addr", c.Consensus.AdvertisedAddr,
		"storage_dir", c.StorageDir,
		"join_addrs", joinAddrs,
	)
	logger.Info("ha config",
		"raft_listen", fmt.Sprintf("%s:%d", c.Consensus.ListenAddr, c.Consensus.ListenPort),
		"rpc_listen", fmt.Sprintf("%s:%d", c.RPC.ListenAddr, c.RPC.ListenPort),
		"heartbeat", c.Timeout.Heartbeat,
		"leader_lease", c.Timeout.LeaderLease,
		"trailing_logs", c.Snapshot.TrailingLogs,
	)
	logger.Info("=========================================")
}
