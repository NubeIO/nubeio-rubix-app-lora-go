package setup

import (
	"github.com/NubeIO/nubeio-rubix-app-lora-go/serial"
	"log"
)

func InitSerial() error {
	var args serial.Params
	args.UseConfigFile = false

	var config serial.Serial
	config.Port = "/dev/ttyACM0"
	config.BaudRate = 38400

	err := serial.SetSerialConfig(config, args); if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

