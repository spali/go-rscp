package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/spali/go-rscp/rscp"
)

type JSONSimpleMessage map[rscp.Tag]interface{}

func NewJSONSimpleMessage(message rscp.Message) JSONSimpleMessage {
	jm := JSONSimpleMessage{}
	if messages, isContainer := message.Value.([]rscp.Message); isContainer {
		jm[message.Tag] = NewJSONSimpleMessages(messages)
	} else {
		jm[message.Tag] = message.Value
	}
	return jm
}

func NewJSONSimpleMessages(messages []rscp.Message) []JSONSimpleMessage {
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
