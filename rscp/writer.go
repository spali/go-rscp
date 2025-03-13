package rscp

import (
	"bytes"
	"crypto/cipher"
	"encoding/binary"
	"hash/crc32"
	"time"
)

// Now returns the current local time.
// patchable for tests.
var Now = time.Now

// write wraps the binary.Write for little endian.
// supports some extra type's.
func write(buf *bytes.Buffer, v interface{}) error {
	if v == nil {
		return nil
	}
	// dereference in case we got a pointer to a value
	switch v := dereferencePtr(v).(type) {
	default:
		if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
			// not testable
			return err
		}
	case string:
		if err := write(buf, []byte(v)); err != nil {
			// not testable
			return err
		}
	case time.Time:
		if err := write(buf, v.Unix()); err != nil {
			// not testable
			return err
		}
		//nolint:gosec,mnd
		if err := write(buf, int32(v.UnixNano()-(v.Unix()*1e9))); err != nil {
			// not testable
			return err
		}
	case []Message:
		for i := 0; i < len(v); i++ {
			if err := writeMessage(buf, v[i]); err != nil {
				// not testable
				return err
			}
		}
	}
	return nil
}

// writeMessage writes a message recursive.
func writeMessage(buf *bytes.Buffer, m Message) error {
	if err := write(buf, m.Tag); err != nil {
		return err
	}
	if err := write(buf, m.DataType); err != nil {
		return err
	}
	l := m.valueSize()
	if err := write(buf, &l); err != nil {
		return err
	}
	if err := write(buf, m.Value); err != nil {
		return err
	}
	return nil
}

// writeFrame writes the rscp frame.
func writeFrame(messages []Message, useChecksum bool) ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, RSCP_FRAME_HEADER_SIZE))

	// write header
	if err := write(buf, RSCP_MAGIC); err != nil {
		return nil, err
	}
	ctrl := (RSCP_CTRL_BIT_MASK_VERSION & (uint16(RSCP_VERSION_1_0) << RSCP_FLAG_BIT_VERSION))
	if useChecksum {
		ctrl |= (RSCP_CTRL_BIT_MASK_CRC & (uint16(RSCP_CRC_ENABLED) << RSCP_FLAG_BIT_CRC))
	} else {
		ctrl |= (RSCP_CTRL_BIT_MASK_CRC & (uint16(RSCP_CRC_DISABLED) << RSCP_FLAG_BIT_CRC))
	}
	if err := write(buf, ctrl); err != nil {
		return nil, err
	}
	t := Now().UTC()
	if err := write(buf, t); err != nil {
		return nil, err
	}
	l := messagesSize(messages)
	if err := write(buf, l); err != nil {
		return nil, err
	}

	// write length and data
	if err := write(buf, messages); err != nil {
		return nil, err
	}

	// write checksum
	if useChecksum {
		if err := write(buf, crc32.ChecksumIEEE(buf.Bytes())); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

// Write writes the messages and returns the encrypted rscp frame.
func Write(mode *cipher.BlockMode, messages []Message, useChecksum bool) ([]byte, error) {
	Log.Debugf("write %s", messages)
	var (
		d   []byte
		err error
	)
	if d, err = writeFrame(messages, useChecksum); err != nil {
		return nil, err
	}
	Log.Tracef("write plain %#v", d)
	// padding
	if len(d)%int(RSCP_CRYPT_BLOCK_SIZE) != 0 {
		d = append(d, bytes.Repeat([]byte{RSCP_CRYPT_BLOCK_PADDING}, int(RSCP_CRYPT_BLOCK_SIZE)-len(d)%int(RSCP_CRYPT_BLOCK_SIZE))...)
	}
	// encrypt
	(*mode).CryptBlocks(d, d)
	Log.Tracef("write crypt %#v", d)
	return d, nil
}
