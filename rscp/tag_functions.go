package rscp

import (
	"encoding/json"
	"fmt"
)

// TypeFlagBit is the position of the bit indicating if the tag is a request or response
const TypeFlagBit uint8 = 23

// isRequest returns if the tag is request tag
func (t Tag) isRequest() bool {
	return ((t >> TypeFlagBit) & 1) == 0
}

// isResponse returns if the tag is a response tag
func (t Tag) isResponse() bool {
	return ((t >> TypeFlagBit) & 1) == 1
}

// MarshalJSON implements the json.Marshaler interface for Tag
func (t Tag) MarshalJSON() ([]byte, error) {
	if t.IsATag() {
		return json.Marshal(t.String())
	}
	return []byte(fmt.Sprintf("%d", t)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Tag
func (t *Tag) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		var i uint32
		if err := json.Unmarshal(data, &i); err != nil {
			return fmt.Errorf("Tag should be a string or uint32, got %s", data)
		}
		*t = Tag(i)
		return nil
	}

	var err error
	*t, err = TagString(s)
	return err
}
