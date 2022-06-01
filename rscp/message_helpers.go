package rscp

import "fmt"

// messagesSize returns the size of all messages including the headers
func messagesSize(ms []Message) uint16 {
	var l uint16
	for _, m := range ms {
		l += RSCP_DATA_HEADER_SIZE + m.valueSize()
	}
	return l
}

// validateResponses checks the integrity of the response
// each request must contain a valid tag and data type and the data type must match the value
func validateResponses(messages []Message) error {
	for i, m := range messages {
		if err := m.validateResponse(); err != nil {
			return fmt.Errorf("message at index %d: %w", i, err)
		}
	}
	return nil
}
