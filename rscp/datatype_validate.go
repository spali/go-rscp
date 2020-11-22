package rscp

import (
	"time"
)

//  defines the expected Length of the data type or 0 for variable Length
var validateMap = map[DataType]func(v interface{}) bool{
	None:     func(v interface{}) bool { return v == nil },
	Bool:     func(v interface{}) bool { _, isType := v.(bool); return isType },
	Char8:    func(v interface{}) bool { _, isType := v.(int8); return isType },
	UChar8:   func(v interface{}) bool { _, isType := v.(uint8); return isType },
	Int16:    func(v interface{}) bool { _, isType := v.(int16); return isType },
	UInt16:   func(v interface{}) bool { _, isType := v.(uint16); return isType },
	Int32:    func(v interface{}) bool { _, isType := v.(int32); return isType },
	Uint32:   func(v interface{}) bool { _, isType := v.(uint32); return isType },
	Int64:    func(v interface{}) bool { _, isType := v.(int64); return isType },
	Uint64:   func(v interface{}) bool { _, isType := v.(uint64); return isType },
	Float32:  func(v interface{}) bool { _, isType := v.(float32); return isType },
	Double64: func(v interface{}) bool { _, isType := v.(float64); return isType },
	// TODO: currently no known tag returns this type,
	//       so not sure if this is correct (BAT_REQ_STATUS_CODE and BAT_REQ_ERROR_CODE documented to return it, but thats wrong).
	Bitfield:  func(v interface{}) bool { _, isType := v.(byte); return isType },
	CString:   func(v interface{}) bool { _, isType := v.(string); return isType },
	Container: func(v interface{}) bool { _, isType := v.([]Message); return isType },
	Timestamp: func(v interface{}) bool { _, isType := v.(time.Time); return isType },
	// TODO: could not be tested,
	//       couldn't get a request working for the only known Tag that should response with WB_EXTERN_DATA.
	ByteArray: func(v interface{}) bool { _, isType := v.([]byte); return isType },
	Error:     func(v interface{}) bool { _, isType := v.(RscpError); return isType },
}

// isValidValue validates if the Value matches the expected go type
func (d DataType) isValidValue(v interface{}) bool {
	return validateMap[d](v)
}
