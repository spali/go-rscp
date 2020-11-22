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
	ErrMissingMessage  = errors.New("missing message argument")
)

type config struct {
	version  bool
	host     string
	port     uint
	user     string
	password string
	key      string
	message  string
	detail   bool
	debug    uint
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

func parseFlags() {
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "E3DC", flag.ContinueOnError)
	fs.BoolVar(&conf.version, "version", false, "output version details")
	fs.String(flag.DefaultConfigFlagname, ".config", "path to config file")
	fs.StringVar(&conf.host, "host", "", "e3dc server host")
	fs.UintVar(&conf.port, "port", 5033, "e3dc server host port")
	fs.StringVar(&conf.user, "user", "", "e3dc user")
	fs.StringVar(&conf.password, "password", "", "e3dc password (consider using a config file or environment variable)")
	fs.StringVar(&conf.key, "key", "", "rscp key")
	fs.BoolVar(&conf.detail, "detail", false, "return detailed json output")
	fs.UintVar(&conf.debug, "debug", 0, "enable set debug messages to stderr by setting log level (0-6)")
	_ = fs.Parse(os.Args[1:])
	checkFlags(fs)
}

func checkFlags(fs *flag.FlagSet) {
	if conf.version {
		printVersion()
		os.Exit(0)
	}
	if conf.host == "" {
		fmt.Fprintf(os.Stderr, "error: %s\n", ErrMissingHost)
		printUsage(fs)
		os.Exit(1)
	}
	if conf.user == "" {
		fmt.Fprintf(os.Stderr, "error: %s\n", ErrMissingUser)
		printUsage(fs)
		os.Exit(1)
	}
	if conf.password == "" {
		fmt.Fprintf(os.Stderr, "error: %s\n", ErrMissingPassword)
		printUsage(fs)
		os.Exit(1)
	}
	if conf.key == "" {
		fmt.Fprintf(os.Stderr, "error: %s\n", ErrMissingKey)
		printUsage(fs)
		os.Exit(1)
	}

	if fs.NArg() > 0 {
		conf.message = fs.Arg(0)
	} else {
		stat, _ := os.Stdin.Stat()
		if stdin := (stat.Mode() & os.ModeCharDevice) == 0; stdin {
			var bytes []byte
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return
			}
			conf.message = string(bytes)
		}
	}
	if conf.message == "" {
		fmt.Fprintf(os.Stderr, "error: %s\n", ErrMissingMessage)
		printUsage(fs)
		os.Exit(1)
	}
}
