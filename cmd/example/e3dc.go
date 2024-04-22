package main

import (
	"fmt"
	"os"

	"github.com/jnovack/flag"
	"github.com/sirupsen/logrus"
	"github.com/spali/go-rscp/rscp"
)

type config struct {
	host     string
	port     uint
	user     string
	password string
	key      string
}

var conf = config{}

//nolint:funlen
func run() int {
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "E3DC", flag.ContinueOnError)
	fs.String(flag.DefaultConfigFlagname, ".config", "path to config file")
	fs.StringVar(&conf.host, "host", "", "e3dc server host")
	fs.UintVar(&conf.port, "port", 5033, "e3dc server host port") //nolint:gomnd
	fs.StringVar(&conf.user, "user", "", "e3dc user")
	fs.StringVar(&conf.password, "password", "", "e3dc password (consider using a config file or environment variable)")
	fs.StringVar(&conf.key, "key", "", "rscp key")
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "config parse error: %s\n", err)
	}
	c, err := rscp.NewClient(rscp.ClientConfig{
		Address:     conf.host,
		Port:        uint16(conf.port),
		Username:    conf.user,
		Password:    conf.password,
		Key:         conf.key,
		UseChecksum: true,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connection error: %s\n", err)
		return 1
	}
	rscp.Log.SetLevel(logrus.WarnLevel)
	rscp.Log.SetOutput(os.Stderr)
	defer func() { _ = c.Disconnect() }() // make sure to disconnect at the end
	{
		//
		// single simple request with simple response.
		//
		request := rscp.Message{
			Tag:      rscp.INFO_REQ_MAC_ADDRESS,
			DataType: rscp.None,
			Value:    nil,
		}
		var response *rscp.Message
		if response, err = c.Send(request); err != nil {
			fmt.Fprintf(os.Stderr, "Send error: %s\n", err)
			return 1
		}
		// prints (prettified and replaced Tag and DataType numbers with constants):
		//
		// rscp.Message{
		//   Tag:      rscp.INFO_MAC_ADDRESS,
		//   DataType: rscp.CString,
		//   Value:    "00:00:00:00:00:00",
		// }
		fmt.Printf("%#v\n", *response)
	}
	{
		//
		// single simple request with container response.
		//
		request := rscp.Message{
			Tag:      rscp.INFO_REQ_INFO,
			DataType: rscp.None,
			Value:    nil,
		}
		var response *rscp.Message
		if response, err = c.Send(request); err != nil {
			fmt.Fprintf(os.Stderr, "Send error: %s\n", err)
			return 1
		}
		// prints (prettified and replaced Tag and DataType numbers with constants):
		//
		// rscp.Message{
		//   Tag:      rscp.INFO_INFO,
		//   DataType: rscp.Container,
		//   Value:    []rscp.Message{
		//     {
		//       Tag:      rscp.INFO_SERIAL_NUMBER,
		//       DataType: rscp.CString,
		//       Value:    "Q10-000000000000",
		//     },
		//     {
		//       Tag:      rscp.INFO_PRODUCTION_DATE,
		//       DataType: rscp.Timestamp,
		//       Value:    time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
		//     },
		//     {
		//       Tag:      rscp.INFO_MAC_ADDRESS,
		//       DataType: rscp.CString,
		//       Value:    "00:00:00:00:00:00",
		//     },
		//   },
		// }
		fmt.Printf("%#v\n", *response)
	}
	{
		//
		// single container request with container response.
		//
		request := rscp.Message{
			Tag:      rscp.BAT_REQ_DATA,
			DataType: rscp.Container,
			Value: []rscp.Message{
				{
					Tag:      rscp.BAT_INDEX,
					DataType: rscp.UInt16,
					Value:    uint16(0),
				},
				{
					Tag:      rscp.BAT_REQ_SPECIFICATION,
					DataType: rscp.None,
					Value:    nil,
				},
			},
		}
		var response *rscp.Message
		if response, err = c.Send(request); err != nil {
			fmt.Fprintf(os.Stderr, "Send error: %s\n", err)
			return 1
		}
		// prints (prettified and replaced Tag and DataType numbers with constants):
		//
		// rscp.Message{
		//   Tag:      rscp.BAT_DATA,
		//   DataType: rscp.Container,
		//   Value:    []rscp.Message{
		//     rscp.Message{
		//       Tag:      rscp.BAT_INDEX,
		//       DataType: rscp.UInt16,
		//       Value:    0,
		//     },
		//     rscp.Message{
		//       Tag:      rscp.BAT_SPECIFICATION,
		//       DataType: rscp.Container,
		//       Value:    []rscp.Message{
		//         rscp.Message{
		//           Tag:      rscp.BAT_SPECIFIED_CAPACITY,
		//           DataType: rscp.Int32,
		//           Value:    26107,
		//         },
		//         rscp.Message{
		//           Tag:      rscp.BAT_SPECIFIED_DSCHARGE_POWER,
		//           DataType: rscp.Int32,
		//           Value:    12000,
		//         },
		//         rscp.Message{
		//           Tag:      rscp.BAT_SPECIFIED_CHARGE_POWER,
		//           DataType: rscp.Int32,
		//           Value:    12000,
		//         },
		//         rscp.Message{
		//           Tag:      rscp.BAT_SPECIFIED_MAX_DCB_COUNT,
		//           DataType: rscp.Int32,
		//           Value:    8,
		//         },
		//         rscp.Message{
		//           Tag:      rscp.BAT_ROLE,
		//           DataType: rscp.Uint32,
		//           Value:    257,
		//         },
		//       },
		//     },
		//   },
		// }
		fmt.Printf("%#v\n", *response)
	}
	{
		//
		// all together from above in a single request and response.
		//
		requests := []rscp.Message{
			{
				Tag:      rscp.INFO_REQ_MAC_ADDRESS,
				DataType: rscp.None,
				Value:    nil,
			},
			{
				Tag:      rscp.INFO_REQ_INFO,
				DataType: rscp.None,
				Value:    nil,
			},
			{
				Tag:      rscp.BAT_REQ_DATA,
				DataType: rscp.Container,
				Value: []rscp.Message{
					{
						Tag:      rscp.BAT_INDEX,
						DataType: rscp.UInt16,
						Value:    uint16(0),
					},
					{
						Tag:      rscp.BAT_REQ_SPECIFICATION,
						DataType: rscp.None,
						Value:    nil,
					},
				},
			},
		}
		var responses []rscp.Message
		if responses, err = c.SendMultiple(requests); err != nil {
			fmt.Fprintf(os.Stderr, "Send error: %s\n", err)
			return 1
		}
		// prints (prettified and replaced Tag and DataType numbers with constants):
		// []rscp.Message{
		//   {
		//     Tag:      rscp.INFO_MAC_ADDRESS,
		//     DataType: rscp.CString,
		//     Value:    "00:00:00:00:00:00",
		//   },
		//   {
		//     Tag:      rscp.INFO_INFO,
		//     DataType: rscp.Container,
		//     Value:    []rscp.Message{
		//       {
		//         Tag:      rscp.INFO_SERIAL_NUMBER,
		//         DataType: rscp.CString,
		//         Value:    "Q10-000000000000",
		//       },
		//       {
		//         Tag:      rscp.INFO_PRODUCTION_DATE,
		//         DataType: rscp.Timestamp,
		//         Value:    time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
		//       },
		//       {
		//         Tag:      rscp.INFO_MAC_ADDRESS,
		//         DataType: rscp.CString,
		//         Value:    "00:00:00:00:00:00",
		//       },
		//     },
		//   },
		//   {
		//     Tag:      rscp.BAT_DATA,
		//     DataType: rscp.Container,
		//     Value:    []rscp.Message{
		//       {
		//         Tag:      rscp.BAT_INDEX,
		//         DataType: rscp.UInt16,
		//         Value:    0,
		//       },
		//       {
		//         Tag:      rscp.BAT_SPECIFICATION,
		//         DataType: rscp.Container,
		//         Value:    []rscp.Message{
		//           {
		//             Tag:      rscp.BAT_SPECIFIED_CAPACITY,
		//             DataType: rscp.Int32,
		//             Value:    26107,
		//           },
		//           {
		//             Tag:      rscp.BAT_SPECIFIED_DSCHARGE_POWER,
		//             DataType: rscp.Int32,
		//             Value:    12000,
		//           },
		//           {
		//             Tag:      rscp.BAT_SPECIFIED_CHARGE_POWER,
		//             DataType: rscp.Int32,
		//             Value:    12000,
		//           },
		//           {
		//             Tag:      rscp.BAT_SPECIFIED_MAX_DCB_COUNT,
		//             DataType: rscp.Int32,
		//             Value:    8,
		//           },
		//           {
		//             Tag:      rscp.BAT_ROLE,
		//             DataType: rscp.Uint32,
		//             Value:    257,
		//           },
		//         },
		//       },
		//     },
		//   },
		// }
		fmt.Printf("%#v\n", responses)
	}
	return 0
}

