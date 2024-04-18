package rscp

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestMessage_String(t *testing.T) {
	type fields struct {
		Tag      Tag
		DataType DataType
		Value    interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"non secret message",
			fields{RSCP_AUTHENTICATION_USER, CString, "testuser"},
			"{ RSCP_AUTHENTICATION_USER CString testuser }",
		},
		{"secret message",
			fields{RSCP_AUTHENTICATION_PASSWORD, CString, "testpassword"},
			fmt.Sprintf("{ RSCP_AUTHENTICATION_PASSWORD CString %s }", secretReplaceString),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Message{
				Tag:      tt.fields.Tag,
				DataType: tt.fields.DataType,
				Value:    tt.fields.Value,
			}
			if got := m.String(); got != tt.want {
				t.Errorf("Message.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Message_validate(t *testing.T) {
	type args struct {
		message Message
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{"empty message",
			args{Message{}},
			nil,
		},
		{"message of type None",
			args{Message{Tag: INFO_REQ_UTC_TIME}},
			nil,
		},
		{"message with wrong value fir data type",
			args{Message{Tag: INFO_REQ_UTC_TIME, DataType: None, Value: "notvalid"}},
			ErrDataTypeValueMismatch,
		},
		{"valid container message",
			args{Message{Tag: RSCP_REQ_AUTHENTICATION, DataType: Container, Value: []Message{{Tag: RSCP_AUTHENTICATION_USER, DataType: CString, Value: "testuser"}, {Tag: RSCP_AUTHENTICATION_PASSWORD, DataType: CString, Value: "testpassword"}}}},
			nil,
		},
		{"invalid container message",
			args{Message{Tag: RSCP_REQ_AUTHENTICATION, DataType: Container, Value: []Message{{Tag: RSCP_AUTHENTICATION_USER, DataType: CString, Value: nil}, {Tag: RSCP_AUTHENTICATION_PASSWORD, DataType: CString, Value: nil}}}},
			ErrDataTypeValueMismatch,
		},
		{"invalid size",
			args{Message{Tag: RSCP_AUTHENTICATION_USER, DataType: CString, Value: strings.Repeat("x", int(RSCP_DATA_MAX_DATA_SIZE)+1)}},
			ErrRscpDataLimitExceeded,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.args.message.validate(); (err != nil || tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("Message.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
