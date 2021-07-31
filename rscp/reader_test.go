package rscp

import (
	"bytes"
	"crypto/cipher"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/azihsoyn/rijndael256"
	"github.com/go-test/deep"
	log "github.com/sirupsen/logrus"
)

func Test_readHeader(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name          string
		args          args
		wantCrcFlag   bool
		wantFrameSize uint32
		wantDataSize  uint16
		wantErr       error
	}{
		{"valid crc header for 0 data",
			args{[]byte{0xe3, 0xdc, 0x0, 0x11, 0x7b, 0xf3, 0xcc, 0x5f, 0x0, 0x0, 0x0, 0x0, 0xb8, 0x1f, 0xcc, 0x15, 0x0, 0x0}},
			true,
			18 + 0 + 4,
			0,
			nil,
		},
		{"valid non crc header for 0 data",
			args{[]byte{0xe3, 0xdc, 0x0, 0x1, 0x7b, 0xf3, 0xcc, 0x5f, 0x0, 0x0, 0x0, 0x0, 0xb8, 0x1f, 0xcc, 0x15, 0x0, 0x0}},
			false,
			18 + 0 + 0,
			0,
			nil,
		},
		{"valid crc header for 10 data",
			args{[]byte{0xe3, 0xdc, 0x0, 0x11, 0x7b, 0xf3, 0xcc, 0x5f, 0x0, 0x0, 0x0, 0x0, 0xb8, 0x1f, 0xcc, 0x15, 0xa, 0x0}},
			true,
			18 + 10 + 4,
			10,
			nil,
		},
		{"valid non crc header for 10 data",
			args{[]byte{0xe3, 0xdc, 0x0, 0x1, 0x7b, 0xf3, 0xcc, 0x5f, 0x0, 0x0, 0x0, 0x0, 0xb8, 0x1f, 0xcc, 0x15, 0xa, 0x0}},
			false,
			18 + 10 + 0,
			10,
			nil,
		},
		{"invalid magic",
			args{[]byte{0x0, 0xdc, 0x0, 0x11, 0x7b, 0xf3, 0xcc, 0x5f, 0x0, 0x0, 0x0, 0x0, 0xb8, 0x1f, 0xcc, 0x15, 0x0, 0x0}},
			false,
			0,
			0,
			ErrRscpInvalidMagic,
		},
		{"wrong protocol version",
			args{[]byte{0xe3, 0xdc, 0x0, 0x10, 0x7b, 0xf3, 0xcc, 0x5f, 0x0, 0x0, 0x0, 0x0, 0xb8, 0x1f, 0xcc, 0x15, 0x0, 0x0}},
			false,
			0,
			0,
			ErrRscpProtVersionMismatch,
		},
		{"invalid control field",
			args{[]byte{0xe3, 0xdc, 0x1, 0x11, 0x7b, 0xf3, 0xcc, 0x5f, 0x0, 0x0, 0x0, 0x0, 0xb8, 0x1f, 0xcc, 0x15, 0x0, 0x0}},
			false,
			0,
			0,
			ErrRscpInvalidControl,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCrcFlag, gotFrameSize, gotDataSize, err := readHeader(tt.args.data)
			if (err != nil || tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("readHeader() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotCrcFlag != tt.wantCrcFlag {
				t.Errorf("readHeader() gotCrcFlag = %v, want %v", gotCrcFlag, tt.wantCrcFlag)
			}
			if gotFrameSize != tt.wantFrameSize {
				t.Errorf("readHeader() gotFrameSize = %v, want %v", gotFrameSize, tt.wantFrameSize)
			}
			if gotDataSize != tt.wantDataSize {
				t.Errorf("readHeader() gotDataSize = %v, want %v", gotDataSize, tt.wantDataSize)
			}

		})
	}
}

