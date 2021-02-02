package rscp

import (
	"errors"
	"testing"

	"github.com/go-test/deep"
)

func TestCreateRequests(t *testing.T) {
	type args struct {
		values [][]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []Message
		wantErr error
	}{

		{"Test empty",
			args{[][]interface{}{}},
			nil,
			ErrNoArguments,
		},
		{"Test Request with nil value",
			args{[][]interface{}{
				{INFO_REQ_UTC_TIME, nil},
			}},
			[]Message{{INFO_REQ_UTC_TIME, None, nil}},
			nil,
		},
		{"Test Request with value",
			args{[][]interface{}{
				{EMS_REQ_SET_ERROR_BUZZER_ENABLED, true},
			}},
			[]Message{{EMS_REQ_SET_ERROR_BUZZER_ENABLED, Bool, true}},
			nil,
		},
		{"Test Container Request with values",
			args{[][]interface{}{
				{RSCP_REQ_AUTHENTICATION,
					RSCP_AUTHENTICATION_USER, "testuser",
					RSCP_AUTHENTICATION_PASSWORD, "testpassword",
				},
			}},
			[]Message{{RSCP_REQ_AUTHENTICATION, Container, []Message{{RSCP_AUTHENTICATION_USER, CString, "testuser"}, {RSCP_AUTHENTICATION_PASSWORD, CString, "testpassword"}}}},
			nil,
		},
		{"Test Container Request with mixed values",
			args{[][]interface{}{
				{BAT_REQ_DATA,
					BAT_INDEX, uint16(0),
					BAT_REQ_DEVICE_STATE,
					BAT_REQ_RSOC,
					BAT_REQ_STATUS_CODE,
				},
			}},
			[]Message{{BAT_REQ_DATA, Container, []Message{{BAT_INDEX, UInt16, uint16(0)}, {BAT_REQ_DEVICE_STATE, None, nil}, {BAT_REQ_RSOC, None, nil}, {BAT_REQ_STATUS_CODE, None, nil}}}},
			nil,
		},
		{"Test Container Request without values",
			args{[][]interface{}{
				{RSCP_REQ_AUTHENTICATION,
					RSCP_AUTHENTICATION_USER,
					RSCP_AUTHENTICATION_PASSWORD,
				},
			}},
			nil,
			ErrDataTypeValueMismatch,
		},
		{"Test Container Request with one value (1)",
			args{[][]interface{}{
				{RSCP_REQ_AUTHENTICATION,
					RSCP_AUTHENTICATION_USER, "testuser",
					RSCP_AUTHENTICATION_PASSWORD,
				},
			}},
			nil,
			ErrMissingValue,
		},
		{"Test Container Request with one value (2)",
			args{[][]interface{}{
				{RSCP_REQ_AUTHENTICATION,
					RSCP_AUTHENTICATION_USER,
					RSCP_AUTHENTICATION_PASSWORD, "testpassword",
				},
			}},
			nil,
			ErrDataTypeValueMismatch,
		},
		{"Test Empty Container Request",
			args{[][]interface{}{
				{RSCP_REQ_AUTHENTICATION},
			}},
			[]Message{{RSCP_REQ_AUTHENTICATION, Container, []Message{}}},
			nil,
		},
		{"Test Container with value instead of tag",
			args{[][]interface{}{
				{RSCP_REQ_AUTHENTICATION, "testvalue"},
			}},
			nil,
			ErrValidTag,
		},
		{"Test multi Request with mixed value",
			args{[][]interface{}{
				{INFO_REQ_UTC_TIME, nil},
				{EMS_REQ_SET_ERROR_BUZZER_ENABLED, true},
			}},
			[]Message{
				{INFO_REQ_UTC_TIME, None, nil},
				{EMS_REQ_SET_ERROR_BUZZER_ENABLED, Bool, true},
			},
			nil,
		},
		{"Test multi Request with containers",
			args{[][]interface{}{
				{INFO_REQ_UTC_TIME, nil},
				{EMS_REQ_SET_ERROR_BUZZER_ENABLED, true},
				{BAT_REQ_DATA,
					BAT_INDEX, uint16(0),
					BAT_REQ_DEVICE_STATE,
					BAT_REQ_RSOC,
					BAT_REQ_STATUS_CODE,
				},
			}},
			[]Message{
				{INFO_REQ_UTC_TIME, None, nil},
				{EMS_REQ_SET_ERROR_BUZZER_ENABLED, Bool, true},
				{BAT_REQ_DATA, Container, []Message{
					{BAT_INDEX, UInt16, uint16(0)},
					{BAT_REQ_DEVICE_STATE, None, nil},
					{BAT_REQ_RSOC, None, nil},
					{BAT_REQ_STATUS_CODE, None, nil},
				}},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateRequests(tt.args.values...)
			if (err != nil || tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("CreateRequests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("CreateRequests() = %v, want %v\n%s", got, tt.want, diff)
			}
		})
	}
}

func TestValidateRequests(t *testing.T) {
	type args struct {
		messages []Message
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{"empty message",
			args{[]Message{{}}},
			ErrValidTag,
		},
		{"not a request",
			args{[]Message{{INFO_UTC_TIME, Timestamp, nil}}},
			ErrNotARequestTag,
		},
		{"valid request",
			args{[]Message{{INFO_REQ_UTC_TIME, None, nil}}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateRequests(tt.args.messages)
			if (err != nil || tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("ValidateRequests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
