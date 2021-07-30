package rscp

import (
	"crypto/cipher"
	"fmt"
	"net"
	"time"

	"errors"

	"github.com/azihsoyn/rijndael256"
	log "github.com/sirupsen/logrus"
)

// Client for rscp protocol
type Client struct {
	config           ClientConfig
	connectionString string
	isConnected      bool
	isAuthenticated  bool
	conn             net.Conn
	encrypter        cipher.BlockMode
	decrypter        cipher.BlockMode
}

// NewClient creates a new client
func NewClient(config ClientConfig) (*Client, error) {
	if err := config.check(); err != nil {
		return nil, err
	}
	key := createAESKey(config.Key)
	initIV := newIV()
	cipherBlock, _ := rijndael256.NewCipher(key[:]) // implementation does not return an error
	// Intitialize the Client structure.
	c := &Client{
		config:           config,
		connectionString: fmt.Sprintf("%s:%d", config.Address, config.Port),
		isConnected:      false,
		isAuthenticated:  false,
		encrypter:        cipher.NewCBCEncrypter(cipherBlock, initIV[:]),
		decrypter:        cipher.NewCBCDecrypter(cipherBlock, initIV[:]),
	}
	return c, nil
}

// send message
func (c *Client) send(messages []Message) error {
	if err := validateRequests(messages); err != nil {
		return err
	}
	var (
		msg []byte
		err error
	)
	if msg, err = Write(&c.encrypter, messages, c.config.UseChecksum.(bool)); err != nil {
		return err
	}
	if err := c.conn.SetWriteDeadline(time.Now().Add(c.config.SendTimeout)); err != nil {
		return err
	}
	if _, err := c.conn.Write(msg); err != nil {
		return err
	}
	return nil
}

// receive listens for a response and decodes the response
func (c *Client) receive() ([]Message, error) {
	if err := c.conn.SetReadDeadline(time.Now().Add(c.config.ReceiveTimeout)); err != nil {
		return nil, err
	}

	buf := make([]byte, 0, RSCP_FRAME_MAX_SIZE)
	var crcFlag bool
	var frameSize uint32
	var dataSize uint16
	var m []Message

	for i, new := 0, make([]byte, uint32(RSCP_CRYPT_BLOCK_SIZE)*uint32(c.config.ReceiveBufferBlockSize)); ; {
		var err error
		if i, err = c.conn.Read(new); err != nil {
			return nil, fmt.Errorf("error during receive response: %w", err)
		} else if i == 0 {
			return nil, ErrRscpInvalidFrameLength
		}
		switch m, err = Read(&c.decrypter, &buf, &crcFlag, &frameSize, &dataSize, new[:i]); {
		case errors.Is(err, ErrRscpInvalidFrameLength):
			// frame not complete
			continue
		case err != nil:
			return nil, err
		case m != nil:
			// frame complete
			return m, nil
		}
	}
}

// connect create connection
func (c *Client) connect() error {
	log.Infof("Connecting to %s", c.connectionString)
	var (
		conn net.Conn
		err  error
	)
	if conn, err = net.DialTimeout("tcp", c.connectionString, c.config.ConnectionTimeout); err != nil {
		c.isConnected = false
		return err
	}
	c.conn = conn
	c.isConnected = true
	log.Infof("successfully connected to %s", c.conn.RemoteAddr())
	return nil
}

// authenticate authenticates the connection
func (c *Client) authenticate() error {
	if msg, err := CreateRequest(RSCP_REQ_AUTHENTICATION,
		RSCP_AUTHENTICATION_USER, c.config.Username, RSCP_AUTHENTICATION_PASSWORD, c.config.Password); err != nil {
		return err
	} else if err := c.send([]Message{*msg}); err != nil {
		return err
	}
	var (
		messages []Message
		err      error
	)
	if messages, err = c.receive(); err != nil {
		return fmt.Errorf("authentication error: %w", err)
	}
	if messages[0].Tag != RSCP_AUTHENTICATION {
		c.isAuthenticated = false
		return fmt.Errorf("authentication failed: %+v", messages[0])
	}
	switch v := messages[0].Value.(type) {
	default:
		c.isAuthenticated = false
		return fmt.Errorf("authentication failed, received unexpected auth level data type %+v", messages[0])
	case int32:
		// wrong credentials returns 0 as Int32 instead of Uint8
		if v == int32(AUTH_LEVEL_NO_AUTH) {
			c.isAuthenticated = false
			return fmt.Errorf("authentication failed: %+v", AuthLevel(v))
		}
	case uint8:
		if v == uint8(AUTH_LEVEL_NO_AUTH) {
			c.isAuthenticated = false
			return fmt.Errorf("authentication failed: %+v", AuthLevel(v))
		}
	}
	c.isAuthenticated = true
	log.Infof("successfully authenticated (level: %s)", AuthLevel(messages[0].Value.(uint8)))
	return nil
}

// Disconnect the client
func (c *Client) Disconnect() error {
	c.isAuthenticated = false
	c.isConnected = false

	if c.isConnected && c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return err
		}
	}
	log.Info("disconnected")
	return nil
}

// Send a message and return the response.
//
// connects and authenticates the first time used.
func (c *Client) Send(request Message) (*Message, error) {
	var (
		responses []Message
		err       error
	)
	if responses, err = (c.SendMultiple([]Message{request})); err != nil {
		return nil, err
	}
	return &responses[0], nil
}

// Send multiple messages in one round-trip and return the response.
//
// connects and authenticates the first time used.
func (c *Client) SendMultiple(requests []Message) ([]Message, error) {
	if !c.isConnected {
		if err := c.connect(); err != nil {
			return nil, err
		}
	}
	if !c.isAuthenticated {
		if err := c.authenticate(); err != nil {
			return nil, err
		}
	}
	if err := c.send(requests); err != nil {
		return nil, err
	}
	var (
		responses []Message
		err       error
	)
	if responses, err = c.receive(); err != nil {
		return nil, err
	}
	return responses, nil
}
