package rscp

import (
	"fmt"
	"strings"
	"time"
)

// ClientConfig allows to configure the client behavior
type ClientConfig struct {
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
//
//nolint:gomnd
var defaultClientConfig = ClientConfig{
	Port:                   5033,
	HeartbeatInterval:      time.Second * 10,
	ConnectionTimeout:      time.Second * 3,
	SendTimeout:            time.Second * 3,
	ReceiveTimeout:         time.Second * 3,
	ReceiveBufferBlockSize: 1,
	UseChecksum:            true,
}

// check does set default values on missing or fail if required
func (c *ClientConfig) check() error {
	var missing []string
	if len(c.Address) == 0 {
		missing = append(missing, "address")
	}
	if len(c.Username) == 0 {
		missing = append(missing, "username")
	}
	if len(c.Password) == 0 {
		missing = append(missing, "password")
	}
	if len(c.Key) == 0 {
		missing = append(missing, "key")
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing config values: %s", strings.Join(missing, ", "))
	}
	if c.Port == 0 {
		c.Port = defaultClientConfig.Port
	}
	if c.HeartbeatInterval <= time.Second {
		c.HeartbeatInterval = defaultClientConfig.HeartbeatInterval
	}
	if c.ConnectionTimeout <= 0 {
		c.ConnectionTimeout = defaultClientConfig.ConnectionTimeout
	}
	if c.SendTimeout <= 0 {
		c.SendTimeout = defaultClientConfig.SendTimeout
	}
	if c.ReceiveTimeout <= 0 {
		c.ReceiveTimeout = defaultClientConfig.ReceiveTimeout
	}
	if c.ReceiveBufferBlockSize == 0 || c.ReceiveBufferBlockSize > RSCP_FRAME_MAX_BLOCK_SIZE {
		c.ReceiveBufferBlockSize = defaultClientConfig.ReceiveBufferBlockSize
	}
	if _, isBool := c.UseChecksum.(bool); !isBool && c.UseChecksum != nil {
		return fmt.Errorf("config value UseChecksum has to be nil or of type bool, got %T", c.UseChecksum)
	}
	if c.UseChecksum == nil {
		c.UseChecksum = defaultClientConfig.UseChecksum
	}
	return nil
}
