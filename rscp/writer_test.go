package rscp

import (
	"bytes"
	"crypto/cipher"
	"errors"
	"testing"
	"time"

	"github.com/azihsoyn/rijndael256"
	"github.com/go-test/deep"
)

func Test_write(t *testing.T) {
	tests := []struct {
		name         string
		data         interface{} // go representation of the data to write
		wantBufBytes []byte      // byte representation of the data written
		wantErr      error
	}{

		{"nil value",
			nil,
			[]byte{},
			nil,
		},
		{"simple type",
			uint8(255),
			[]byte{255},
			nil,
		},
		{"string",
			"test",
			[]byte{116, 101, 115, 116},
			nil,
		},
		{"time",
			time.Date(1234, 5, 6, 7, 8, 9, 123456000, time.UTC), // 0xd9, 0x74, 0x46, 0x98, 0xfa, 0xff, 0xff, 0xff, 0x0, 0xca, 0x5b, 0x7
			[]byte{0xd9, 0x74, 0x46, 0x98, 0xfa, 0xff, 0xff, 0xff, 0x0, 0xca, 0x5b, 0x7},
			nil,
		},
		{"byte array",
			[]Message{{WB_EXTERN_DATA, ByteArray, []byte{0x0, 0x1, 0x2}}},
			[]byte{0x10, 0x20, 0x4, 0xe, 0x10, 0x3, 0x0, 0x0, 0x1, 0x2},
			nil,
		},
		{"messages",
			[]Message{{RSCP_AUTHENTICATION, UChar8, uint8(10)}, {RSCP_AUTHENTICATION, UChar8, uint8(10)}},
			[]byte{0x1, 0x0, 0x80, 0x0, 0x3, 0x1, 0x0, 0xa, 0x1, 0x0, 0x80, 0x0, 0x3, 0x1, 0x0, 0xa},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte{})
			err := write(buf, tt.data)
			if (err != nil || tt.wantErr != nil) && !(errors.Is(err, tt.wantErr) || err.Error() == tt.wantErr.Error()) {
				t.Errorf("write() error = %#v, wantErr %#v", err, tt.wantErr)
			}
			if diff := deep.Equal(buf.Bytes(), tt.wantBufBytes); diff != nil {
				t.Errorf("write() = %#v, want %#v\n%s", buf.Bytes(), tt.wantBufBytes, diff)
			}
		})
	}
}

func Test_writeFrame(t *testing.T) {
	tests := []struct {
		name        string
		messages    []Message // go representation of the messages to write
		useChecksum bool
		want        []byte // byte representation of the data written
		wantErr     error
	}{
		{"valid frame with crc",
			[]Message{
				{RSCP_AUTHENTICATION, UChar8, uint8(10)},
			},
			true,
			[]byte{0xe3, 0xdc, 0x0, 0x11, 0xd9, 0x74, 0x46, 0x98, 0xfa, 0xff, 0xff, 0xff, 0x0, 0xca, 0x5b, 0x7, 0x8, 0x0, 0x1, 0x0, 0x80, 0x0, 0x3, 0x1, 0x0, 0xa, 0x7f, 0xb3, 0xef, 0xc3},
			nil,
		},
		{"valid frame without crc",
			[]Message{
				{RSCP_AUTHENTICATION, UChar8, uint8(10)},
			},
			false,
			[]byte{0xe3, 0xdc, 0x0, 0x1, 0xd9, 0x74, 0x46, 0x98, 0xfa, 0xff, 0xff, 0xff, 0x0, 0xca, 0x5b, 0x7, 0x8, 0x0, 0x1, 0x0, 0x80, 0x0, 0x3, 0x1, 0x0, 0xa},
			nil,
		},
	}
	for _, tt := range tests {
		// patch now function
		Now = testTime
		t.Run(tt.name, func(t *testing.T) {
			got, err := writeFrame(tt.messages, tt.useChecksum)
			if (err != nil || tt.wantErr != nil) && !(errors.Is(err, tt.wantErr) || err.Error() == tt.wantErr.Error()) {
				t.Errorf("writeFrame() error = %#v, wantErr %#v", err, tt.wantErr)
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("writeFrame() = %#v, want %#v\n%s", got, tt.want, diff)
			}
		})
	}
}

func TestWrite(t *testing.T) {
	key := createAESKey("testkey")
	initIV := newIV()
	cipherBlock, _ := rijndael256.NewCipher(key[:]) // implementation does not return an error
	tests := []struct {
		name        string
		mode        cipher.BlockMode
		messages    []Message // go representation of the messages to write
		useChecksum bool
		want        []byte // byte representation of the data written
		wantErr     error
	}{

		{"valid frame with crc",
			cipher.NewCBCEncrypter(cipherBlock, initIV[:]),
			[]Message{{RSCP_AUTHENTICATION, UChar8, uint8(10)}},
			true,
			[]byte{0xa5, 0x28, 0x58, 0xa1, 0x26, 0x80, 0xc2, 0xb7, 0x12, 0xbb, 0xc9, 0xcb, 0x8c, 0xe0, 0xf0, 0x40, 0x6f, 0x1, 0xf5, 0xd1, 0xc, 0xa0, 0x6e, 0x9b, 0xda, 0x41, 0xfb, 0x46, 0xb3, 0x11, 0xc2, 0x26},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Write(&tt.mode, tt.messages, tt.useChecksum)
			if (err != nil || tt.wantErr != nil) && !(errors.Is(err, tt.wantErr) || err.Error() == tt.wantErr.Error()) {
				t.Errorf("Write() error = %#v, wantErr %#v", err, tt.wantErr)
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("Write() = %#v, want %#v\n%s", got, tt.want, diff)
			}
		})
	}
}
