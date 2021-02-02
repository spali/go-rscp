// Code generated by "enumer -type=RscpError -json"; DO NOT EDIT.

//
package rscp

import (
	"encoding/json"
	"fmt"
)

const (
	_RscpErrorName_0 = "ERR_NOT_HANDLEDERR_ACCESS_DENIEDERR_FORMATERR_AGAINERR_OUT_OF_BOUNDSERR_NOT_AVAILABLEERR_UNKNOWN_TAGERR_ALREADY_IN_USE"
	_RscpErrorName_1 = "ERR_UNEXPECTED"
)

var (
	_RscpErrorIndex_0 = [...]uint8{0, 15, 32, 42, 51, 68, 85, 100, 118}
	_RscpErrorIndex_1 = [...]uint8{0, 14}
)

func (i RscpError) String() string {
	switch {
	case 1 <= i && i <= 8:
		i -= 1
		return _RscpErrorName_0[_RscpErrorIndex_0[i]:_RscpErrorIndex_0[i+1]]
	case i == 4294967295:
		return _RscpErrorName_1
	default:
		return fmt.Sprintf("RscpError(%d)", i)
	}
}

var _RscpErrorValues = []RscpError{1, 2, 3, 4, 5, 6, 7, 8, 4294967295}

var _RscpErrorNameToValueMap = map[string]RscpError{
	_RscpErrorName_0[0:15]:    1,
	_RscpErrorName_0[15:32]:   2,
	_RscpErrorName_0[32:42]:   3,
	_RscpErrorName_0[42:51]:   4,
	_RscpErrorName_0[51:68]:   5,
	_RscpErrorName_0[68:85]:   6,
	_RscpErrorName_0[85:100]:  7,
	_RscpErrorName_0[100:118]: 8,
	_RscpErrorName_1[0:14]:    4294967295,
}

// RscpErrorString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func RscpErrorString(s string) (RscpError, error) {
	if val, ok := _RscpErrorNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to RscpError values", s)
}

// RscpErrorValues returns all values of the enum
func RscpErrorValues() []RscpError {
	return _RscpErrorValues
}

// IsARscpError returns "true" if the value is listed in the enum definition. "false" otherwise
func (i RscpError) IsARscpError() bool {
	for _, v := range _RscpErrorValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for RscpError
func (i RscpError) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for RscpError
func (i *RscpError) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("RscpError should be a string, got %s", data)
	}

	var err error
	*i, err = RscpErrorString(s)
	return err
}