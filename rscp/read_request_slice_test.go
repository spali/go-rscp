package rscp

import (
	"errors"
	"testing"

	"github.com/go-test/deep"
	"github.com/spali/go-slicereader"
)

func Test_readRequestSlice(t *testing.T) {
	var zero Message
	tests := []struct {
		name    string
		args    []interface{}
		want    Message
		wantErr error
	}{
		{
			"nil",
			nil,
			zero,
			slicereader.EOS,
		},
		{
			"empty slice",
			[]interface{}{},
			zero,
			slicereader.EOS,
		},
		{
			"not a tag",
			[]interface{}{"not a tag"},
			zero,
			ErrValidTag,
		},
		{
			"simple tag",
			[]interface{}{INFO_REQ_UTC_TIME},
			Message{INFO_REQ_UTC_TIME, INFO_REQ_UTC_TIME.DataType(), nil},
			nil,
		},
		{
			"simple tag with additional message ignored",
			[]interface{}{INFO_REQ_UTC_TIME, INFO_REQ_UTC_TIME},
			Message{INFO_REQ_UTC_TIME, INFO_REQ_UTC_TIME.DataType(), nil},
			nil,
		},
		{
			"tag with value",
			[]interface{}{RSCP_REQ_SET_ENCRYPTION_PASSPHRASE, "test"},
			Message{RSCP_REQ_SET_ENCRYPTION_PASSPHRASE, RSCP_REQ_SET_ENCRYPTION_PASSPHRASE.DataType(), "test"},
			nil,
		},
		{
			"tag with tag as value",
			[]interface{}{RSCP_REQ_SET_ENCRYPTION_PASSPHRASE, INFO_REQ_UTC_TIME},
			zero,
			ErrDataTypeValueMismatch,
		},
		{
			"tag with data type as value",
			[]interface{}{RSCP_REQ_SET_ENCRYPTION_PASSPHRASE, Bool},
			zero,
			ErrDataTypeValueMismatch,
		},
		{
			"tag with value and additional container ignored",
			[]interface{}{INFO_REQ_UTC_TIME, RSCP_REQ_AUTHENTICATION, RSCP_AUTHENTICATION_USER, "user", RSCP_AUTHENTICATION_PASSWORD, "password"},
			Message{INFO_REQ_UTC_TIME, INFO_REQ_UTC_TIME.DataType(), nil},
			nil,
		},
		{
			"tag with missing value",
			[]interface{}{RSCP_REQ_SET_ENCRYPTION_PASSPHRASE},
			zero,
			ErrMissingValue,
		},
		{
			"empty container tag",
			[]interface{}{RSCP_REQ_AUTHENTICATION},
			Message{RSCP_REQ_AUTHENTICATION, RSCP_REQ_AUTHENTICATION.DataType(), []Message{}},
			nil,
		},
		{
			"container tag with values",
			[]interface{}{RSCP_REQ_AUTHENTICATION, RSCP_AUTHENTICATION_USER, "user", RSCP_AUTHENTICATION_PASSWORD, "password"},
			Message{RSCP_REQ_AUTHENTICATION, RSCP_REQ_AUTHENTICATION.DataType(), []Message{{RSCP_AUTHENTICATION_USER, RSCP_AUTHENTICATION_USER.DataType(), "user"}, {RSCP_AUTHENTICATION_PASSWORD, RSCP_AUTHENTICATION_PASSWORD.DataType(), "password"}}},
			nil,
		},
		{
			"unsupported case: container tag with values with additional message ignored",
			[]interface{}{RSCP_REQ_AUTHENTICATION, RSCP_AUTHENTICATION_USER, "user", RSCP_AUTHENTICATION_PASSWORD, "password", INFO_REQ_UTC_TIME},
			Message{RSCP_REQ_AUTHENTICATION, RSCP_REQ_AUTHENTICATION.DataType(), []Message{{RSCP_AUTHENTICATION_USER, RSCP_AUTHENTICATION_USER.DataType(), "user"}, {RSCP_AUTHENTICATION_PASSWORD, RSCP_AUTHENTICATION_PASSWORD.DataType(), "password"}, {INFO_REQ_UTC_TIME, INFO_REQ_UTC_TIME.DataType(), nil}}},
			nil,
		},
		{
			"container tag with missing value",
			[]interface{}{RSCP_REQ_AUTHENTICATION, RSCP_AUTHENTICATION_USER, RSCP_AUTHENTICATION_PASSWORD, "password"},
			zero,
			ErrDataTypeValueMismatch,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readRequestSlice(tt.args)
			if (err != nil || tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("readRequestSlice() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("readRequestSlice() = %v, want %v\n%s", got, tt.want, diff)
			}
		})
	}
}
