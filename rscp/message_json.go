package rscp

import (
	"bytes"
	"encoding/json"
	"fmt"
)

//nolint: funlen
func (m *Message) UnmarshalJSON(b []byte) (err error) {
	switch isArray, isObject, isString := peekJSONType(b); {
	case isString:
		// treat a string as a tag
		if err := json.Unmarshal(b, &m.Tag); err != nil {
			return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
		}
	case isArray:
		// treat an array as a tuple
		tmp := []interface{}{}
		if err := json.Unmarshal(b, &tmp); err != nil {
			return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
		}
		// TODO: there maybe a more efficient way to support optional tuple elements
		//nolint: gomnd
		switch len(tmp) {
		case 1:
			tmp := []interface{}{&m.Tag}
			if err := json.Unmarshal(b, &tmp); err != nil {
				return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
			}
		case 2:
			tmp := []interface{}{&m.Tag, &m.Value}
			if err := json.Unmarshal(b, &tmp); err != nil {
				return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
			}
			// detect unusual case that a data type is specified but no value
			if s, isString := m.Value.(string); isString {
				if dt, err := DataTypeString(s); err == nil {
					m.DataType = dt
					m.Value = nil
				}
			}
		case 3:
			tmp := []interface{}{&m.Tag, &m.DataType, &m.Value}
			if err := json.Unmarshal(b, &tmp); err != nil {
				return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
			}
		}
	case isObject:
		// treat an object as Message, but we use an anonymous struct to not get a stack overflow and copy the fields after unmarshalling
		tmp := JSONMessage{}
		if err := json.Unmarshal(b, &tmp); err != nil {
			return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
		}
		tmp.CopyTo(m)
	default:
		return fmt.Errorf("%w: %s", ErrJSONUnmarshal, "invalid message, expected tuple or object")
	}
	// infer data type (was either not specified or wrong)
	if m.DataType == None {
		m.DataType = m.Tag.DataType()
	}
	// convert number values to expected data type
	if v, isFloat := m.Value.(float64); isFloat {
		if m.Value, err = m.DataType.New(v); err != nil {
			return fmt.Errorf("%w: could not convert number value %f for data type %s", ErrJSONUnmarshal, v, m.DataType)
		}
	}
	if err = m.validate(); err != nil {
		return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
	}
	return nil
}

type JSONMessage struct {
	Tag      Tag         `json:"Tag"`
	DataType DataType    `json:"DataType"`
	Value    interface{} `json:"Value"`
}

type JSONSimpleMessage map[Tag]interface{}

func NewJSONSimpleMessage(message Message) JSONSimpleMessage {
	jm := JSONSimpleMessage{}
	if messages, isContainer := message.Value.([]Message); isContainer {
		jm[message.Tag] = NewJSONSimpleMessages(messages)
	} else {
		jm[message.Tag] = message.Value
	}
	return jm
}

func NewJSONSimpleMessages(messages []Message) []JSONSimpleMessage {
	jms := []JSONSimpleMessage{}
	for _, message := range messages {
		jms = append(jms, NewJSONSimpleMessage(message))
	}
	return jms
}

func (m JSONSimpleMessage) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	length := len(m)
	count := 0
	for key, value := range m {
		count++
		var (
			b   []byte
			vb  []byte
			err error
		)
		if b, err = json.Marshal(key); err != nil {
			return nil, err
		}
		if vb, err = json.Marshal(value); err != nil {
			return nil, err
		}
		buffer.WriteString(fmt.Sprintf("%s:%s", b, vb))
		if count < length {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

func (j *JSONMessage) CopyTo(m *Message) {
	m.Tag, m.DataType, m.Value = j.Tag, j.DataType, j.Value
}
