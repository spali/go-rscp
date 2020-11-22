package rscp

import (
	"bytes"
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

// peekJSONType peeks the byte slice for basic json types
func peekJSONType(data []byte) (isArray bool, isObject bool, isString bool) {
	x := bytes.TrimLeft(data, " \t\r\n")
	isString = len(x) > 0 && x[0] == '"'
	isArray = len(x) > 0 && x[0] == '['
	isObject = len(x) > 0 && x[0] == '{'
	return
}
