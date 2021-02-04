package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/spali/go-rscp/rscp"
)

func TestNewJSONSimpleMessages(t *testing.T) {
	tests := []struct {
		name     string
		messages []rscp.Message
		want     []JSONSimpleMessage
	}{
		{"multi message with a container",
			[]rscp.Message{
				{Tag: rscp.INFO_UTC_TIME, DataType: rscp.Timestamp, Value: time.Date(1234, 5, 6, 7, 8, 9, 123456000, time.UTC)},
				{Tag: rscp.BAT_DATA, DataType: rscp.Container, Value: []rscp.Message{
					{Tag: rscp.BAT_INDEX, DataType: rscp.UInt16, Value: uint16(0)},
				}},
			},
			[]JSONSimpleMessage{
				{rscp.INFO_UTC_TIME: time.Date(1234, 5, 6, 7, 8, 9, 123456000, time.UTC)},
				{rscp.BAT_DATA: []JSONSimpleMessage{
					{rscp.BAT_INDEX: uint16(0)},
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewJSONSimpleMessages(tt.messages)
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("NewJSONSimpleMessages() = %v, want %v\n%s", got, tt.want, diff)
			}
		})
	}
}

func TestJSONSimpleMessage_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		m       JSONSimpleMessage
		want    []byte
		wantErr bool
	}{
		{"container",
			JSONSimpleMessage{rscp.BAT_DATA: []JSONSimpleMessage{
				{rscp.BAT_INDEX: uint16(0)},
				{rscp.BAT_RSOC: float32(0)},
			}},
			[]byte(`{"BAT_DATA":[{"BAT_INDEX":0},{"BAT_RSOC":0}]}`),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONSimpleMessage.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONSimpleMessage.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}
