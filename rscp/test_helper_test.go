package rscp

import (
	"time"
)

func pointerOfUint8(i uint8) *uint8        { return &i }
func pointerOfString(s string) *string     { return &s }
func pointerOfTime(t time.Time) *time.Time { return &t }

func testTime() time.Time {
	return time.Date(1234, 5, 6, 7, 8, 9, 123456000, time.UTC)
}
