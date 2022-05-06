package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/spali/go-rscp/rscp"
)

func TestNewJSONMergedMessages(t *testing.T) {
	tests := []struct {
		name     string
		messages []rscp.Message
		want     JSONMessage
	}{
		{"multi message with a container",
			[]rscp.Message{
				{Tag: rscp.INFO_UTC_TIME, DataType: rscp.Timestamp, Value: time.Date(1234, 5, 6, 7, 8, 9, 123456000, time.UTC)},
				{Tag: rscp.BAT_DATA, DataType: rscp.Container, Value: []rscp.Message{
					{Tag: rscp.BAT_INDEX, DataType: rscp.UInt16, Value: uint16(0)},
				}},
			},
			JSONMessage{
				rscp.INFO_UTC_TIME: time.Date(1234, 5, 6, 7, 8, 9, 123456000, time.UTC),
				rscp.BAT_DATA: JSONMessage{
					rscp.BAT_INDEX: uint16(0),
				},
			},
		},
		{"multi message with a container containing duplicate keys",
			[]rscp.Message{
				{Tag: rscp.INFO_UTC_TIME, DataType: rscp.Timestamp, Value: time.Date(1234, 5, 6, 7, 8, 9, 123456000, time.UTC)},
				{Tag: rscp.EMS_GET_SYS_SPECS, DataType: rscp.Container, Value: []rscp.Message{
					{Tag: rscp.EMS_SYS_SPEC, DataType: rscp.Container, Value: []rscp.Message{
						{Tag: rscp.EMS_SYS_SPEC_INDEX, DataType: rscp.Int32, Value: int32(0)},
						{Tag: rscp.EMS_SYS_SPEC_NAME, DataType: rscp.CString, Value: "maxChargePower"},
						{Tag: rscp.EMS_SYS_SPEC_VALUE_INT, DataType: rscp.Int32, Value: int32(4500)},
					}},
					{Tag: rscp.EMS_SYS_SPEC, DataType: rscp.Container, Value: []rscp.Message{
						{Tag: rscp.EMS_SYS_SPEC_INDEX, DataType: rscp.Int32, Value: int32(1)},
						{Tag: rscp.EMS_SYS_SPEC_NAME, DataType: rscp.CString, Value: "maxDischargePower"},
						{Tag: rscp.EMS_SYS_SPEC_VALUE_INT, DataType: rscp.Int32, Value: int32(4500)},
					}},
				}},
			},
			JSONMessage{
				rscp.INFO_UTC_TIME: time.Date(1234, 5, 6, 7, 8, 9, 123456000, time.UTC),
				rscp.EMS_GET_SYS_SPECS: JSONMessage{
					rscp.EMS_SYS_SPEC: []JSONMessage{
						{
							rscp.EMS_SYS_SPEC_INDEX:     int32(0),
							rscp.EMS_SYS_SPEC_NAME:      "maxChargePower",
							rscp.EMS_SYS_SPEC_VALUE_INT: int32(4500),
						},
						{
							rscp.EMS_SYS_SPEC_INDEX:     int32(1),
							rscp.EMS_SYS_SPEC_NAME:      "maxDischargePower",
							rscp.EMS_SYS_SPEC_VALUE_INT: int32(4500),
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewJSONMergedMessages(tt.messages)
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("NewJSONMergedMessages() = %v, want %v\n%s", got, tt.want, diff)
			}
		})
	}
}

func TestNewJSONSimpleMessages(t *testing.T) {
	tests := []struct {
		name     string
		messages []rscp.Message
		want     []JSONMessage
	}{
		{"multi message with a container",
			[]rscp.Message{
				{Tag: rscp.INFO_UTC_TIME, DataType: rscp.Timestamp, Value: time.Date(1234, 5, 6, 7, 8, 9, 123456000, time.UTC)},
				{Tag: rscp.BAT_DATA, DataType: rscp.Container, Value: []rscp.Message{
					{Tag: rscp.BAT_INDEX, DataType: rscp.UInt16, Value: uint16(0)},
				}},
			},
			[]JSONMessage{
				{rscp.INFO_UTC_TIME: time.Date(1234, 5, 6, 7, 8, 9, 123456000, time.UTC)},
				{rscp.BAT_DATA: []JSONMessage{
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
		m       JSONMessage
		want    []byte
		wantErr bool
	}{
		{"container simple",
			JSONMessage{rscp.BAT_DATA: []JSONMessage{
				{rscp.BAT_INDEX: uint16(0)},
				{rscp.BAT_RSOC: float32(0)},
			}},
			[]byte(`{"BAT_DATA":[{"BAT_INDEX":0},{"BAT_RSOC":0}]}`),
			false,
		},
		{"container merged sorted",
			JSONMessage{rscp.BAT_DATA: JSONMessage{
				rscp.BAT_INDEX: uint16(0),
				rscp.BAT_RSOC:  float32(0),
			}},
			[]byte(`{"BAT_DATA":{"BAT_INDEX":0,"BAT_RSOC":0}}`),
			false,
		},
		// required to test the sort of keys and be sure to return deterministic objects
		{"container merged unsorted",
			JSONMessage{rscp.BAT_DATA: JSONMessage{
				rscp.BAT_RSOC:  float32(0),
				rscp.BAT_INDEX: uint16(0),
			}},
			[]byte(`{"BAT_DATA":{"BAT_INDEX":0,"BAT_RSOC":0}}`),
			false,
		},
		{"byte array as array bot as string",
			JSONMessage{
				rscp.WB_EXTERN_DATA: []uint8{0, 0, 6, 0, 0, 0, 0, 0},
			},
			[]byte(`{"WB_EXTERN_DATA":[0,0,6,0,0,0,0,0]}`),
			false,
		},
		{"valid time (test workaround for https://github.com/spali/go-rscp/issues/17)",
			JSONMessage{rscp.EMS_MANUAL_CHARGE_LASTSTART: time.Date(1234, 5, 6, 7, 8, 9, 123456000, time.UTC)},
			[]byte(`{"EMS_MANUAL_CHARGE_LASTSTART":"1234-05-06T07:08:09.123456Z"}`),
			false,
		},
		{"invalid negative time (test workaround for https://github.com/spali/go-rscp/issues/17)",
			JSONMessage{rscp.EMS_MANUAL_CHARGE_LASTSTART: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)},
			[]byte(`{"EMS_MANUAL_CHARGE_LASTSTART":"1970-01-01T00:00:00Z"}`),
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
