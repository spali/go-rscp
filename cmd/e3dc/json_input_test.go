package main

import (
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/spali/go-rscp/rscp"
)

func Test_unmarshalJSONRequest(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    rscp.Message
		wantErr bool
	}{
		{`tuple with data type and value`,
			`["RSCP_AUTHENTICATION_USER","CString","testuser"]`,
			rscp.Message{Tag: rscp.RSCP_AUTHENTICATION_USER, DataType: rscp.CString, Value: "testuser"},
			false,
		},
		{"empty results in error",
			"",
			rscp.Message{},
			true,
		},
		{`string`,
			`"INFO_REQ_UTC_TIME"`,
			rscp.Message{Tag: rscp.INFO_REQ_UTC_TIME},
			false,
		},
		{`uint32 (valid known uint32 tag)`,
			`1`,
			rscp.Message{Tag: 1, DataType: rscp.Container},
			false,
		},
		{`uint32 (valid unknown uint32 tag)`,
			`4294967295`,
			rscp.Message{Tag: 4294967295},
			false,
		},
		{`string (invalid tag)`,
			`"INVALID_TAG"`,
			rscp.Message{},
			true,
		},
		{`tuple`,
			`["INFO_REQ_UTC_TIME"]`,
			rscp.Message{Tag: rscp.INFO_REQ_UTC_TIME},
			false,
		},
		{`tuple with data type`,
			`["INFO_REQ_UTC_TIME","None"]`,
			rscp.Message{Tag: rscp.INFO_REQ_UTC_TIME},
			false,
		},
		{`tuple with value`,
			`["RSCP_AUTHENTICATION_USER","testuser"]`,
			rscp.Message{Tag: rscp.RSCP_AUTHENTICATION_USER, DataType: rscp.CString, Value: "testuser"},
			false,
		},
		{`tuple with data type and value`,
			`["RSCP_AUTHENTICATION_USER","CString","testuser"]`,
			rscp.Message{Tag: rscp.RSCP_AUTHENTICATION_USER, DataType: rscp.CString, Value: "testuser"},
			false,
		},
		{`tuple with nested data`,
			`["BAT_REQ_DATA",[["BAT_INDEX",0],"BAT_REQ_DEVICE_STATE"]]`,
			rscp.Message{Tag: rscp.BAT_REQ_DATA, DataType: rscp.Container, Value: []rscp.Message{{Tag: rscp.BAT_INDEX, DataType: rscp.UInt16, Value: uint16(0)}, {Tag: rscp.BAT_REQ_DEVICE_STATE}}},
			false,
		},
		{`tuple (invalid empty)`,
			`[]`,
			rscp.Message{},
			true,
		},
		{`tuple (invalid)`,
			`[x]`,
			rscp.Message{},
			true,
		},
		{`tuple (valid tag with datatype override)`,
			`[1, "None", null ]`,
			rscp.Message{Tag: rscp.RSCP_REQ_AUTHENTICATION, DataType: rscp.None, Value: nil},
			false,
		},
		{`tuple (valid known uint32 tag with datatype override)`,
			`[ 1, "None", null ]`,
			rscp.Message{Tag: 1, DataType: rscp.None, Value: nil},
			false,
		},
		{`tuple (valid unknown uint32 tag)`,
			`[ 4294967295, "None", null ]`,
			rscp.Message{Tag: 4294967295, DataType: rscp.None, Value: nil},
			false,
		},
		{`tuple (invalid string tag)`,
			`[ "INVALID_TAG" ]`,
			rscp.Message{},
			true,
		},
		{`tuple (override data type) allowed`,
			`["INFO_REQ_UTC_TIME", "UChar8"]`,
			rscp.Message{Tag: rscp.INFO_REQ_UTC_TIME, DataType: rscp.UChar8},
			false,
		},
		{`tuple (override data type) allowed`,
			`["INFO_REQ_UTC_TIME", "UChar8", 0]`,
			rscp.Message{Tag: rscp.INFO_REQ_UTC_TIME, DataType: rscp.UChar8, Value: uint8(0)},
			false,
		},
		{`tuple (invalid value) get's fixed`,
			`["INFO_REQ_MAC_ADDRESS","None",""]`,
			rscp.Message{Tag: rscp.INFO_REQ_MAC_ADDRESS},
			false,
		},
		{`object`,
			`{ "Tag": "INFO_REQ_UTC_TIME" }`,
			rscp.Message{Tag: rscp.INFO_REQ_UTC_TIME},
			false,
		},
		{`object with value`,
			`{ "Tag": "BAT_INDEX", "Value": 0 }`,
			rscp.Message{Tag: rscp.BAT_INDEX, DataType: rscp.UInt16, Value: uint16(0)},
			false,
		},
		{`object all fields`,
			`{ "Tag": "BAT_INDEX", "DataType": "UInt16", "Value": 0 }`,
			rscp.Message{Tag: rscp.BAT_INDEX, DataType: rscp.UInt16, Value: uint16(0)},
			false,
		},
		{`object with nested data`,
			`{ "Tag": "BAT_REQ_DATA", "Value": [ { "Tag": "BAT_INDEX", "Value": 0 }, { "Tag": "BAT_REQ_DEVICE_STATE" } ] }`,
			rscp.Message{Tag: rscp.BAT_REQ_DATA, DataType: rscp.Container, Value: []rscp.Message{{Tag: rscp.BAT_INDEX, DataType: rscp.UInt16, Value: uint16(0)}, {Tag: rscp.BAT_REQ_DEVICE_STATE}}},
			false,
		},
		{`object (valid tag with datatype override)`,
			`{ "Tag": 1, "DataType": "None", "Value": null }`,
			rscp.Message{Tag: rscp.RSCP_REQ_AUTHENTICATION, DataType: rscp.None, Value: nil},
			false,
		},
		{`object (valid known uint32 tag with datatype override)`,
			`{ "Tag": 1, "DataType": "None", "Value": null }`,
			rscp.Message{Tag: 1, DataType: rscp.None, Value: nil},
			false,
		},
		{`object (valid unknown uint32 tag)`,
			`{ "Tag": 4294967295, "DataType": "None", "Value": null }`,
			rscp.Message{Tag: 4294967295, DataType: rscp.None, Value: nil},
			false,
		},
		{`object (invalid string tag)`,
			`{ "Tag": "INVALID_TAG" }`,
			rscp.Message{},
			true,
		},
		{`time value`,
			`["INFO_SET_TIME","1234-05-06T07:08:09.123456Z"]`,
			rscp.Message{Tag: rscp.INFO_SET_TIME, DataType: rscp.Timestamp, Value: time.Date(1234, 5, 6, 7, 8, 9, 123456000, time.UTC)},
			false,
		},
		{`bytearray value`,
			`["WB_EXTERN_DATA",[0,6,0,0,0,0]]`,
			rscp.Message{Tag: rscp.WB_EXTERN_DATA, DataType: rscp.ByteArray, Value: []byte{0x0, 0x6, 0x0, 0x0, 0x0, 0x0}},
			false,
		},
		{`empty bytearray value`,
			`["WB_EXTERN_DATA",[]]`,
			rscp.Message{Tag: rscp.WB_EXTERN_DATA, DataType: rscp.ByteArray, Value: []byte{}},
			false,
		},
		{`bytearray invalid number`,
			`["WB_EXTERN_DATA",1]`,
			rscp.Message{Tag: rscp.WB_EXTERN_DATA, DataType: rscp.ByteArray, Value: nil},
			true,
		},
		{`bytearray invalid string`,
			`["WB_EXTERN_DATA","S"]`,
			rscp.Message{Tag: rscp.WB_EXTERN_DATA, DataType: rscp.ByteArray, Value: nil},
			true,
		},
		{`bytearray invalid string array`,
			`["WB_EXTERN_DATA",["S"]]`,
			rscp.Message{Tag: rscp.WB_EXTERN_DATA, DataType: rscp.ByteArray, Value: nil},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := rscp.Message{}
			err := unmarshalJSONRequest([]byte(tt.message), &m)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalJSONRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(m, tt.want); diff != nil {
				t.Errorf("unmarshalJSONRequest() = %v, want %v\n%s", m, tt.want, diff)
			}
		})
	}
}

