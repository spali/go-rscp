package main

import (
	"encoding/json"
	"errors"

	"github.com/spali/go-rscp/rscp"
)

var (
	ErrInputNotAnArray   = errors.New("request input has always to be an in an array")
	ErrInputInvalidTuple = errors.New("request contains an invalid tuple")
)

func unmarshalJSONValue(b []byte, m *rscp.Message) error {
	if m.DataType == rscp.Container {
		var err error
		if m.Value, err = unmarshalJSONRequests(b); err != nil {
			return err
		}
		return nil
	}
	if err := m.UnmarshalJSONValue(b); err != nil {
		return err
	}
	return nil
}

//nolint:gomnd
func unmarshalJSONRequest(b []byte, m *rscp.Message) error {
	if isJSONEmpty(b) {
		return ErrInputInvalidTuple
	}
	if isJSONArray(b) {
		// parse tuple
		t := []json.RawMessage{}
		if err := json.Unmarshal(b, &t); err != nil {
			return err
		}
		l := len(t)
		if l < 1 || l > 3 {
			return ErrInputInvalidTuple
		}
		// parse tag
		if err := json.Unmarshal(t[0], &(m.Tag)); err != nil {
			return err
		}
		if l > 1 {
			// check for data type in second element
			if isDataType, dt := isJSONDataType(t[1]); isDataType {
				m.DataType = *dt
			} else {
				// infer data type
				m.DataType = m.Tag.DataType()
				// unmarshal value
				if err := unmarshalJSONValue(t[1], m); err != nil {
					return err
				}
			}
		} else {
			// infer data type
			m.DataType = m.Tag.DataType()
		}
		if l == 3 {
			// unmarshal value
			if err := unmarshalJSONValue(t[2], m); err != nil {
				return err
			}
		}
		return nil
	}
	if isJSONString(b) {
		// parse tag
		if err := json.Unmarshal(b, &(m.Tag)); err != nil {
			return err
		}
		// infer data type
		m.DataType = m.Tag.DataType()
		return nil
	}
	// use Message default Unmarshal
	if err := json.Unmarshal(b, m); err != nil {
		return err
	}
	return nil
}

func unmarshalJSONRequests(b []byte) ([]rscp.Message, error) {
	if !isJSONArray(b) {
		return nil, ErrInputNotAnArray
	}
	r := []json.RawMessage{}
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, err
	}
	m := make([]rscp.Message, len(r))
	for i, v := range r {
		if err := unmarshalJSONRequest(v, &m[i]); err != nil {
			return nil, err
		}
	}
	return m, nil
}
