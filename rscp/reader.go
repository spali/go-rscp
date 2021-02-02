package rscp

import (
	"bytes"
	"crypto/cipher"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"reflect"
	"time"

	log "github.com/sirupsen/logrus"
)

// readHeader reads and checks the header
func readHeader(data []byte) (crcFlag bool, frameSize uint32, dataSize uint16, err error) {
	if binary.LittleEndian.Uint16(data[RSCP_FRAME_MAGIC_POS:]) != RSCP_MAGIC {
		return false, 0, 0, ErrRscpInvalidMagic
	}
	c := binary.LittleEndian.Uint16(data[RSCP_FRAME_CTRL_POS:])
	if (c | RSCP_CTRL_BIT_MASK) != RSCP_CTRL_BIT_MASK {
		return false, 0, 0, ErrRscpInvalidControl
	}
	if (c & RSCP_CTRL_BIT_MASK_VERSION) != (uint16(RSCP_VERSION_1_0) << RSCP_FLAG_BIT_VERSION) {
		return false, 0, 0, ErrRscpProtVersionMismatch
	}
	crcFlag = (c & RSCP_CTRL_BIT_MASK_CRC) == (uint16(RSCP_CRC_ENABLED) << RSCP_FLAG_BIT_CRC)
	dataSize = binary.LittleEndian.Uint16(data[RSCP_FRAME_LENGTH_POS:])
	frameSize = uint32(dataSize + RSCP_FRAME_HEADER_SIZE + (((c & RSCP_CTRL_BIT_MASK_CRC) >> RSCP_FLAG_BIT_CRC) * RSCP_FRAME_CRC_SIZE))
	return crcFlag, frameSize, dataSize, nil
}

// truncatePadding checks for padding after the frame and removes it.
// returns an error if there is unexpected data after the frame.
func truncatePadding(data *[]byte, frameSize uint32) error {
	// remove padding
	for i := len(*data); i > int(frameSize) && (*data)[i-1] == RSCP_CRYPT_BLOCK_PADDING; i-- {
		*data = (*data)[:i-1]
	}
	// check for unexpected remaining data
	if len(*data) > int(frameSize) {
		return fmt.Errorf("unexpected data after the frame found: %w", ErrRscpInvalidFrameLength)
	}
	return nil
}

// read wraps the binary.Read for little endian to read from the given buffer.
// for fixed size type's, "s" is not used and can be any valid uint16.
// wraps some rscp specific type's to read from.
func read(buf *bytes.Reader, v interface{}, size uint16) error {
	switch v := v.(type) {
	default:
		if err := binary.Read(buf, binary.LittleEndian, v); err != nil {
			return err
		}
	case *string:
		d := make([]byte, size)
		if err := read(buf, &d, size); err != nil {
			return err
		}
		*v = string(d)
	case *time.Time:
		var tS int64
		var tNs int32
		if err := read(buf, &tS, 8); err != nil {
			return err
		}
		if err := read(buf, &tNs, 4); err != nil {
			return err
		}
		*v = time.Unix(tS, int64(tNs)).UTC()
	case *[]Message:
		for s := buf.Len() - int(size); buf.Len() > s; {
			var (
				m   *Message
				err error
			)
			if m, err = readMessage(buf); err != nil {
				return fmt.Errorf("unexpected error reading data frame %d: %w", len(*v)+1, err)
			}
			*v = append(*v, *m)
		}
	}
	return nil
}

// readMessage reads a message recursive.
func readMessage(buf *bytes.Reader) (*Message, error) {
	m := new(Message)

	// read and check header
	if err := read(buf, &m.Tag, RSCP_DATA_TAG_SIZE); err != nil {
		return nil, err
	}
	if err := read(buf, &m.DataType, RSCP_DATA_DATATYPE_SIZE); err != nil {
		return nil, err
	}

	// Note: see Issue https://github.com/spali/go-e3dc/issues/1
	//
	// i.e request EMS_GET_SYS_SPECS is returning some Uint32 and some Int32 data types for EMS_SYS_SPEC_INDEX
	//  if m.DataType != m.Tag.DataType() && m.DataType != Error {
	//  	return nil, fmt.Errorf("expected data type %s for tag %s, got %s: %w", m.Tag.DataType(), m.Tag, m.DataType, ErrTagDataTypeMismatch)
	//  }

	var l uint16
	if err := read(buf, &l, RSCP_DATA_LENGTH_SIZE); err != nil {
		return nil, err
	}
	if l > RSCP_DATA_MAX_DATA_SIZE {
		return nil, ErrRscpDataLimitExceeded
	}
	// test length against known length of data type's
	if (m.DataType.length() != 0 || m.DataType == None) && m.DataType.length() != l {
		return nil, fmt.Errorf("length %d does not match expected length of data type %s: %w", l, m.DataType, ErrRscpDataLimitExceeded)
	}
	// read data
	v := m.DataType.newEmpty(l)
	if err := read(buf, v, l); err != nil {
		return nil, fmt.Errorf("reading message %s: %w", m.Tag, err)
	}
	m.Value = reflect.ValueOf(v).Elem().Interface() // dereference pointer value

	return m, nil
}

// Read decrypts and reads the data appends it to the buffer and returns the messages once the frame is complete.
func Read(mode *cipher.BlockMode, buf *[]byte, crcFlag *bool, frameSize *uint32, dataSize *uint16, data []byte) ([]Message, error) {
	if len(data) < int(RSCP_CRYPT_BLOCK_SIZE) || len(data)%int(RSCP_CRYPT_BLOCK_SIZE) != 0 {
		return nil, fmt.Errorf("require at least a block of %d bytes and the length must be a multiple of the block size: %w",
			RSCP_CRYPT_BLOCK_SIZE, ErrRscpInvalidFrameLength)
	}
	(*mode).CryptBlocks(data, data)

	if len(*buf) == 0 {
		var err error
		if *crcFlag, *frameSize, *dataSize, err = readHeader(data); err != nil {
			return nil, err
		}
		*buf = make([]byte, 0, *frameSize)
	}
	*buf = append(*buf, data...)

	if len(*buf) >= int(*frameSize) {
		if err := truncatePadding(buf, *frameSize); err != nil {
			return nil, err
		}

		// read frame data
		r := bytes.NewReader((*buf)[RSCP_FRAME_HEADER_SIZE:])
		var m []Message
		if err := read(r, &m, *dataSize); err != nil {
			return nil, err
		}

		log.Tracef("read plain %#v", *buf)
		log.Debugf("read %s", m)

		// read and check crc
		if *crcFlag {
			var crc uint32
			// can not fail, because frame/data size checks would catch it before, or padding would "fix" it, resulting in a wrong checksum below
			_ = read(r, &crc, RSCP_FRAME_CRC_SIZE)
			if crc != crc32.ChecksumIEEE((*buf)[:*frameSize-uint32(RSCP_FRAME_CRC_SIZE)]) {
				return nil, ErrRscpInvalidCrc
			}
		}
		return m, nil
	}
	return nil, ErrRscpInvalidFrameLength
}
