package rscp

//nolint:revive,stylecheck
const (
	RSCP_FRAME_MAGIC_SIZE  uint16 = 0x2 // magic field size of frame
	RSCP_FRAME_CTRL_SIZE   uint16 = 0x2 // control field size of frame
	RSCP_FRAME_TIME_SIZE   uint16 = 0xc // time field size of frame
	RSCP_FRAME_LENGTH_SIZE uint16 = 0x2 // length field size of frame
	RSCP_FRAME_HEADER_SIZE uint16 = RSCP_FRAME_MAGIC_SIZE + RSCP_FRAME_CTRL_SIZE +
		RSCP_FRAME_TIME_SIZE + RSCP_FRAME_LENGTH_SIZE // header size of frame

	RSCP_FRAME_MAGIC_POS  uint8 = 0                                                     // magic field position in frame
	RSCP_FRAME_CTRL_POS   uint8 = RSCP_FRAME_MAGIC_POS + uint8(RSCP_FRAME_MAGIC_SIZE)   // control field position in frame
	RSCP_FRAME_TIME_POS   uint8 = RSCP_FRAME_CTRL_POS + uint8(RSCP_FRAME_CTRL_SIZE)     // time field position in frame
	RSCP_FRAME_LENGTH_POS uint8 = RSCP_FRAME_TIME_POS + uint8(RSCP_FRAME_TIME_SIZE)     // length field position in frame
	RSCP_FRAME_DATA_POS   uint8 = RSCP_FRAME_LENGTH_POS + uint8(RSCP_FRAME_LENGTH_SIZE) // data field position in frame

	RSCP_FRAME_CRC_SIZE      uint16 = 0x4    // crc field size of frame
	RSCP_FRAME_MAX_DATA_SIZE uint16 = 0xffff // max. data size of a frame (limited by lenth field size uint16)
	RSCP_FRAME_MAX_SIZE      uint32 = uint32(RSCP_FRAME_HEADER_SIZE) +
		uint32(RSCP_FRAME_MAX_DATA_SIZE) + uint32(RSCP_FRAME_CRC_SIZE) // max frame size
	RSCP_FRAME_MAX_BLOCK_SIZE uint16 = uint16((RSCP_FRAME_MAX_SIZE +
		(uint32(RSCP_CRYPT_BLOCK_SIZE) - 1)) / uint32(RSCP_CRYPT_BLOCK_SIZE))

	RSCP_DATA_TAG_SIZE      uint16 = 0x4 // tag field size of data field
	RSCP_DATA_DATATYPE_SIZE uint16 = 0x1 // data type field size of data field
	RSCP_DATA_LENGTH_SIZE   uint16 = 0x2 // length field size of data field

	// header size of data field
	RSCP_DATA_HEADER_SIZE uint16 = RSCP_DATA_TAG_SIZE + RSCP_DATA_DATATYPE_SIZE + RSCP_DATA_LENGTH_SIZE

	// max size a data field's data (is limited by the data field's header size and the frames uint16 max of length field)
	RSCP_DATA_MAX_DATA_SIZE uint16 = RSCP_FRAME_MAX_DATA_SIZE - RSCP_DATA_HEADER_SIZE
)

// control field of frame
//
//	Byte |        0        |        1        |
//	Bit  | 7 6 5 4 3 2 1 0 | 7 6 5 4 3 2 1 0 |
//	     | R R R C V V V V | R R R R R R R R |
//	     | R R R           | R R R R R R R R | Reserviert für zukünfige Erweiterungen -> RSCP_CTRL_BIT_MASK
//	     | 0 0 0           | 0 0 0 0 0 0 0 0 | Derzeitiger Zustand für die reservierten Bit
//	     |       C         |                 | Checksummen Flag -> RSCP_CTRL_BIT_MASK
//	     |       0         |                 | Checksumme wird nicht verwendet
//	     |       1         |                 | Checksumme wird verwendet. Das Feld CRC am Ende des Frames ist ein Pflichtfeld, ansonsten wird der Frame verworfen! -> RSCP_CRC_FLAG
//	     |         V V V V |                 | Versionskennzeichnung -> RSCP_CTRL_BIT_MASK_VERSION
//	     |         0 0 0 1 |                 | Version 1.0 (Momentan einzig zugelassener Wert) -> RSCP_VERSION_1_0
//
//nolint:lll,revive,stylecheck
const (
	RSCP_CTRL_BIT_MASK_CRC     uint16 = 0b0001000000000000 // allowed bit's for crc field of control field
	RSCP_FLAG_BIT_CRC          uint8  = 12                 // bit start position of crc field of control field
	RSCP_CRC_DISABLED          uint8  = 0b0                // do not use CRC
	RSCP_CRC_ENABLED           uint8  = 0b1                // use CRC
	RSCP_CTRL_BIT_MASK_VERSION uint16 = 0b0000111100000000 // allowed bit's for version field of control field
	RSCP_FLAG_BIT_VERSION      uint8  = 8                  // bit start position of version field of control field
	RSCP_VERSION_1_0           uint8  = 0b0001             // version 1.0

	RSCP_CTRL_BIT_MASK uint16 = RSCP_CTRL_BIT_MASK_CRC | RSCP_CTRL_BIT_MASK_VERSION // allowed bit's for control field

	// header constant for magic field of frame
	RSCP_MAGIC uint16 = 0xDCE3
)
