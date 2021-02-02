package rscp

import (
	"reflect"
)

// dereferencePtr dereferences a interface containing a pointer to a value.
//
// returns the value of the contained pointer the interface v contains
// or the containing value if the interface v does not contain a value.
func dereferencePtr(v interface{}) interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		return vv.Elem().Interface()
	}
	return v
}
