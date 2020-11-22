package rscp

type DataType uint8

//go:generate enumer -type=DataType -json
const (
	None      DataType = 0x00
	Bool      DataType = 0x01
	Char8     DataType = 0x02
	UChar8    DataType = 0x03
	Int16     DataType = 0x04
	UInt16    DataType = 0x05
	Int32     DataType = 0x06
	Uint32    DataType = 0x07
	Int64     DataType = 0x08
	Uint64    DataType = 0x09
	Float32   DataType = 0x0A
	Double64  DataType = 0x0B
	Bitfield  DataType = 0x0C
	CString   DataType = 0x0D
	Container DataType = 0x0E
	// 64Bit Sekunden + 32Bit Nanosekunden seit 1970
	Timestamp DataType = 0x0F
	ByteArray DataType = 0x10
	Error     DataType = 0xFF
)
