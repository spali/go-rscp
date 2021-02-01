package rscp

import (
	"fmt"
)

type Message struct {
	Tag      Tag
	DataType DataType
	Value    interface{}
}

// NewMessage creates a new message (infer the data type from the tag).
func NewMessage(tag Tag, value interface{}) *Message {
	return &Message{Tag: tag, DataType: tag.DataType(), Value: value}
}

const secretReplaceString = "********"

// String converter function for a message
func (m Message) String() string {
	if m.Tag.IsSecret() {
		return fmt.Sprintf("{ %s %s %v }", m.Tag, m.DataType, secretReplaceString)
	}
	return fmt.Sprintf("{ %s %s %v }", m.Tag, m.DataType, m.Value)
}

// valueSize returns actual the value size of a message
func (m *Message) valueSize() uint16 {
	if m.DataType.length() == 0 && m.DataType != None {
		switch v := m.Value.(type) {
		default:
			panic("not implemented")
		case nil:
			return 0
		case string:
			return uint16(len(v))
		case []byte:
			return uint16(len(v))
		case []Message:
			return messagesSize(v)
		}
	} else {
		return m.DataType.length()
	}
}

// isValidDataType check if the data type is the expected one for the message tag
func (m *Message) isValidDataType() bool {
	return m.DataType != m.Tag.DataType()
}

// validateResponse checks the integrity of the response
//
// must contain a valid tag and data type and the data type must match the value
func (m *Message) validateResponse() error {
	if !m.Tag.isResponse() {
		return fmt.Errorf("%s: %w", m.Tag, ErrNotAResponseTag)
	}
	return m.validate()
}

// validate checks the integrity of the message
//
// must contain a valid tag and data type and the data type must match the value
func (m *Message) validate() error {
	// check if message data type matches expected data type of the tag
	if !m.Tag.IsATag() {
		return fmt.Errorf("%s: %w", m.Tag, ErrValidTag)
	} else if m.isValidDataType() {
		return fmt.Errorf("%s expects %s, got %s: %w", m.Tag, m.Tag.DataType(), m.DataType, ErrTagDataTypeMismatch)
	}
	if !m.DataType.isValidValue(m.Value) {
		return fmt.Errorf("expected %T got %T : %w", m.DataType.newEmpty(0), m.Value, ErrDataTypeValueMismatch)
	}
	if m.valueSize() > RSCP_DATA_MAX_DATA_SIZE {
		return ErrRscpDataLimitExceeded
	}
	if m.DataType == Container {
		for i := range m.Value.([]Message) {
			if err := m.Value.([]Message)[i].validate(); err != nil {
				return err
			}
		}
	}
	return nil
}