func Test_unmarshalJSONRequests(t *testing.T) {

	tests := []struct {
		name    string
		message string
		want    []rscp.Message
		wantErr bool
	}{
		{"empty results in error",
			"",
			nil,
			true,
		},
		{`array of tags`,
			`["INFO_REQ_MAC_ADDRESS","INFO_REQ_UTC_TIME"]`,
			[]rscp.Message{{Tag: rscp.INFO_REQ_MAC_ADDRESS}, {Tag: rscp.INFO_REQ_UTC_TIME}},
			false,
		},
		{`array of tuples`,
			`[["INFO_REQ_MAC_ADDRESS"],["INFO_REQ_UTC_TIME"]]`,
			[]rscp.Message{{Tag: rscp.INFO_REQ_MAC_ADDRESS}, {Tag: rscp.INFO_REQ_UTC_TIME}},
			false,
		},
		{`array of messages`,
			`[{ "Tag": "INFO_REQ_MAC_ADDRESS" }, { "Tag": "INFO_REQ_UTC_TIME" }]`,
			[]rscp.Message{{Tag: rscp.INFO_REQ_MAC_ADDRESS}, {Tag: rscp.INFO_REQ_UTC_TIME}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unmarshalJSONRequests([]byte(tt.message))
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalJSONRequests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("unmarshalJSONRequests() = %v, want %v\n%s", got, tt.want, diff)
			}
		})
	}
}
