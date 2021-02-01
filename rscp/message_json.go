package rscp

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// unmarshalJSONTuple unmarshalls a tuple into a JSONMessage
func unmarshalJSONTuple(x *jsonMessage, b []byte) error {
	// treat an array as a tuple
	tmpS := []interface{}{}
	if err := json.Unmarshal(b, &tmpS); err != nil {
		return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
	}
	// TODO: there maybe a more efficient way to support optional tuple elements
	//nolint: gomnd
	switch len(tmpS) {
	case 1:
		// tag only tuple
		tmp := []interface{}{&x.Tag}
		if err := json.Unmarshal(b, &tmp); err != nil {
			return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
		}
	case 2:
		// tag, value pair
		tmp := []interface{}{&x.Tag, &x.Value}
		if err := json.Unmarshal(b, &tmp); err != nil {
			return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
		}
		// detect unusual case that a data type is specified but no value
		if _, _, isString := peekJSONType(x.Value); isString {
			dtS := string("")
			if err := json.Unmarshal(x.Value, &dtS); err == nil {
				if dt, err := DataTypeString(dtS); err == nil && dt.IsADataType() {
					x.DataType = dt
					x.Value = nil
				}
			}
		}
	case 3:
		// tuple
		tmp := []interface{}{&x.Tag, &x.DataType, &x.Value}
		if err := json.Unmarshal(b, &tmp); err != nil {
			return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
		}
	}
	return nil
}

func (m *Message) UnmarshalJSON(b []byte) (err error) {
	x := jsonMessage{}
	switch isArray, isObject, isString := peekJSONType(b); {
	case isString:
		// treat a string as a tag
		if err := json.Unmarshal(b, &x.Tag); err != nil {
			return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
		}
	case isArray:
		if err := unmarshalJSONTuple(&x, b); err != nil {
			return err
		}
	case isObject:
		// treat an object as Message, but we do not use Message type to prevent a stack overflow and copy the fields after unmarshalling
		if err := json.Unmarshal(b, &x); err != nil {
			return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
		}
	default:
		return fmt.Errorf("%w: %s", ErrJSONUnmarshal, "invalid request syntax, expected tuple or object")
	}
	// infer data type if not given or just wrong
	x.DataType = x.Tag.DataType()
	if x.DataType == Container {
		tmp := []Message{}
		if err := json.Unmarshal(x.Value, &tmp); err != nil {
			return fmt.Errorf("%w: could not convert value '%s' for data type %s: %s", ErrJSONUnmarshal, x.Value, x.DataType, err)
		}
		m.Value = tmp
	} else if x.Value != nil {
		tmp := reflect.ValueOf(x.DataType.newEmpty(0)).Elem().Interface()
		if err := json.Unmarshal(x.Value, &tmp); err != nil {
			return fmt.Errorf("%w: could not convert value '%s' for data type %s: %s", ErrJSONUnmarshal, x.Value, x.DataType, err)
		}
		// convert number values to expected data type (json does by default unmarshal to float64)
		if v, isFloat := tmp.(float64); isFloat {
			if tmp, err = x.DataType.new(v); err != nil {
				return fmt.Errorf("%w: could not convert number value '%f' for data type %s", ErrJSONUnmarshal, v, x.DataType)
			}
		}
		m.Value = tmp
	}

	m.Tag, m.DataType = x.Tag, x.DataType

	if err = m.validate(); err != nil {
		return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
	}
	return nil
}

type jsonMessage struct {
	Tag      Tag
	DataType DataType
	Value    json.RawMessage
}
