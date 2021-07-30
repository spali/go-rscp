package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/spali/go-rscp/rscp"
)

type JSONMessage map[rscp.Tag]interface{}

func NewJSONMergedMessages(messages []rscp.Message) JSONMessage {
	jm := JSONMessage{}
	var arrVal []JSONMessage
	for _, message := range messages {
		if messages, isContainer := message.Value.([]rscp.Message); isContainer {
			if _, exists := jm[message.Tag]; exists && arrVal == nil {
				arrVal = []JSONMessage{jm[message.Tag].(JSONMessage)}
			}
			if arrVal != nil {
				arrVal = append(arrVal, NewJSONMergedMessages(messages))
			} else {
				jm[message.Tag] = NewJSONMergedMessages(messages)
			}
			if arrVal != nil {
				jm[message.Tag] = arrVal
			}
		} else {
			jm[message.Tag] = message.Value
		}
	}
	return jm
}

// NewJSONSimpleMessage returns the message as simplified json
func NewJSONSimpleMessage(message rscp.Message) JSONMessage {
	jm := JSONMessage{}
	if messages, isContainer := message.Value.([]rscp.Message); isContainer {
		jm[message.Tag] = NewJSONSimpleMessages(messages)
	} else {
		jm[message.Tag] = message.Value
	}
	return jm
}

// NewJSONSimpleMessages returns the messages as simplified json
func NewJSONSimpleMessages(messages []rscp.Message) []JSONMessage {
	jms := []JSONMessage{}
	for _, message := range messages {
		jms = append(jms, NewJSONSimpleMessage(message))
	}
	return jms
}

// MarshalJSON marshalls the message as json
func (m JSONMessage) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")

	// store the keys sorted to create deterministic json objects and match the default implementation
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = int(k)
		i++
	}
	sort.Ints(keys)

	lastIndex := len(keys) - 1
	for i, key := range keys {
		var (
			b   []byte
			err error
		)
		t := rscp.Tag(key)
		if b, err = json.Marshal(t); err != nil {
			return nil, err
		}
		switch v := m[t].(type) {
		default:
			if v, err = json.Marshal(v); err != nil {
				return nil, err
			}
			buffer.WriteString(fmt.Sprintf("%s:%s", b, v))
		case []uint8:
			// bytearray as json array of numbers
			va := strings.Join(strings.Fields(fmt.Sprintf("%d", v)), ",")
			buffer.WriteString(fmt.Sprintf("%s:%s", b, va))
		}
		if i < lastIndex {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}