//nolint:lll
func main() {
	// for reference, the whole output exactly as printed:
	//
	// rscp.Message{Tag:0xa80000a, DataType:0xd, Value:"00:00:00:00:00:00"}
	// rscp.Message{Tag:0xa800011, DataType:0xe, Value:[]rscp.Message{rscp.Message{Tag:0xa800001, DataType:0xd, Value:"Q10-000000000000"}, rscp.Message{Tag:0xa800002, DataType:0xf, Value:time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)}, rscp.Message{Tag:0xa80000a, DataType:0xd, Value:"00:00:00:00:00:00"}}}
	// rscp.Message{Tag:0x3840000, DataType:0xe, Value:[]rscp.Message{rscp.Message{Tag:0x3040001, DataType:0x5, Value:0x0}, rscp.Message{Tag:0x3800043, DataType:0xe, Value:[]rscp.Message{rscp.Message{Tag:0x3800125, DataType:0x6, Value:26107}, rscp.Message{Tag:0x3800126, DataType:0x6, Value:12000}, rscp.Message{Tag:0x3800127, DataType:0x6, Value:12000}, rscp.Message{Tag:0x3800128, DataType:0x6, Value:8}, rscp.Message{Tag:0x3800129, DataType:0x7, Value:0x101}}}}}
	// []rscp.Message{rscp.Message{Tag:0xa80000a, DataType:0xd, Value:"00:00:00:00:00:00"}, rscp.Message{Tag:0xa800011, DataType:0xe, Value:[]rscp.Message{rscp.Message{Tag:0xa800001, DataType:0xd, Value:"Q10-000000000000"}, rscp.Message{Tag:0xa800002, DataType:0xf, Value:time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)}, rscp.Message{Tag:0xa80000a, DataType:0xd, Value:"00:00:00:00:00:00"}}}, rscp.Message{Tag:0x3840000, DataType:0xe, Value:[]rscp.Message{rscp.Message{Tag:0x3040001, DataType:0x5, Value:0x0}, rscp.Message{Tag:0x3800043, DataType:0xe, Value:[]rscp.Message{rscp.Message{Tag:0x3800125, DataType:0x6, Value:26107}, rscp.Message{Tag:0x3800126, DataType:0x6, Value:12000}, rscp.Message{Tag:0x3800127, DataType:0x6, Value:12000}, rscp.Message{Tag:0x3800128, DataType:0x6, Value:8}, rscp.Message{Tag:0x3800129, DataType:0x7, Value:0x101}}}}}}
	os.Exit(run())
}
