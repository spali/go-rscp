package rscp

import (
	"crypto/cipher"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"time"

	"github.com/azihsoyn/rijndael256"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

// Client for rscp protocol
type Client struct {
	config           ClientConfig
	connectionString string
	isAuthenticated  bool
	conn             net.Conn
	encrypter        cipher.BlockMode
	decrypter        cipher.BlockMode
}

const RequiredAuthLogLevel = 99

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
	msg, err := Write(&c.encrypter, messages, c.config.UseChecksum.(bool))
	if err != nil {
		return err
	}
	if err := c.conn.SetWriteDeadline(time.Now().Add(c.config.SendTimeout)); err != nil {
		_ = c.Disconnect()
		return err
	}
	if _, err := c.conn.Write(msg); err != nil {
		_ = c.Disconnect()
		return err
	}
	return nil
}

// receive listens for a response and decodes the response
func (c *Client) receive() ([]Message, error) {
	if err := c.conn.SetReadDeadline(time.Now().Add(c.config.ReceiveTimeout)); err != nil {
		_ = c.Disconnect()
		return nil, err
	}

	buf := make([]byte, 0, RSCP_FRAME_MAX_SIZE)
	var crcFlag bool
	var frameSize uint32
	var dataSize uint16
	var m []Message

	for i, data := 0, make([]byte, uint32(RSCP_CRYPT_BLOCK_SIZE)*uint32(c.config.ReceiveBufferBlockSize)); ; {
		var err error

		if i, err = c.conn.Read(data); err != nil {
			_ = c.Disconnect()
			return nil, err
		} else if i == 0 {
			return nil, ErrRscpInvalidFrameLength
		}

		switch m, err = Read(&c.decrypter, &buf, &crcFlag, &frameSize, &dataSize, data[:i]); {
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
	Log.Infof("Connecting to %s", c.connectionString)

	conn, err := net.DialTimeout("tcp", c.connectionString, c.config.ConnectionTimeout)
	if err != nil {
		return err
	}

	Log.Infof("successfully connected to %s", conn.RemoteAddr())
	c.conn = conn

	return nil
}

// authenticate authenticates the connection
func (c *Client) authenticate() error {
	orgLogLevel := Log.GetLevel()
	if orgLogLevel < RequiredAuthLogLevel {
		Log.Infof("hiding auth request for security, use debug >= %d to debug authentication", RequiredAuthLogLevel)
		Log.SetLevel(logrus.Level((math.Min(float64(orgLogLevel), float64(logrus.InfoLevel)))))
	}
	if msg, err := CreateRequest(RSCP_REQ_AUTHENTICATION,
		RSCP_AUTHENTICATION_USER, c.config.Username, RSCP_AUTHENTICATION_PASSWORD, c.config.Password); err != nil {
		if orgLogLevel < RequiredAuthLogLevel {
			Log.SetLevel(orgLogLevel)
		}
		return err
	} else if err := c.send([]Message{*msg}); err != nil {
		if orgLogLevel < RequiredAuthLogLevel {
			Log.SetLevel(orgLogLevel)
		}
		return err
	}
	if orgLogLevel < RequiredAuthLogLevel {
		Log.SetLevel(orgLogLevel)
	}

	messages, err := c.receive()
	if err != nil {
		if errors.Is(err, io.EOF) {
			Log.Warnf("Hint: EOF during authentification usually is due a wrong rscp key")
		}
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
	Log.Infof("successfully authenticated (level: %s)", AuthLevel(messages[0].Value.(uint8)))
	return nil
}

// Disconnect the client
func (c *Client) Disconnect() (err error) {
	c.isAuthenticated = false

	if c.conn != nil {
		err = c.conn.Close()
		c.conn = nil
		Log.Info("disconnected")
	}

	return err
}

// Send a message and return the response.
//
// connects and authenticates the first time used.
func (c *Client) Send(request Message) (*Message, error) {
	responses, err := c.SendMultiple([]Message{request})
	if err != nil {
		return nil, err
	}

	return &responses[0], nil
}

// Send multiple messages in one round-trip and return the response.
//
// connects and authenticates the first time used.
func (c *Client) SendMultiple(requests []Message) ([]Message, error) {
	if c.conn == nil {
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
	return c.receive()
}
