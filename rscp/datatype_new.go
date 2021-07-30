package rscp

import (
	"fmt"
	"time"

	"github.com/cstockton/go-conv"
)

// returns pointer to new interface of the expected go type
var newEmptyMap = map[DataType]func(size uint16) interface{}{
	None:     func(size uint16) interface{} { return nil },
	Bool:     func(size uint16) interface{} { return new(bool) },
	Char8:    func(size uint16) interface{} { return new(int8) },
	UChar8:   func(size uint16) interface{} { return new(uint8) },
	Int16:    func(size uint16) interface{} { return new(int16) },
	UInt16:   func(size uint16) interface{} { return new(uint16) },
	Int32:    func(size uint16) interface{} { return new(int32) },
	Uint32:   func(size uint16) interface{} { return new(uint32) },
	Int64:    func(size uint16) interface{} { return new(int64) },
	Uint64:   func(size uint16) interface{} { return new(uint64) },
	Float32:  func(size uint16) interface{} { return new(float32) },
	Double64: func(size uint16) interface{} { return new(float64) },
	// TODO: currently no known tag returns this type,
	//       so not sure if this is correct (BAT_REQ_STATUS_CODE and BAT_REQ_ERROR_CODE documented to return it, but thats wrong).
	Bitfield:  func(size uint16) interface{} { return new(byte) },
	CString:   func(size uint16) interface{} { return new(string) },
	Container: func(size uint16) interface{} { return new([]Message) },
	Timestamp: func(size uint16) interface{} { return new(time.Time) },
	// TODO: could not be tested,
	//       couldn't get a request working for the only known Tag that should response with WB_EXTERN_DATA.
	ByteArray: func(size uint16) interface{} { return make([]byte, size) },
	Error:     func(size uint16) interface{} { return new(RscpError) },
}

// newEmpty returns pointer to new interface of the expected go type
func (d DataType) newEmpty(s uint16) interface{} {
	return newEmptyMap[d](s)
}

// returns pointer to new interface of the expected go type with the provided value
var newMap = map[DataType]func(v interface{}) (interface{}, error){
	None:     func(v interface{}) (interface{}, error) { return nil, nil },
	Bool:     func(v interface{}) (interface{}, error) { return conv.Bool(v) },
	Char8:    func(v interface{}) (interface{}, error) { return conv.Uint8(v) },
	UChar8:   func(v interface{}) (interface{}, error) { return conv.Uint8(v) },
	Int16:    func(v interface{}) (interface{}, error) { return conv.Int16(v) },
	UInt16:   func(v interface{}) (interface{}, error) { return conv.Uint16(v) },
	Int32:    func(v interface{}) (interface{}, error) { return conv.Int32(v) },
	Uint32:   func(v interface{}) (interface{}, error) { return conv.Uint32(v) },
	Int64:    func(v interface{}) (interface{}, error) { return conv.Int64(v) },
	Uint64:   func(v interface{}) (interface{}, error) { return conv.Uint64(v) },
	Float32:  func(v interface{}) (interface{}, error) { return conv.Float32(v) },
	Double64: func(v interface{}) (interface{}, error) { return conv.Float64(v) },
	// TODO: currently no known tag returns this type,
	//       so not sure if this is correct (BAT_REQ_STATUS_CODE and BAT_REQ_ERROR_CODE documented to return it, but thats wrong).
	Bitfield: func(v interface{}) (interface{}, error) { b, err := conv.Uint8(v); return b, err },
	CString:  func(v interface{}) (interface{}, error) { return conv.String(v) },
	Container: func(v interface{}) (interface{}, error) {
		if m, ok := v.([]Message); ok {
			return m, nil
		}
		return nil, fmt.Errorf("cannot convert %T to []Message", v)
	},
	Timestamp: func(v interface{}) (interface{}, error) { return conv.Time(v) },
	ByteArray: func(v interface{}) (interface{}, error) { s, err := conv.String(v); return []byte(s), err },
	Error: func(v interface{}) (interface{}, error) {
		if e, ok := v.(RscpError); ok {
			return e, nil
		}
		return nil, fmt.Errorf("cannot convert %T to RscpError", v)
	},
}

// new returns pointer to new interface of the expected go type with the provided value
func (d DataType) new(v interface{}) (interface{}, error) {
	return newMap[d](v)
}
