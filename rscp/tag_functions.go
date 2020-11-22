package rscp

// TypeFlagBit is the position of the bit indicating if the tag is a request or response
const TypeFlagBit uint8 = 23

// isRequest returns if the tag is request tag
func (t Tag) isRequest() bool {
	return ((t >> TypeFlagBit) & 1) == 0
}

// isResponse returns if the tag is a response tag
func (t Tag) isResponse() bool {
	return ((t >> TypeFlagBit) & 1) == 1
}
