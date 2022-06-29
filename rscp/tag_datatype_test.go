package rscp

import (
	"testing"
)

func TestTag_DataTypeDefined(t *testing.T) {
	for _, tag := range _TagValues {
		t.Run(tag.String(), func(t *testing.T) {
			if _, ok := dataTypeMap[tag]; !ok {
				t.Errorf("tag %s has no datatype defined", tag)
			}
		})
	}
}
