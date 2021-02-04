package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jnovack/flag"
)

// set during built time
var (
	name      string = "e3dc"
	source    string = "unknown"
	version   string = "unknown"
	commit    string = "unknown"
	platform  string = "unknown"
	buildTime string = "unknown"
)

var (
	ErrMissingHost     = errors.New("missing host argument")
	ErrMissingUser     = errors.New("missing user argument")
	ErrMissingPassword = errors.New("missing password argument")
	ErrMissingKey      = errors.New("missing key argument")
	ErrMissingRequest  = errors.New("missing request argument")
)

type config struct {
	help          bool
	version       bool
	file          string
	host          string
	port          uint
	user          string
	password      string
	key           string
	request       string
	output        string
	debug         uint
	splitrequests bool
}

var conf = config{}

func printVersion() {
	fmt.Fprintln(os.Stderr, name)
	fmt.Fprintf(os.Stderr, "%s\n", strings.Repeat("-", len(name)))
	fmt.Fprintf(os.Stderr, "Source:     %s\n", source)
	fmt.Fprintf(os.Stderr, "Version:    %s\n", version)
	fmt.Fprintf(os.Stderr, "Commit:     %s\n", commit)
	fmt.Fprintf(os.Stderr, "Platform:   %s\n", platform)
	fmt.Fprintf(os.Stderr, "Build Time: %s\n", buildTime)
}

func printUsage(fs *flag.FlagSet) {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] 'json request'\n", name)
	fs.PrintDefaults()
}

func parseFlags() (*flag.FlagSet, error) {
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "E3DC", flag.ContinueOnError)
	fs.BoolVar(&conf.help, "help", false, "output this help")
	fs.BoolVar(&conf.help, "h", false, "output this help")
	fs.BoolVar(&conf.version, "version", false, "output version details")
	fs.String(flag.DefaultConfigFlagname, ".config", "path to config file")
	fs.StringVar(&conf.file, "file", "", "path to request file")
	fs.StringVar(&conf.host, "host", "", "e3dc server host")
	fs.UintVar(&conf.port, "port", 5033, "e3dc server host port")
	fs.StringVar(&conf.user, "user", "", "e3dc user")
	fs.StringVar(&conf.password, "password", "", "e3dc password (consider using a config file or environment variable)")
	fs.StringVar(&conf.key, "key", "", "rscp key")
	fs.StringVar(&conf.output, "output", "jsonsimple", "control the output, possible values:\n"+
		"  jsonsimple: array with simple objects with tag key for the value\n"+
		"  json:       array of full message objects")
	fs.UintVar(&conf.debug, "debug", 0, "enable set debug messages to stderr by setting log level (0-6)")
	fs.BoolVar(&conf.splitrequests, "splitrequests", false, "split the request array to multiple requests.\n"+
		"this can help if the server sends a timeout on big requests")
	_ = fs.Parse(os.Args[1:])
	return checkFlags(fs)
}

func checkFlags(fs *flag.FlagSet) (*flag.FlagSet, error) {
	if conf.version {
		return fs, nil
	}
	if conf.host == "" {
		return fs, ErrMissingHost
	}
	if conf.user == "" {
		return fs, ErrMissingUser
	}
	if conf.password == "" {
		return fs, ErrMissingPassword
	}
	if conf.key == "" {
		return fs, ErrMissingKey
	}

	if fs.NArg() > 0 {
		conf.request = fs.Arg(0)
	} else {
		if conf.file != "" {
			var (
				m   []byte
				err error
			)
			if m, err = ioutil.ReadFile(conf.file); err != nil {
				return fs, fmt.Errorf("could not read input file: %s", err)
			}
			conf.request = string(m)
		} else {
			stat, _ := os.Stdin.Stat()
			if stdin := (stat.Mode() & os.ModeCharDevice) == 0; stdin {
				var bytes []byte
				bytes, err := ioutil.ReadAll(os.Stdin)
				if err != nil {
					return fs, fmt.Errorf("could not read stdin: %s", err)
				}
				conf.request = string(bytes)
			}
		}
	}
	if conf.request == "" {
		return fs, ErrMissingRequest
	}
	return fs, nil
}
