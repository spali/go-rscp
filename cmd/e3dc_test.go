package main

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/spali/go-e3dc/rscp"
)

func Test_parseJson(t *testing.T) {

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
			got, err := parseJSON(tt.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("CreateRequests() = %v, want %v\n%s", got, tt.want, diff)
			}
		})
	}
}
