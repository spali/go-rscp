package rscp

import (
	"fmt"
)

// validateRequest checks the integrity of the request
// must contain a valid tag and data type and the data type must match the value
func validateRequest(message Message) error {
	if !message.Tag.isRequest() {
		return fmt.Errorf("%s: %w", message.Tag, ErrNotARequestTag)
	}
	return message.validate()
}

// ValidateRequests checks the integrity of the requests
// each request must contain a valid tag and data type and the data type must match the value
func validateRequests(messages []Message) error {
	for i, m := range messages {
		if err := validateRequest(m); err != nil {
			return fmt.Errorf("message at index %d: %w", i, err)
		}
	}
	return nil
}

// CreateRequest creates a new message (infer the data type from the tag)
// if the tag's data type is a Container, following tag's will be nested as "sub" requests within the container.
// Every tag that has a data type other than DATATYPE_None requires a following value.
// Examples:
//
//	CreateRequest(INFO_REQ_UTC_TIME)
//	CreateRequest(EMS_REQ_SET_ERROR_BUZZER_ENABLED, true)
//	CreateRequest(BAT_REQ_DATA, BAT_INDEX, uint16(0), BAT_REQ_DEVICE_STATE, BAT_REQ_RSOC, BAT_REQ_STATUS_CODE)
func CreateRequest(values ...interface{}) (msg Message, err error) {
	return readRequestSlice(values)
}

// CreateRequests creates multiple new requests (infer the data type from the tag)
// if the tag's data type is a Container, provided values will be converted to "sub" requests, separated by the provided tag's
// Examples:
//
//	CreateRequests([]interface{}{INFO_REQ_UTC_TIME})
//	CreateRequests([]interface{}{EMS_REQ_SET_ERROR_BUZZER_ENABLED, true})
//	CreateRequests([]interface{}{BAT_REQ_DATA, BAT_INDEX, uint16(0), BAT_REQ_DEVICE_STATE, BAT_REQ_RSOC, BAT_REQ_STATUS_CODE})
func CreateRequests(values ...[]interface{}) ([]Message, error) {
	if len(values) == 0 {
		return nil, ErrNoArguments
	}
	msgs := make([]Message, 0)
	for _, subValues := range values {
		msg, err := CreateRequest(subValues...)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}
	return msgs, nil
}
