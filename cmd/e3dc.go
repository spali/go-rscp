package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/spali/go-e3dc/client"
	"github.com/spali/go-e3dc/rscp"
)

func parseJSON(message string) ([]rscp.Message, error) {
	data := []rscp.Message{}
	if err := json.Unmarshal([]byte(message), &data); err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	parseFlags()
	if conf.debug > 0 {
		logrus.SetLevel(logrus.Level(conf.debug))
		logrus.SetOutput(os.Stderr)
	} else {
		logrus.SetLevel(logrus.PanicLevel)
	}
	c, err := client.New(client.Config{
		Address:     conf.host,
		Port:        uint16(conf.port),
		Username:    conf.user,
		Password:    conf.password,
		Key:         conf.key,
		UseChecksum: true,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	m, err := parseJSON(conf.request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	r, err := c.SendMultiple(m)
	_ = c.Disconnect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	var rb []byte
	if conf.detail {
		if rb, err = json.Marshal(r); err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
	} else {
		if rb, err = json.Marshal(NewJSONSimpleMessages(r)); err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
	}
	fmt.Printf("%s\n", rb)
}
