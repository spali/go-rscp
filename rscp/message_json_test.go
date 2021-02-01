package rscp

import (
	"errors"
	"testing"

	"github.com/go-test/deep"
)

func TestMessage_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    *Message
		wantErr error
	}{
		{"empty results in error",
			"",
			nil,
			ErrJSONUnmarshal,
		},
		{`string`,
			`"INFO_REQ_UTC_TIME"`,
			&Message{Tag: INFO_REQ_UTC_TIME},
			nil,
		},
		{`tuple`,
			`["INFO_REQ_UTC_TIME"]`,
			&Message{Tag: INFO_REQ_UTC_TIME},
			nil,
		},
		{`tuple with data type`,
			`["INFO_REQ_UTC_TIME","None"]`,
			&Message{Tag: INFO_REQ_UTC_TIME},
			nil,
		},
		{`tuple with value`,
			`["RSCP_AUTHENTICATION_USER","testuser"]`,
			&Message{Tag: RSCP_AUTHENTICATION_USER, DataType: CString, Value: "testuser"},
			nil,
		},
		{`tuple with data type and value`,
			`["RSCP_AUTHENTICATION_USER","CString","testuser"]`,
			&Message{Tag: RSCP_AUTHENTICATION_USER, DataType: CString, Value: "testuser"},
			nil,
		},
		{`tuple with nested data`,
			`["BAT_REQ_DATA",[["BAT_INDEX",0],"BAT_REQ_DEVICE_STATE"]]`,
			&Message{Tag: BAT_REQ_DATA, DataType: Container, Value: []Message{{Tag: BAT_INDEX, DataType: UInt16, Value: uint16(0)}, {Tag: BAT_REQ_DEVICE_STATE}}},
			nil,
		},
		{`tuple (invalid empty)`,
			`[]`,
			nil,
			ErrJSONUnmarshal,
		},
		{`tuple (invalid)`,
			`[x]`,
			nil,
			ErrJSONUnmarshal,
		},
		{`tuple (invalid tag)`,
			`[1]`,
			nil,
			ErrJSONUnmarshal,
		},
		{`tuple (invalid tag and datatype)`,
			`[1,1]`,
			nil,
			ErrJSONUnmarshal,
		},
		{`tuple (invalid tag, datatype and value)`,
			`[1,1,1]`,
			nil,
			ErrJSONUnmarshal,
		},
		{`tuple (invalid data type) get's fixed`,
			`["INFO_REQ_UTC_TIME", "UChar8"]`,
			&Message{Tag: INFO_REQ_UTC_TIME},
			nil,
		},
		{`tuple (invalid value) get's fixed`,
			`["INFO_REQ_MAC_ADDRESS","None",""]`,
			&Message{Tag: INFO_REQ_MAC_ADDRESS},
			nil,
		},
		{`object`,
			`{ "Tag": "INFO_REQ_UTC_TIME" }`,
			&Message{Tag: INFO_REQ_UTC_TIME},
			nil,
		},
		{`object with value`,
			`{ "Tag": "BAT_INDEX", "Value": 0 }`,
			&Message{Tag: BAT_INDEX, DataType: UInt16, Value: uint16(0)},
			nil,
		},
		{`object all fields`,
			`{ "Tag": "BAT_INDEX", "DataType": "UInt16", "Value": 0 }`,
			&Message{Tag: BAT_INDEX, DataType: UInt16, Value: uint16(0)},
			nil,
		},
		{`object with nested data`,
			`{ "Tag": "BAT_REQ_DATA", "Value": [ { "Tag": "BAT_INDEX", "Value": 0 }, { "Tag": "BAT_REQ_DEVICE_STATE" } ] }`,
			&Message{Tag: BAT_REQ_DATA, DataType: Container, Value: []Message{{Tag: BAT_INDEX, DataType: UInt16, Value: uint16(0)}, {Tag: BAT_REQ_DEVICE_STATE}}},
			nil,
		},
		{`object (invalid tag)`,
			`{ "Tag": 1 }`,
			nil,
			ErrJSONUnmarshal,
		},
		{`string (invalid tag)`,
			`"INVALID_TAG"`,
			nil,
			ErrJSONUnmarshal,
		},
		{`string (invalid data type)`,
			`"RSCP_AUTHENTICATION"`,
			nil,
			ErrJSONUnmarshal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{}
			err := m.UnmarshalJSON([]byte(tt.message))
			if (err != nil || tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("Message.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			} else if err == nil {
				if diff := deep.Equal(m, tt.want); diff != nil {
					t.Errorf("Message.UnmarshalJSON() = %v, want %v\n%s", m, tt.want, diff)
				}
			}
		})
	}
}
