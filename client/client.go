package client

import (
	"crypto/cipher"
	"fmt"
	"net"
	"time"

	"errors"

	"github.com/azihsoyn/rijndael256"
	log "github.com/sirupsen/logrus"
	"github.com/spali/go-e3dc/rscp"
)

// Client for rscp protocol
type Client struct {
	config           Config
	connectionString string
	isConnected      bool
	isAuthenticated  bool
	conn             net.Conn
	encrypter        cipher.BlockMode
	decrypter        cipher.BlockMode
}

// New creates a new client
func New(config Config) (*Client, error) {
	config, err := check(config)
	if err != nil {
		return nil, err
	}
	key := rscp.CreateAESKey(config.Key)
	initIV := rscp.NewIV()
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
func (c *Client) send(messages []rscp.Message) error {
	if err := rscp.ValidateRequests(messages); err != nil {
		return err
	}
	var (
		msg []byte
		err error
	)
	if msg, err = rscp.Write(&c.encrypter, messages, c.config.UseChecksum.(bool)); err != nil {
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
func (c *Client) receive() ([]rscp.Message, error) {
	if err := c.conn.SetReadDeadline(time.Now().Add(c.config.ReceiveTimeout)); err != nil {
		return nil, err
	}

	buf := make([]byte, 0, rscp.RSCP_FRAME_MAX_SIZE)
	var crcFlag bool
	var frameSize uint32
	var dataSize uint16
	var m []rscp.Message

	for i, new := 0, make([]byte, uint32(rscp.RSCP_CRYPT_BLOCK_SIZE)*uint32(c.config.ReceiveBufferBlockSize)); ; {
		var err error
		if i, err = c.conn.Read(new); err != nil {
			return nil, fmt.Errorf("error during receive response: %w", err)
		} else if i == 0 {
			return nil, rscp.ErrRscpInvalidFrameLength
		}
		switch m, err = rscp.Read(&c.decrypter, &buf, &crcFlag, &frameSize, &dataSize, new[:i]); {
		case errors.Is(err, rscp.ErrRscpInvalidFrameLength):
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
	if msg, err := rscp.CreateRequest(rscp.RSCP_REQ_AUTHENTICATION,
		rscp.RSCP_AUTHENTICATION_USER, c.config.Username, rscp.RSCP_AUTHENTICATION_PASSWORD, c.config.Password); err != nil {
		return err
	} else if err := c.send([]rscp.Message{*msg}); err != nil {
		return err
	}
	var (
		messages []rscp.Message
		err      error
	)
	if messages, err = c.receive(); err != nil {
		return fmt.Errorf("authentication error: %w", err)
	}
	if messages[0].Tag != rscp.RSCP_AUTHENTICATION || messages[0].Value.(uint8) == uint8(rscp.AUTH_LEVEL_NO_AUTH) {
		c.isAuthenticated = false
		return fmt.Errorf("authentication failed: %+v", messages[0])
	}
	c.isAuthenticated = true
	log.Infof("successfully authenticated (level: %s)", messages[0].Value)
	return nil
}

// Disconnect the client
func (c *Client) Disconnect() error {
	c.isAuthenticated = false
	c.isConnected = false

	if err := c.conn.Close(); err != nil {
		return err
	}
	log.Info("disconnected")
	return nil
}

// Send a message and return the response.
//
// connects and authenticates the first time used.
func (c *Client) Send(request rscp.Message) (*rscp.Message, error) {
	var (
		responses []rscp.Message
		err       error
	)
	if responses, err = (c.SendMultiple([]rscp.Message{request})); err != nil {
		return nil, err
	}
	return &responses[0], nil
}

// Send multiple messages in one round-trip and return the response.
//
// connects and authenticates the first time used.
func (c *Client) SendMultiple(requests []rscp.Message) ([]rscp.Message, error) {
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
		responses []rscp.Message
		err       error
	)
	if responses, err = c.receive(); err != nil {
		return nil, err
	}
	return responses, nil
}
