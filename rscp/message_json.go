package rscp

import (
	"encoding/json"
	"fmt"
)

//nolint: funlen,gomnd
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

func NewJSONSimpleMessage(messages []Message) JSONSimpleMessage {
	jm := JSONSimpleMessage{}
	for _, message := range messages {
		if sm, isMessages := message.Value.([]Message); isMessages {
			jm[message.Tag] = NewJSONSimpleMessage(sm)
		} else {
			jm[message.Tag] = message.Value
		}
	}
	return jm
}

func (m JSONSimpleMessage) MarshalJSON() ([]byte, error) {
	mb := []byte("{")
	l := len(m)
	cnt := 0
	for t, v := range m {
		cnt++
		var (
			b   []byte
			vb  []byte
			err error
		)
		if b, err = json.Marshal(t); err != nil {
			return nil, err
		}
		if vb, err = json.Marshal(v); err != nil {
			return nil, err
		}
		mb = append(mb, b...)
		mb = append(mb, []byte(":")...)
		mb = append(mb, vb...)
		if cnt < l {
			mb = append(mb, []byte(",")...)
		}
	}
	mb = append(mb, []byte("}")...)
	return mb, nil
}

func (j *JSONMessage) CopyTo(m *Message) {
	m.Tag, m.DataType, m.Value = j.Tag, j.DataType, j.Value
}
