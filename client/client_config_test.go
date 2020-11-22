package client

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/spali/go-e3dc/rscp"
)

func Test_check(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		want    Config
		wantErr error
	}{
		{"empty config",
			Config{},
			Config{},
			errors.New("missing config values: address, username, password, key"),
		},
		{"valid minimal config",
			Config{Address: "127.0.0.1", Username: "testuser", Password: "testpassword", Key: "testkey"},
			Config{Address: "127.0.0.1", Port: defaultClientConfig.Port, Username: "testuser", Password: "testpassword", Key: "testkey", HeartbeatInterval: defaultClientConfig.HeartbeatInterval, ConnectionTimeout: defaultClientConfig.ConnectionTimeout, SendTimeout: defaultClientConfig.SendTimeout, ReceiveTimeout: defaultClientConfig.ReceiveTimeout, ReceiveBufferBlockSize: defaultClientConfig.ReceiveBufferBlockSize, UseChecksum: defaultClientConfig.UseChecksum},
			nil,
		},
		{"valid full config overriding all defaults",
			Config{Address: "127.0.0.1", Port: 5555, Username: "testuser", Password: "testpassword", Key: "testkey", HeartbeatInterval: time.Second * 60, ConnectionTimeout: 1, SendTimeout: 1, ReceiveTimeout: 1, ReceiveBufferBlockSize: 1, UseChecksum: false},
			Config{Address: "127.0.0.1", Port: 5555, Username: "testuser", Password: "testpassword", Key: "testkey", HeartbeatInterval: time.Second * 60, ConnectionTimeout: 1, SendTimeout: 1, ReceiveTimeout: 1, ReceiveBufferBlockSize: 1, UseChecksum: false},
			nil,
		},
		{"autofix values",
			Config{Address: "127.0.0.1", Port: 0, Username: "testuser", Password: "testpassword", Key: "testkey", HeartbeatInterval: -1, ConnectionTimeout: -1, SendTimeout: -1, ReceiveTimeout: -1, ReceiveBufferBlockSize: rscp.RSCP_FRAME_MAX_BLOCK_SIZE + 1},
			Config{Address: "127.0.0.1", Port: defaultClientConfig.Port, Username: "testuser", Password: "testpassword", Key: "testkey", HeartbeatInterval: defaultClientConfig.HeartbeatInterval, ConnectionTimeout: defaultClientConfig.ConnectionTimeout, SendTimeout: defaultClientConfig.SendTimeout, ReceiveTimeout: defaultClientConfig.ReceiveTimeout, ReceiveBufferBlockSize: defaultClientConfig.ReceiveBufferBlockSize, UseChecksum: defaultClientConfig.UseChecksum},
			nil,
		},
		{"wrong UseChecksum type",
			Config{Address: "127.0.0.1", Username: "testuser", Password: "testpassword", Key: "testkey", UseChecksum: "notvalid"},
			Config{Address: "127.0.0.1", Port: defaultClientConfig.Port, Username: "testuser", Password: "testpassword", Key: "testkey", HeartbeatInterval: defaultClientConfig.HeartbeatInterval, ConnectionTimeout: defaultClientConfig.ConnectionTimeout, SendTimeout: defaultClientConfig.SendTimeout, ReceiveTimeout: defaultClientConfig.ReceiveTimeout, ReceiveBufferBlockSize: defaultClientConfig.ReceiveBufferBlockSize, UseChecksum: "notvalid"},
			fmt.Errorf("config value UseChecksum has to be nil or of type bool, got %T", "notvalid"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := check(tt.config)
			if (err != nil || tt.wantErr != nil) && fmt.Sprintf("%s", err) != fmt.Sprintf("%s", tt.wantErr) {
				t.Errorf("check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("check() = %v, want %v\n%s", got, tt.want, diff)
			}
		})
	}
}
