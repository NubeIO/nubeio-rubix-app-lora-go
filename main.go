package main

import (
	"github.com/NubeIO/nubeio-rubix-app-lora-go/controller/devices"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/controller/devices/device"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/controller/networks"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/controller/networks/network"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/controller/points"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/controller/points/point"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/serial"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/setup"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/logs"
	"github.com/NubeIO/nubeio-rubix-lib-mqtt-go/pkg/mqtt_lib"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/rest"
	"log"
)

var DisableLogging bool = false

func init() {
	logs.DisableLogging(DisableLogging)
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

	err := setup.InitMQTT();if err != nil {
		log.Println(err)
		return
	}
	db, err := setup.InitDB(DisableLogging);if err != nil {
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
	err = app.Run(":1920");if err != nil {
		log.Println(err)
		return
	}


}
