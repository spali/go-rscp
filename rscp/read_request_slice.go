package rscp

import (
	"fmt"

	"github.com/spali/go-slicereader"
)

// readRequestSlice reads a single request message (recursive for containers) from a slice.
func readRequestSlice(values []interface{}) (Message, error) {
	sr := slicereader.NewSliceReader(values)
	return readRequestSliceReader(sr)
}

// readRequestSliceReader reads a single request message (recursive for containers) from the SliceReader.
func readRequestSliceReader(sr *slicereader.SliceReader) (Message, error) {
	var (
		t     interface{}
		tag   Tag
		isTag bool
		err   error
	)
	if t, err = sr.Read(); err != nil {
		return Message{}, err
	}
	if tag, isTag = t.(Tag); !isTag {
		return Message{}, fmt.Errorf("element at index %d: %w", int(sr.Size()-int64(sr.Len())-1), ErrValidTag)
	}
	msg := NewMessage(tag, nil)
	if msg.DataType == None {
		return msg, nil
	}
	if msg.DataType == Container {
		msg.Value = make([]Message, 0)
		for sr.Len() > 0 {
			subMsg, err := readRequestSliceReader(sr)
			if err != nil {
				return Message{}, err
			}
			msg.Value = append(msg.Value.([]Message), subMsg)
		}
	} else {
		if msg.Value, err = sr.Read(); err != nil {
			return Message{}, fmt.Errorf("expect value after tag %s with data type %s: %w", msg.Tag, msg.DataType, ErrMissingValue)
		}
		switch msg.Value.(type) {
		case Tag, DataType:
			return Message{}, fmt.Errorf("expect element at index %d to be a value for tag %s, got type: %T: %w",
				int(sr.Size()-int64(sr.Len())-1), msg.Tag, msg.Value, ErrDataTypeValueMismatch)
		}
	}
	return msg, nil
}
