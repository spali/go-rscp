package rscp

import (
	"errors"
	"testing"
	"time"
)

func Test_validateResponses(t *testing.T) {
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
			ErrNotAResponseTag,
		},
		{"not a response",
			args{[]Message{{INFO_REQ_UTC_TIME, None, nil}}},
			ErrNotAResponseTag,
		},
		{"valid response",
			args{[]Message{{INFO_UTC_TIME, Timestamp, time.Now()}}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateResponses(tt.args.messages)
			if (err != nil || tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("validateResponses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
