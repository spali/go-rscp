package rscp

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/go-test/deep"
)

func Test_check(t *testing.T) {
	tests := []struct {
		name    string
		config  ClientConfig
		want    ClientConfig
		wantErr error
	}{
		{"empty config",
			ClientConfig{},
			ClientConfig{},
			errors.New("missing config values: address, username, password, key"),
		},
		{"valid minimal config",
			ClientConfig{Address: "127.0.0.1", Username: "testuser", Password: "testpassword", Key: "testkey"},
			ClientConfig{Address: "127.0.0.1", Port: defaultClientConfig.Port, Username: "testuser", Password: "testpassword", Key: "testkey", HeartbeatInterval: defaultClientConfig.HeartbeatInterval, ConnectionTimeout: defaultClientConfig.ConnectionTimeout, SendTimeout: defaultClientConfig.SendTimeout, ReceiveTimeout: defaultClientConfig.ReceiveTimeout, ReceiveBufferBlockSize: defaultClientConfig.ReceiveBufferBlockSize, UseChecksum: defaultClientConfig.UseChecksum},
			nil,
		},
		{"valid full config overriding all defaults",
			ClientConfig{Address: "127.0.0.1", Port: 5555, Username: "testuser", Password: "testpassword", Key: "testkey", HeartbeatInterval: time.Second * 60, ConnectionTimeout: 1, SendTimeout: 1, ReceiveTimeout: 1, ReceiveBufferBlockSize: 1, UseChecksum: false},
			ClientConfig{Address: "127.0.0.1", Port: 5555, Username: "testuser", Password: "testpassword", Key: "testkey", HeartbeatInterval: time.Second * 60, ConnectionTimeout: 1, SendTimeout: 1, ReceiveTimeout: 1, ReceiveBufferBlockSize: 1, UseChecksum: false},
			nil,
		},
		{"autofix values",
			ClientConfig{Address: "127.0.0.1", Port: 0, Username: "testuser", Password: "testpassword", Key: "testkey", HeartbeatInterval: -1, ConnectionTimeout: -1, SendTimeout: -1, ReceiveTimeout: -1, ReceiveBufferBlockSize: RSCP_FRAME_MAX_BLOCK_SIZE + 1},
			ClientConfig{Address: "127.0.0.1", Port: defaultClientConfig.Port, Username: "testuser", Password: "testpassword", Key: "testkey", HeartbeatInterval: defaultClientConfig.HeartbeatInterval, ConnectionTimeout: defaultClientConfig.ConnectionTimeout, SendTimeout: defaultClientConfig.SendTimeout, ReceiveTimeout: defaultClientConfig.ReceiveTimeout, ReceiveBufferBlockSize: defaultClientConfig.ReceiveBufferBlockSize, UseChecksum: defaultClientConfig.UseChecksum},
			nil,
		},
		{"wrong UseChecksum type",
			ClientConfig{Address: "127.0.0.1", Username: "testuser", Password: "testpassword", Key: "testkey", UseChecksum: "notvalid"},
			ClientConfig{Address: "127.0.0.1", Port: defaultClientConfig.Port, Username: "testuser", Password: "testpassword", Key: "testkey", HeartbeatInterval: defaultClientConfig.HeartbeatInterval, ConnectionTimeout: defaultClientConfig.ConnectionTimeout, SendTimeout: defaultClientConfig.SendTimeout, ReceiveTimeout: defaultClientConfig.ReceiveTimeout, ReceiveBufferBlockSize: defaultClientConfig.ReceiveBufferBlockSize, UseChecksum: "notvalid"},
			fmt.Errorf("config value UseChecksum has to be nil or of type bool, got %T", "notvalid"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.check()
			if (err != nil || tt.wantErr != nil) && fmt.Sprintf("%s", err) != fmt.Sprintf("%s", tt.wantErr) {
				t.Errorf("check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(tt.config, tt.want); diff != nil {
				t.Errorf("check() = %v, want %v\n%s", tt.config, tt.want, diff)
			}
		})
	}
}
