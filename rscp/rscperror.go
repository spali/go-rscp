package rscp

//nolint:revive
type RscpError uint32

// all errors as constant
//
//nolint:revive,stylecheck
//go:generate go run github.com/alvaroloes/enumer -type=RscpError -json
const (
	ERR_NOT_HANDLED    RscpError = 0x01
	ERR_ACCESS_DENIED  RscpError = 0x02
	ERR_FORMAT         RscpError = 0x03
	ERR_AGAIN          RscpError = 0x04
	ERR_OUT_OF_BOUNDS  RscpError = 0x05
	ERR_NOT_AVAILABLE  RscpError = 0x06
	ERR_UNKNOWN_TAG    RscpError = 0x07
	ERR_ALREADY_IN_USE RscpError = 0x08
	// undocumented, but happens for example in get_db_data if time and span is invalid (not available)
	ERR_UNEXPECTED RscpError = 0xFFFFFFFF
)
