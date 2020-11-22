package client

import (
	"fmt"
	"strings"
	"time"

	"github.com/spali/go-e3dc/rscp"
)

// Config allows to configure the client behavior
type Config struct {
	// Host address
	Address string
	// Port
	Port uint16
	// username for authentication
	Username string
	// password for authentication
	Password string
	// key for encryption
	Key string
	// time inverall between heatbeat connection requests
	HeartbeatInterval time.Duration
	// Timeout for connection attempt
	ConnectionTimeout time.Duration
	// Timeout waiting for sending data (should not occur under normal conditions)
	SendTimeout time.Duration
	// Timeout waiting for response
	ReceiveTimeout time.Duration
	// define if CRC is uses and checked during communication. if nil, default setting is used. Has to be nil or of type bool.
	UseChecksum interface{}
	// amount of blocks of the receiving buffer size
	ReceiveBufferBlockSize uint16
}

// defaultClientConfig defines the default config values used when not provided by the user.
//nolint: gomnd
var defaultClientConfig = Config{
	Port:                   5033,
	HeartbeatInterval:      time.Second * 10,
	ConnectionTimeout:      time.Second * 3,
	SendTimeout:            time.Second * 3,
	ReceiveTimeout:         time.Second * 3,
	ReceiveBufferBlockSize: 1,
	UseChecksum:            true,
}

// check does set default values on missing or fail if required
func check(config Config) (Config, error) {
	var missing []string
	if len(config.Address) == 0 {
		missing = append(missing, "address")
	}
	if len(config.Username) == 0 {
		missing = append(missing, "username")
	}
	if len(config.Password) == 0 {
		missing = append(missing, "password")
	}
	if len(config.Key) == 0 {
		missing = append(missing, "key")
	}
	if len(missing) > 0 {
		return config, fmt.Errorf("missing config values: %s", strings.Join(missing, ", "))
	}
	if config.Port == 0 {
		config.Port = defaultClientConfig.Port
	}
	if config.HeartbeatInterval <= time.Second {
		config.HeartbeatInterval = defaultClientConfig.HeartbeatInterval
	}
	if config.ConnectionTimeout <= 0 {
		config.ConnectionTimeout = defaultClientConfig.ConnectionTimeout
	}
	if config.SendTimeout <= 0 {
		config.SendTimeout = defaultClientConfig.SendTimeout
	}
	if config.ReceiveTimeout <= 0 {
		config.ReceiveTimeout = defaultClientConfig.ReceiveTimeout
	}
	if config.ReceiveBufferBlockSize == 0 || config.ReceiveBufferBlockSize > rscp.RSCP_FRAME_MAX_BLOCK_SIZE {
		config.ReceiveBufferBlockSize = defaultClientConfig.ReceiveBufferBlockSize
	}
	if _, isBool := config.UseChecksum.(bool); !isBool && config.UseChecksum != nil {
		return config, fmt.Errorf("config value UseChecksum has to be nil or of type bool, got %T", config.UseChecksum)
	}
	if config.UseChecksum == nil {
		config.UseChecksum = defaultClientConfig.UseChecksum
	}
	return config, nil
}