func Test_truncatePadding(t *testing.T) {
	type args struct {
		data      []byte
		frameSize uint32
	}
	tests := []struct {
		name     string
		args     args
		wantData []byte
		wantErr  error
	}{
		{"empty",
			args{
				[]byte{},
				0,
			},
			[]byte{},
			nil,
		},
		{"data after single block (1)",
			args{
				[]byte{0xff},
				0,
			},
			[]byte{0xff},
			ErrRscpInvalidFrameLength,
		},
		{"data after single block (2)",
			args{
				[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff},
				5,
			},
			[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff},
			ErrRscpInvalidFrameLength,
		},
		{"data after single block (3)",
			args{
				[]byte{0xff},
				0,
			},
			[]byte{0xff},
			ErrRscpInvalidFrameLength,
		},
		{"two blocks",
			args{
				[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
				33,
			},
			[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			nil,
		},
		{"data after two blocks ",
			args{
				[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0x0, 0x0},
				32,
			},
			[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff},
			ErrRscpInvalidFrameLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := truncatePadding(&tt.args.data, tt.args.frameSize)
			if (err != nil || tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("truncatePadding() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := deep.Equal(tt.args.data, tt.wantData); diff != nil {
				t.Errorf("truncatePadding() = %v, want %v\n%s", tt.args.data, tt.wantData, diff)
			}
		})
	}
}

func Test_read(t *testing.T) {
	tests := []struct {
		name      string
		bytes     []byte      // byte representation of the data
		emptydata interface{} // go representation of the data (empty input)
		size      uint16      // size of the byte data to read
		want      interface{} // go representation of the data readed
		wantErr   error
	}{
		{"simple type",
			[]byte{255},
			new(uint8),
			1,
			pointerOfUint8(255),
			nil,
		},
		{"simple type error",
			[]byte{255},
			new(uint16),
			2,
			new(uint16),
			io.ErrUnexpectedEOF,
		},
		{"string",
			[]byte{0x74, 0x65, 0x73, 0x74},
			new(string),
			4,
			pointerOfString("test"),
			nil,
		},
		{"string error unexpected EOF",
			[]byte{0x74, 0x65, 0x73, 0x74},
			new(string),
			6,
			new(string),
			io.ErrUnexpectedEOF,
		},
		{"string error EOF",
			[]byte{},
			new(string),
			4,
			new(string),
			io.EOF,
		},
		{"time",
			[]byte{0xd9, 0x74, 0x46, 0x98, 0xfa, 0xff, 0xff, 0xff, 0x0, 0xca, 0x5b, 0x7},
			new(time.Time),
			12,
			pointerOfTime(testTime()), // 0xd9, 0x74, 0x46, 0x98, 0xfa, 0xff, 0xff, 0xff, 0x0, 0xca, 0x5b, 0x7
			nil,
		},
		{"time no ns",
			[]byte{0xd9, 0x74, 0x46, 0x98, 0xfa, 0xff, 0xff, 0xff},
			new(time.Time),
			12,
			new(time.Time),
			io.EOF,
		},
		{"time partial ns",
			[]byte{0xd9, 0x74, 0x46, 0x98, 0xfa, 0xff, 0xff, 0xff, 0x0, 0xca},
			new(time.Time),
			12,
			new(time.Time),
			io.ErrUnexpectedEOF,
		},
		{"time no s",
			[]byte{},
			new(time.Time),
			12,
			new(time.Time),
			io.EOF,
		},
		{"time partial s",
			[]byte{0xd9, 0x74},
			new(time.Time),
			12,
			new(time.Time),
			io.ErrUnexpectedEOF,
		},
		{"messages",
			[]byte{0x1, 0x0, 0x80, 0x0, 0x3, 0x1, 0x0, 0xa, 0x1, 0x0, 0x80, 0x0, 0x3, 0x1, 0x0, 0xa},
			new([]Message),
			16,
			&[]Message{{RSCP_AUTHENTICATION, UChar8, uint8(10)}, {RSCP_AUTHENTICATION, UChar8, uint8(10)}},
			nil,
		},
		{"message EOF due wrong length",
			[]byte{0x1, 0x0, 0x80, 0x0, 0x3, 0x1, 0x0, 0xa},
			new([]Message),
			16,
			&[]Message{{RSCP_AUTHENTICATION, UChar8, uint8(10)}},
			io.EOF,
		},
		{"message EOF due missing data in message",
			[]byte{0x1, 0x0, 0x80, 0x0, 0x3, 0x1, 0x0},
			new([]Message),
			8,
			new([]Message),
			io.EOF,
		},
		{"message unexpected EOF",
			[]byte{0x1, 0x0, 0x80},
			new([]Message),
			16,
			new([]Message),
			io.ErrUnexpectedEOF,
		},
		{"message with wrong length field",
			[]byte{0x1, 0x0, 0x80, 0x0, 0x3, 0x0, 0x0, 0xa},
			new([]Message),
			8,
			new([]Message),
			ErrRscpDataLimitExceeded,
		},
		{"message with length field overflow",
			[]byte{0x1, 0x0, 0x80, 0x0, 0x3, 0xff, 0xff, 0xa},
			new([]Message),
			8,
			new([]Message),
			ErrRscpDataLimitExceeded,
		},
		{"message with unexpected eof length field",
			[]byte{0x1, 0x0, 0x80, 0x0, 0x3, 0xff},
			new([]Message),
			8,
			new([]Message),
			io.ErrUnexpectedEOF,
		},
		// Note: see Issue https://github.com/spali/go-e3dc/issues/1
		//
		// {"message with wrong data type",
		// 	[]byte{0x1, 0x0, 0x80, 0x0, 0x0, 0x1, 0x0, 0xa},
		// 	new([]Message),
		// 	8,
		// 	new([]Message),
		// 	ErrTagDataTypeMismatch,
		// },
		{"message with missing data length field",
			[]byte{0x1, 0x0, 0x80, 0x0},
			new([]Message),
			8,
			new([]Message),
			io.EOF,
		},
		{"message with byte array",
			[]byte{0x10, 0x20, 0x4, 0xe, 0x10, 0x3, 0x0, 0x0, 0x1, 0x2},
			new([]Message),
			10,
			&[]Message{{WB_EXTERN_DATA, ByteArray, []byte{0x0, 0x1, 0x2}}},
			nil,
		},
		{"message with unknown tag log warning",
			[]byte{0x0, 0x0, 0x00, 0x0, 0x3, 0x1, 0x0, 0xa},
			new([]Message),
			8,
			&[]Message{{0, UChar8, uint8(10)}},
			nil,
		},
	}
	for _, tt := range tests {
		// to call debug callback function provided to log.DebugFn at least once
		log.SetLevel(log.DebugLevel)
		t.Run(tt.name, func(t *testing.T) {
			reader := bytes.NewReader(tt.bytes)
			err := read(reader, tt.emptydata, tt.size)
			if (err != nil || tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("read() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := deep.Equal(tt.emptydata, tt.want); diff != nil {
				t.Errorf("read() = %v, want %v\n%s", tt.emptydata, tt.want, diff)
			}
		})
	}
}

func TestRead(t *testing.T) {
	key := createAESKey("testkey")
	initIV := newIV()
	cipherBlock, _ := rijndael256.NewCipher(key[:]) // implementation does not return an error
	type args struct {
		mode      cipher.BlockMode
		buf       *[]byte
		crcFlag   *bool
		frameSize *uint32
		dataSize  *uint16
		data      []byte
	}
	emptyBufPointer := func() *[]byte { buf := make([]byte, 0); return &buf }
	tests := []struct {
		name    string
		args    args
		want    []Message
		wantErr error
	}{
		{"valid frame with crc",
			args{
				cipher.NewCBCDecrypter(cipherBlock, initIV[:]),
				emptyBufPointer(),
				new(bool),
				new(uint32),
				new(uint16),
				[]byte{0xd5, 0xa4, 0x8a, 0xb2, 0x67, 0xd7, 0xda, 0xcb, 0xc3, 0xff, 0x35, 0xf, 0x2d, 0x1f, 0x1a, 0x14, 0x7e, 0x31, 0xe4, 0x63, 0xb5, 0xbb, 0x7b, 0xfc, 0x38, 0xc, 0x59, 0x9, 0x5b, 0x64, 0x26, 0xaf, 0x1b, 0xc4, 0x63, 0x7a, 0x57, 0xd4, 0x2, 0x2e, 0x2, 0x6c, 0xd5, 0x51, 0xd9, 0xb6, 0xf8, 0xf9, 0x17, 0x82, 0x2d, 0x3c, 0x65, 0x3, 0x87, 0xb1, 0xe9, 0x2c, 0x1c, 0xce, 0x66, 0x9b, 0xdf, 0x39, 0xec, 0x4d, 0xf2, 0x9a, 0xa3, 0x5a, 0x4f, 0x3c, 0x13, 0xee, 0xaf, 0x64, 0xc, 0xfa, 0xa0, 0x3, 0xce, 0x56, 0x37, 0xe7, 0xa1, 0x5f, 0xc7, 0x50, 0x16, 0x21, 0xbf, 0xb1, 0x8f, 0xa1, 0xda, 0x5a, 0x1b, 0xf2, 0x1d, 0xd0, 0xef, 0xdd, 0xf4, 0xbc, 0xf4, 0xe7, 0x3c, 0x37, 0x71, 0x85, 0x56, 0xbd, 0x65, 0xdc, 0x85, 0xb8, 0x53, 0x67, 0xb3, 0x4f, 0xac, 0x67, 0x29, 0xf7, 0x9f, 0x41, 0x11, 0xb1, 0x6d, 0x76, 0x6e, 0x60, 0x6f, 0xdc, 0x61, 0x87, 0x6d, 0x80, 0x79, 0x8f, 0xd3, 0xf4, 0x88, 0xb1, 0x7e, 0x5c, 0xfe, 0xf9, 0x28, 0x6b, 0x4f, 0x97, 0xe9, 0xa6, 0xbd, 0xf2, 0xd3, 0x78, 0x51, 0x4e, 0x75, 0x4b, 0xe3, 0x63, 0x18, 0x1c, 0xb4, 0x15, 0xa6, 0xd3, 0x94, 0x68, 0x5d, 0xaf, 0xb4, 0x48, 0x53, 0x5a, 0xe8, 0x70, 0x75, 0xf6, 0xa5, 0x99, 0x40, 0x60, 0x30, 0x5b, 0xed, 0x58, 0x2a, 0xa, 0x15, 0x0, 0x12, 0x4b, 0xbf, 0xb7, 0x76, 0x39, 0x61, 0x4a, 0xee, 0x5f, 0x6a, 0x3f, 0xaf, 0x3e, 0xe0, 0x77, 0xe9, 0x27, 0xaf, 0x7c, 0x27, 0x2b, 0x16, 0x6d, 0x1f, 0x48, 0x3f, 0x8e, 0x6, 0xd5},
			},
			[]Message{
				{BAT_DATA, Container, []Message{
					{BAT_INDEX, UInt16, uint16(0)},
					{BAT_DEVICE_STATE, Container, []Message{
						{BAT_DEVICE_CONNECTED, Bool, true},
						{BAT_DEVICE_WORKING, Bool, true},
						{BAT_DEVICE_IN_SERVICE, Bool, false},
					}},
					{BAT_RSOC, Float32, float32(0)},
					{BAT_STATUS_CODE, Uint32, uint32(0)},
					{BAT_ERROR_CODE, Uint32, uint32(0)},
				}},
				{INFO_TIME, Timestamp, testTime()},
				{EMS_GET_POWER_SETTINGS, Container, []Message{
					{EMS_POWER_LIMITS_USED, Bool, true},
					{EMS_MAX_CHARGE_POWER, Int32, int32(4500)},
					{EMS_MAX_DISCHARGE_POWER, Int32, int32(4500)},
					{EMS_DISCHARGE_START_POWER, Uint32, uint32(65)},
					{EMS_POWERSAVE_ENABLED, Bool, false},
					{EMS_WEATHER_REGULATED_CHARGE_ENABLED, Bool, false},
					{EMS_WEATHER_FORECAST_MODE, Int32, int32(0)},
				}},
			},
			nil,
		},
		{"valid frame without crc (no padding, exact 192 bytes)",
			args{
				cipher.NewCBCDecrypter(cipherBlock, initIV[:]),
				emptyBufPointer(),
				new(bool),
				new(uint32),
				new(uint16),
				[]byte{0x7d, 0x3e, 0x30, 0xab, 0x2f, 0x8d, 0x75, 0x61, 0xb6, 0xf2, 0x4c, 0x2c, 0x7, 0xf4, 0x63, 0x8c, 0x3a, 0xa7, 0xf3, 0x86, 0xf8, 0xcd, 0xf5, 0xe2, 0x1d, 0x29, 0xc6, 0x98, 0x21, 0x43, 0x25, 0xe7, 0xf5, 0x64, 0xf9, 0xc4, 0x8d, 0xea, 0x39, 0xfa, 0x3d, 0xf3, 0xb2, 0x97, 0x9f, 0xbd, 0x3f, 0xb4, 0xe4, 0x20, 0xca, 0xaf, 0x6f, 0xa3, 0x5f, 0x6, 0xd6, 0x8e, 0x5f, 0x8e, 0xca, 0x47, 0xa7, 0x43, 0x16, 0xbb, 0xdd, 0xa, 0x3d, 0x3b, 0xc5, 0xc8, 0x72, 0x55, 0x63, 0x18, 0xd5, 0xe, 0x45, 0xf, 0x35, 0x94, 0xb4, 0x29, 0xa2, 0x2, 0x23, 0x39, 0x9, 0x6d, 0x2f, 0x68, 0xcb, 0x99, 0x56, 0xec, 0x4a, 0xd, 0x2, 0xd5, 0x82, 0x87, 0x29, 0x29, 0x4d, 0x72, 0x68, 0xfd, 0x27, 0xa4, 0x7d, 0x59, 0x5a, 0x38, 0xb9, 0xed, 0x5, 0xbf, 0x89, 0x71, 0xfa, 0xef, 0x42, 0xe5, 0xcf, 0xbf, 0xc2, 0x4e, 0xf3, 0x92, 0xee, 0xbf, 0x6e, 0xac, 0x8f, 0x7, 0x7, 0xcb, 0x35, 0x5b, 0x52, 0xe6, 0xaa, 0xc0, 0xba, 0xec, 0x69, 0xd5, 0xce, 0xf3, 0x13, 0x97, 0x1e, 0x9b, 0x86, 0x4d, 0xca, 0xe4, 0xa4, 0xa7, 0x4a, 0xb, 0xc5, 0xf, 0xc2, 0xd7, 0x4d, 0x8e, 0x3a, 0xf, 0xd6, 0x60, 0xa9, 0x8d, 0x80, 0xb9, 0xbd, 0x27, 0x66, 0xbb, 0xad, 0x8d, 0x76, 0x79, 0x14, 0xd3, 0x2b, 0x4b, 0x15, 0xbd, 0x4f, 0x1b},
			},
			[]Message{
				{BAT_DATA, Container, []Message{
					{BAT_INDEX, UInt16, uint16(0)},
					{BAT_DEVICE_STATE, Container, []Message{
						{BAT_DEVICE_CONNECTED, Bool, true},
						{BAT_DEVICE_WORKING, Bool, true},
						{BAT_DEVICE_IN_SERVICE, Bool, false},
					}},
					{BAT_RSOC, Float32, float32(0)},
					{BAT_STATUS_CODE, Uint32, uint32(0)},
					{BAT_ERROR_CODE, Uint32, uint32(0)},
				}},
				{INFO_TIME, Timestamp, testTime()},
				{EMS_GET_POWER_SETTINGS, Container, []Message{
					{EMS_POWER_LIMITS_USED, Bool, true},
					{EMS_MAX_CHARGE_POWER, Int32, int32(4500)},
					{EMS_MAX_DISCHARGE_POWER, Int32, int32(4500)},
					{EMS_DISCHARGE_START_POWER, Uint32, uint32(65)},
					{EMS_POWERSAVE_ENABLED, Bool, false},
					{EMS_WEATHER_REGULATED_CHARGE_ENABLED, Bool, false},
					{EMS_WEATHER_FORECAST_MODE, Int32, int32(0)},
				}},
			},
			nil,
		},
		{"empty read",
			args{
				cipher.NewCBCDecrypter(cipherBlock, initIV[:]),
				emptyBufPointer(),
				new(bool),
				new(uint32),
				new(uint16),
				[]byte{},
			},
			nil,
			ErrRscpInvalidFrameLength,
		},
		{"valid frame with crc",
			args{
				cipher.NewCBCDecrypter(cipherBlock, initIV[:]),
				emptyBufPointer(),
				new(bool),
				new(uint32),
				new(uint16),
				[]byte{0xe4, 0xba, 0xf9, 0x68, 0x0, 0x95, 0x3a, 0x1d, 0x70, 0xea, 0x61, 0x24, 0xa3, 0xa3, 0x1d, 0x75, 0xf6, 0xa0, 0x94, 0x2c, 0x38, 0xe5, 0xcc, 0x21, 0xd8, 0x81, 0x6, 0x7c, 0x45, 0xcf, 0x26, 0x2d},
			},
			[]Message{{RSCP_AUTHENTICATION, UChar8, uint8(10)}},
			nil,
		},
		{"invalid header",
			args{
				cipher.NewCBCDecrypter(cipherBlock, initIV[:]),
				emptyBufPointer(),
				new(bool),
				new(uint32),
				new(uint16),
				[]byte{0xe4, 0x85, 0xf0, 0x39, 0x62, 0x36, 0x33, 0x0, 0xba, 0x26, 0x77, 0x24, 0xb2, 0x95, 0x3e, 0xcd, 0x75, 0xa7, 0x53, 0x18, 0x2, 0x7a, 0x84, 0x86, 0xcb, 0x6d, 0xc6, 0xf0, 0xe0, 0xd, 0xad, 0x54},
			},
			nil,
			ErrRscpInvalidMagic,
		},
		{"unexpected data in padding",
			args{
				cipher.NewCBCDecrypter(cipherBlock, initIV[:]),
				emptyBufPointer(),
				new(bool),
				new(uint32),
				new(uint16),
				[]byte{0xb5, 0xe5, 0x18, 0x7f, 0xeb, 0x7f, 0xf4, 0x91, 0xbb, 0x30, 0xd1, 0x5b, 0x27, 0x83, 0xa4, 0x30, 0xab, 0xcb, 0xc7, 0xfd, 0x9b, 0x7f, 0x89, 0xcc, 0x29, 0x87, 0xe3, 0x22, 0x63, 0x5, 0xd6, 0x1c},
			},
			nil,
			ErrRscpInvalidFrameLength,
		},
		{"missing crc",
			args{
				cipher.NewCBCDecrypter(cipherBlock, initIV[:]),
				emptyBufPointer(),
				new(bool),
				new(uint32),
				new(uint16),
				[]byte{0x67, 0xdb, 0x11, 0x5e, 0xe9, 0x2c, 0xb4, 0x18, 0x52, 0x57, 0x53, 0xd0, 0x9e, 0xc8, 0x3, 0x55, 0x6a, 0x52, 0xc6, 0x35, 0x14, 0xfa, 0x90, 0x8b, 0x12, 0xf6, 0x3d, 0x4e, 0x46, 0x73, 0xd0, 0xbd},
			},
			nil,
			ErrRscpInvalidCrc,
		},
		{"frame/data data length mismatch",
			args{
				cipher.NewCBCDecrypter(cipherBlock, initIV[:]),
				emptyBufPointer(),
				new(bool),
				new(uint32),
				new(uint16),
				[]byte{0x29, 0x9c, 0xd8, 0xc6, 0x46, 0xc0, 0x8d, 0x3f, 0x99, 0x29, 0x6, 0x79, 0x70, 0xa3, 0xf0, 0x65, 0x53, 0x2c, 0xa5, 0x74, 0xad, 0xdf, 0xa3, 0xb8, 0xcb, 0x12, 0x75, 0x7d, 0x90, 0xde, 0xf4, 0x74},
			},
			nil,
			ErrRscpInvalidFrameLength,
		},
		// Note: see Issue https://github.com/spali/go-e3dc/issues/1
		//
		// {"wrong data",
		// 	args{
		// 		cipher.NewCBCDecrypter(cipherBlock, initIV[:]),
		// 		emptyBufPointer(),
		// 		new(bool),
		// 		new(uint32),
		// 		new(uint16),
		// 		[]byte{0xf7, 0xe7, 0x2d, 0x81, 0xd1, 0x4f, 0x5f, 0x9b, 0x37, 0x82, 0x7, 0x92, 0x4c, 0x60, 0x3b, 0x17, 0xda, 0xed, 0xea, 0xb0, 0x63, 0xda, 0x1a, 0xc4, 0xc0, 0xa1, 0x33, 0x26, 0x1, 0xb9, 0xe3, 0x1f},
		// 	},
		// 	nil,
		// 	ErrTagDataTypeMismatch,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read(&tt.args.mode, tt.args.buf, tt.args.crcFlag, tt.args.frameSize, tt.args.dataSize, tt.args.data)
			if (err != nil || tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("Read() = %v, want %v\n%s", got, tt.want, diff)
			}
		})
	}
}
