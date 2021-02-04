package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spali/go-rscp/rscp"
)

func run() ([]byte, error) {
	c, err := rscp.NewClient(rscp.ClientConfig{
		Address:     conf.host,
		Port:        uint16(conf.port),
		Username:    conf.user,
		Password:    conf.password,
		Key:         conf.key,
		UseChecksum: true,
	})
	if err != nil {
		return nil, err
	}
	defer func() { _ = c.Disconnect() }()
	var (
		rb []byte
		ms []rscp.Message
		rs []rscp.Message
	)
	if ms, err = unmarshalJSONRequests([]byte(conf.request)); err != nil {
		return nil, err
	}
	if conf.splitrequests {
		rs = make([]rscp.Message, len(ms))
		for i := range ms {
			var r *rscp.Message
			if r, err = c.Send(ms[i]); err != nil {
				return nil, err
			}
			rs[i] = *r
		}
	} else if rs, err = c.SendMultiple(ms); err != nil {
		return nil, err
	}
	switch conf.output {
	case "json":
		if rb, err = json.Marshal(rs); err != nil {
			return nil, err
		}
	case "jsonsimple":
		if rb, err = json.Marshal(NewJSONSimpleMessages(rs)); err != nil {
			return nil, err
		}
	case "jsonmerged":
		if rb, err = json.Marshal(NewJSONMergedMessages(rs)); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("output %s not supported", conf.output)
	}
	return rb, nil
}

func main() {
	switch fs, err := parseFlags(); {
	case conf.help:
		printUsage(fs)
		os.Exit(0)
	case conf.version:
		printVersion()
		os.Exit(0)
	case err != nil:
		// workaround for https://github.com/jnovack/flag/issues/1
		if !errors.Is(err, ErrFlagError) {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			printUsage(fs)
		}
		os.Exit(1)
	}
	if conf.debug > 0 {
		logrus.SetLevel(logrus.Level(conf.debug))
		logrus.SetOutput(os.Stderr)
	} else {
		logrus.SetLevel(logrus.PanicLevel)
	}
	var (
		rb  []byte
		err error
	)
	if rb, err = run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", rb)
}
