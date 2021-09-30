package setup

import (
	"log"

	"github.com/NubeIO/nubeio-rubix-app-lora-go/serial"
)

func InitSerial(config serial.TSerialSettings) error {
	var args serial.Params
	args.UseConfigFile = false

	err := serial.SetSerialConfig(config, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
