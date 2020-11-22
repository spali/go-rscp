package rscp

// defines the expected Length of the data type or 0 for variable length
//nolint: gomnd
var lengthMap = map[DataType]uint16{
	None:      0,
	Bool:      1,
	Char8:     1,
	UChar8:    1,
	Int16:     2,
	UInt16:    2,
	Int32:     4,
	Uint32:    4,
	Int64:     8,
	Uint64:    8,
	Float32:   4,
	Double64:  8,
	Bitfield:  1,
	CString:   0,
	Container: 0,
	Timestamp: 12,
	ByteArray: 0,
	Error:     4,
}

// Length returns the data type's expected byte length or 0 for variable length datatype
func (d DataType) Length() uint16 {
	return lengthMap[d]
}
