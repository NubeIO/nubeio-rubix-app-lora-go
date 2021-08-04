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
	"github.com/NubeIO/nubeio-rubix-app-lora-go/setup"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/rest"
	"io/ioutil"
	"log"
	"strconv"
)

func enableLogging(enable bool) {
	if enable {
		log.Print("INIT APP: LOGGING IS DISABLED")
		log.SetOutput(ioutil.Discard)
	}
}

func main() {
	appSetting := setup.App()
	serialSetting := setup.Serial()

	flag.Parse()
	fmt.Println("SerialPort:", serialSetting.SerialPort)
	fmt.Println("BaudRate:", serialSetting.BaudRate)
	fmt.Println("Logging:", appSetting.Logging)
	enableLogging(appSetting.Logging)

	err := setup.InitMQTT()
	if err != nil {
		log.Println(err)
		return
	}
	db, err := setup.InitDB(appSetting.Logging, appSetting)
	if err != nil {
		log.Println(err)
		return
	}
	err = setup.InitSerial()
	if err != nil {
		log.Println(err)
		return
	}
	//TODO: uncomment and make it work
	//mqttConnection := mqtt_lib.NewConnection()
	//go serial.SerialOpenAndRead(mqttConnection)

	app := rest.New(3)
	app.Controller(networks.New(db))
	app.Controller(network.New(db))
	app.Controller(devices.New(db))
	app.Controller(device.New(db))
	app.Controller(points.New(db))
	app.Controller(point.New(db))
	app.Controller(points.ByName(db))
	err = app.Run(":" + strconv.Itoa(appSetting.Port))
	if err != nil {
		log.Println(err)
		return
	}
}
