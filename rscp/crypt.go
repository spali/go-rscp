package rscp

import (
	"bytes"

	"github.com/azihsoyn/rijndael256"
)

//nolint: golint,stylecheck,gomnd
const (
	RSCP_CRYPT_BLOCK_SIZE    uint8 = rijndael256.BlockSize
	RSCP_CRYPT_BLOCK_PADDING       = byte(0x00)
	RSCP_CRYPT_KEY_PADDING         = byte(0xff)
	RSCP_CRYPT_IV_PADDING          = byte(0xff)
)

const keySize = 32

type Key = [keySize]byte

// CreateAESKey creates a 32-bit key out of a password string
func CreateAESKey(key string) [keySize]byte {
	aesKey := [32]byte{}
	copy(aesKey[:], key)
	copy(aesKey[len(key):], bytes.Repeat([]byte{RSCP_CRYPT_KEY_PADDING}, keySize-len(key)))
	return aesKey
}

type IV = [RSCP_CRYPT_BLOCK_SIZE]byte

// NewIV returns a new initialized IV
func NewIV() IV {
	var iv IV
	copy(iv[:], bytes.Repeat([]byte{RSCP_CRYPT_IV_PADDING}, int(RSCP_CRYPT_BLOCK_SIZE))[:RSCP_CRYPT_BLOCK_SIZE])
	return iv
}
