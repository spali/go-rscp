package rscp

import (
	"errors"
)

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
