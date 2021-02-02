# go-e3dc

Go library (rscp) and command line utility (e3dc) to communicate with a E3DC system implementing the RSCP (Remote Storage Control Protocol).

Please be gentle with me, this is my first real go project and it's bigger than I thought ;)
There is still a lot to improve to be a good reusable library, but most works for my specific usecase using it as a command line utility.
Contributions are very welcome, see also the check [TODO](#TODO) or search in the code for `TODO:` comments.

## Usage

All arguments can also be set as `E3DC_` prefixed environment variable or in a config file like:
```
host 127.0.0.1
user myuser
password mypassword
key mykey
```

request sent, is always last argument or can be piped through `stdin` as json string.

See the examples below for more information.


### Examples

***Note**: top element has always to be an array*

***Note**: these examples assume you have the default `.config` file setup or exported the environment variables.*

* Short tag only notation just request some information
    ```sh
    echo '[["EMS_REQ_START_MANUAL_CHARGE",0]]' | e3dc
    ```
    Output:
    ```json
    [
        {
            "EMS_START_MANUAL_CHARGE": true
        }
    ]
    ```

* Short tag only notation just request some information
    ```sh
    e3dc '["INFO_REQ_MAC_ADDRESS", "INFO_REQ_UTC_TIME"]'
    ```
    Output:
    ```json
    [
        {
            "INFO_MAC_ADDRESS": "FC:69:47:48:D8:D2"
        },
        {
            "INFO_UTC_TIME": "2021-02-01T21:35:37.000651Z"
        }
    ]
    ```

* Tuple notation of tag's to request information
    ```sh
    e3dc '[["INFO_REQ_MAC_ADDRESS"], ["INFO_REQ_UTC_TIME"]]'
    ```
    Output:
    ```json
    [
        {
            "INFO_MAC_ADDRESS": "FC:69:47:48:D8:D2"
        },
        {
            "INFO_UTC_TIME": "2021-02-01T21:35:37.000651Z"
        }
    ]
    ```

* Tuple notation of tag's and values to send information (data type is inferred by the tag)
    ```sh
    e3dc '[["EMS_REQ_START_MANUAL_CHARGE", 3000]]'
    ```
    Output:
    ```json
    [
        {
            "EMS_START_MANUAL_CHARGE": true
        }
    ]
    ```
* Tuple notation of tag's and values to send information (with explicit data type)
  
    *Note: wrong data type is corrected by inferring it from the tag*
    ```sh
    e3dc '[["INFO_REQ_MAC_ADDRESS","None",""], ["INFO_REQ_UTC_TIME"]]'
    ```
    Output:
    ```json
    [
        {
            "INFO_MAC_ADDRESS": "FC:69:47:48:D8:D2"
        },
        {
            "INFO_UTC_TIME": "2021-02-01T21:35:37.000651Z"
        }
    ]
    ```

## TODO
 - [ ] more testing
 - [ ] more documentation
 - [ ] `cmd`
   - [x] read config file (security)
   - [x] read from environment variables (security)
   - [x] streamline logging 
   - [ ] implement optional (by flag) pair output (same as the tuple input without datatype)
 - [ ] `rscp`
   - [ ] cleanup API
     - [ ] probably expose `Message` as interface and make the struct internal (would allow to move the complete cmd specific json stuff out)
   - [ ] streamline logging
   - [ ] client: improve implementation to make it stable for keeping stable and connected when used in a service
   - [x] move `cmd` specific json marshalling out of `rscp` to command line utility `cmd`
   - [ ] move `cmd` specific json unmarshaling out of `rscp` to command line utility `cmd`