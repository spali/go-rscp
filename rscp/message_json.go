package rscp

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type jsonMessage struct {
	Tag      Tag
	DataType DataType
	Value    json.RawMessage
}

// UnmarshalJSONValue unmarshals a message value from json including nested messages in containers
func (m *Message) UnmarshalJSONValue(jm json.RawMessage) error {
	if m.DataType == Container {
		tmp := []Message{}
		if err := json.Unmarshal(jm, &tmp); err != nil {
			return fmt.Errorf("could not convert value '%s' for data type %s: %s", jm, m.DataType, err)
		}
		m.Value = tmp
	} else if jm != nil && m.DataType != None {
		var tmp interface{}
		// TODO: we need to cleanup generic data type handling somewhen to prevent such hacks
		if m.DataType == Timestamp {
			tmp = m.DataType.newEmpty(0)
		} else {
			tmp = reflect.ValueOf(m.DataType.newEmpty(0)).Elem().Interface()
		}
		if err := json.Unmarshal(jm, &tmp); err != nil {
			return fmt.Errorf("could not convert value '%s' for data type %s: %s", jm, m.DataType, err)
		}
		// convert number values to expected data type (json does by default unmarshal to float64)
		v, isFloat := tmp.(float64)
		if isFloat {
			var err error
			if tmp, err = m.DataType.new(v); err != nil {
				return fmt.Errorf("could not convert number value '%f' for data type %s", v, m.DataType)
			}
		}
		// TODO: we need to cleanup generic data type handling somewhen to prevent such hacks
		if m.DataType == Timestamp {
			m.Value = reflect.ValueOf(tmp).Elem().Interface()
		} else {
			m.Value = tmp
		}
	}
	return nil
}

// UnmarshalJSON unmarshals a message from json including nested messages in containers
func (m *Message) UnmarshalJSON(b []byte) (err error) {
	// treat as Message, but we do not use Message type to prevent a stack overflow and copy the fields after unmarshalling
	jm := jsonMessage{}
	if err := json.Unmarshal(b, &jm); err != nil {
		return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
	}
	// Note: see Issue https://github.com/spali/go-e3dc/issues/1
	if jm.DataType == None && jm.DataType != jm.Tag.DataType() {
		jm.DataType = jm.Tag.DataType()
	}
	m.Tag, m.DataType = jm.Tag, jm.DataType
	if err := m.UnmarshalJSONValue(jm.Value); err != nil {
		return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
	}
	if err = m.validate(); err != nil {
		return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
	}
	return nil
}
