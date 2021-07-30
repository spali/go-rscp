package rscp

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/go-test/deep"
)

func Test_write(t *testing.T) {
	type args struct {
		buf *bytes.Buffer
		v   interface{}
	}
	tests := []struct {
		name         string
		args         args
		wantBufBytes []byte
		wantErr      error
	}{

		{"nil value",
			args{
				bytes.NewBuffer([]byte{}),
				nil,
			},
			[]byte{},
			nil,
		},
		{"simple type",
			args{
				bytes.NewBuffer([]byte{}),
				uint8(255),
			},
			[]byte{255},
			nil,
		},
		{"string",
			args{
				bytes.NewBuffer([]byte{}),
				"test",
			},
			[]byte{116, 101, 115, 116},
			nil,
		},
		{"time",
			args{
				bytes.NewBuffer([]byte{}),
				time.Date(1234, 5, 6, 7, 8, 9, 123456000, time.UTC), // 0xd9, 0x74, 0x46, 0x98, 0xfa, 0xff, 0xff, 0xff, 0x0, 0xca, 0x5b, 0x7
			},
			[]byte{0xd9, 0x74, 0x46, 0x98, 0xfa, 0xff, 0xff, 0xff, 0x0, 0xca, 0x5b, 0x7},
			nil,
		},
		{"byte array",
			args{
				bytes.NewBuffer([]byte{}),
				[]Message{{WB_EXTERN_DATA, ByteArray, []byte{0x0, 0x1, 0x2}}},
			},
			[]byte{0x10, 0x20, 0x4, 0xe, 0x10, 0x3, 0x0, 0x0, 0x1, 0x2},
			nil,
		},
		{"messages",
			args{
				bytes.NewBuffer([]byte{}),
				[]Message{{RSCP_AUTHENTICATION, UChar8, uint8(10)}, {RSCP_AUTHENTICATION, UChar8, uint8(10)}},
			},
			[]byte{0x1, 0x0, 0x80, 0x0, 0x3, 0x1, 0x0, 0xa, 0x1, 0x0, 0x80, 0x0, 0x3, 0x1, 0x0, 0xa},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := write(tt.args.buf, tt.args.v)
			if (err != nil || tt.wantErr != nil) && !(errors.Is(err, tt.wantErr) || err.Error() == tt.wantErr.Error()) {
				t.Errorf("write() error = %#v, wantErr %#v", err, tt.wantErr)
			}
			if diff := deep.Equal(tt.args.buf.Bytes(), tt.wantBufBytes); diff != nil {
				t.Errorf("write() = %#v, want %#v\n%s", tt.args.buf.Bytes(), tt.wantBufBytes, diff)
			}
		})
	}
}
