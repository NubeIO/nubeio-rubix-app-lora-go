package serial

import (
	"bufio"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/controller/points"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/decoder"

	"github.com/NubeIO/nubeio-rubix-lib-mqtt-go/pkg/mqtt_lib"
	"go.bug.st/serial"
	"log"
)

type Connection struct {
	Port           serial.Port
	Enable         bool
	Connected      bool
	Error          bool
	ActivePortList []string
}

var Port Connection

func (c *Connection) Disconnect() error {
	return c.Port.Close()
}

type TMicroEdge struct {
	Sensor        string
}

var _disableDebug = false

func printString(msg string) {
	if !_disableDebug {
		log.Println(msg)
	}
}

func NewSerialConnection(mqttConn *mqtt_lib.MqttConnection, disableDebug bool) {
	if disableDebug {
		_disableDebug = true
	}
	c := GetSerialConfig()
	portName := c.Port
	baudRate := c.BaudRate
	parity := c.Parity
	stopBits := c.StopBits
	dataBits := c.DataBits
	if Port.Connected {
		log.Println("Existing serial port connection by this app is open So! close existing connection")
		err := Port.Disconnect()
		if err != nil {
			return
		}
	}
	m := &serial.Mode{
		BaudRate: baudRate,
		Parity:   parity,
		DataBits: dataBits,
		StopBits: stopBits,
	}
	ports, err := serial.GetPortsList()
	Port.ActivePortList = ports
	log.Println("SerialPort try and connect to", portName)
	log.Println(ports)
	port, err := serial.Open(portName, m)
	Port.Port = port

	if err != nil {
		Port.Error = true
		log.Fatal("ERROR Connected to serial port", err)
	}
	Port.Connected = true
	log.Println("Connected to serial port", portName)
	scanner := bufio.NewScanner(port)
	count := 0

	for scanner.Scan() {
		var data = scanner.Text()
		if decoder.CheckPayloadLength(data) {
			count = count + 1
			_log := fmt.Sprintf("loop count %d", count)
			printString(_log)
			s := decoder.CheckSensorType(data)
			_log = fmt.Sprintf("CheckSensorType %s", data)
			printString(_log)
			me := string(decoder.SensorNames.ME)
			thml := string(decoder.SensorNames.THML)
			if s == me {
				d := decoder.MicroEdge(data, me)
				_log = fmt.Sprintf("decoder.MicroEdge %s", data)
				printString(_log)
				points.PublishMicro(me, d, mqttConn)
				points.GetReturn(me, d)
				fmt.Println(d)
			} else if s == thml {
				d := decoder.Droplet(data, thml)
				_log = fmt.Sprintf("decoder.Droplet %s", data)
				printString(_log)
				points.GetReturn(thml, d)
				points.PublishDroplet(thml, d, mqttConn)
			}
		} else {
			_log := fmt.Sprintf("lora serial messsage size %d", len(data))
			printString(_log)
		}
	}
}
