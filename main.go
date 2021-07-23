package main

import (
	"flag"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/controller/devices"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/controller/devices/device"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/controller/networks"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/controller/networks/network"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/controller/points"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/controller/points/point"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/serial"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/setup"
	"github.com/NubeIO/nubeio-rubix-lib-mqtt-go/pkg/mqtt_lib"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/rest"
	"io/ioutil"
	"log"
)




func enableLogging(enable bool) {
	if enable {
		log.Print("INIT APP: LOGGING IS DISABLED")
		log.SetOutput(ioutil.Discard)
	}

}

// @title GO Nube API
// @version 1.0
// @description nube api docs
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
func main() {

	serialPort := flag.String("serialPort", "/dev/ttyACM0", "serial port by name")
	baudRate := flag.Int("baudRate", 38400, "serial port baud rate")
	logging := flag.Bool("logging", true, "disable logging")

	flag.Parse()
	fmt.Println("serialPort:", *serialPort)
	fmt.Println("baudRate:", *baudRate)
	fmt.Println("logging:", *logging)
	_logging := *logging

	enableLogging(_logging)

	err := setup.InitMQTT();if err != nil {
		log.Println(err)
		return
	}
	db, err := setup.InitDB(_logging);if err != nil {
		log.Println(err)
		return
	}
	err = setup.InitSerial();if err != nil {
		log.Println(err)
		return
	}
	mqttConnection := mqtt_lib.NewConnection()
	go serial.NewSerialConnection(mqttConnection, true)


	app := rest.New(3)
	app.Controller(networks.New(db))
	app.Controller(network.New(db))
	app.Controller(devices.New(db))
	app.Controller(device.New(db))
	app.Controller(points.New(db))
	app.Controller(point.New(db))
	app.Controller(points.ByName(db))
	err = app.Run(":1920");if err != nil {
		log.Println(err)
		return
	}


}
