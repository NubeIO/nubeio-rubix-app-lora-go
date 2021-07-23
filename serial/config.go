package serial

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"go.bug.st/serial"
)

var (
	S Serial
)

type Serial struct {
	Enable   bool
	Port     string
	BaudRate int
	StopBits serial.StopBits
	Parity   serial.Parity
	DataBits int
	Timeout  int
}

type Params struct {
	UseConfigFile bool
	ConfigFile    string
	GenerateFile  bool
}

//SetSerialConfig
// if args Params.GenerateFile is true this will create a json serial_config file and will disregard
// if args Params.UseConfigFile is true is will use a local serial_config file
func SetSerialConfig(config Serial, args Params) error {
	if !args.UseConfigFile {
		S = config
	} else {
		configFile := args.ConfigFile
		if configFile == "" {
			return errors.New("no valid config file passed in")
		}
		file, err := os.Open(configFile)
		_config := Serial{}
		if err != nil {
			_config.Enable = true
			_config.Port = "/dev/ttyACM0"
			_config.BaudRate = 38400
			_config.StopBits = serial.OneStopBit
			_config.Parity = serial.NoParity
			_config.DataBits = 8
			// generate mqtt_config file if dont exist
			if args.GenerateFile {
				j, _ := json.Marshal(_config)
				err = ioutil.WriteFile(configFile, j, 0644)
			}
			S = _config
		} else {
			decoder := json.NewDecoder(file)
			err = decoder.Decode(&_config)
			S = _config
		}
	}
	return nil
}

func GetSerialConfig() Serial {
	return S
}
