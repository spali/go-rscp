package rscp

import "errors"

var ErrNoArguments = errors.New("no arguments")
var ErrNotARequestTag = errors.New("tag is not a request tag")
var ErrNotAResponseTag = errors.New("tag is not a response tag")
var ErrTagDataTypeMismatch = errors.New("tag data type does not match data type")
var ErrDataTypeValueMismatch = errors.New("value does not match data type")
var ErrValidTag = errors.New("not a valid tag")
var ErrMissingValue = errors.New("missing value")

var ErrJSONUnmarshal = errors.New("json unmarshal error")

var ErrRscpInvalidMagic = errors.New("ERR_INVALID_MAGIC")
var ErrRscpInvalidControl = errors.New("ERR_INVALID_CONTROL")
var ErrRscpProtVersionMismatch = errors.New("ERR_PROT_VERSION_MISMATCH")
var ErrRscpInvalidFrameLength = errors.New("ERR_INVALID_FRAME_LENGTH")
var ErrRscpInvalidCrc = errors.New("ERR_INVALID_CRC")
var ErrRscpDataLimitExceeded = errors.New("ERR_DATA_LIMIT_EXCEEDED")

//nolint: golint
type RscpError uint32

// String converter function for Error
func (e RscpError) String() string {
	return ErrorMeta[e].Name
}

// all errors as constant
//nolint: golint,stylecheck
const (
	ERR_NOT_HANDLED    RscpError = 0x01
	ERR_ACCESS_DENIED  RscpError = 0x02
	ERR_FORMAT         RscpError = 0x03
	ERR_AGAIN          RscpError = 0x04
	ERR_OUT_OF_BOUNDS  RscpError = 0x05
	ERR_NOT_AVAILABLE  RscpError = 0x06
	ERR_UNKNOWN_TAG    RscpError = 0x07
	ERR_ALREADY_IN_USE RscpError = 0x08
)

type ErrorMetaType struct {
	Name string
}

// ErrorMeta contains meta informations about each error
var ErrorMeta = map[RscpError]ErrorMetaType{
	ERR_NOT_HANDLED:    {"ERR_NOT_HANDLED"},
	ERR_ACCESS_DENIED:  {"ERR_ACCESS_DENIED"},
	ERR_FORMAT:         {"ERR_FORMAT"},
	ERR_AGAIN:          {"ERR_AGAIN"},
	ERR_OUT_OF_BOUNDS:  {"ERR_OUT_OF_BOUNDS"},
	ERR_NOT_AVAILABLE:  {"ERR_NOT_AVAILABLE"},
	ERR_UNKNOWN_TAG:    {"ERR_UNKNOWN_TAG"},
	ERR_ALREADY_IN_USE: {"ERR_ALREADY_IN_USE"},
}
