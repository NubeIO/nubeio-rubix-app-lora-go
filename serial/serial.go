package serial

import (
	"bufio"
	"log"

	"go.bug.st/serial"

	"github.com/NubeIO/nubeio-rubix-app-lora-go/controller/points"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/decoder"
	"github.com/NubeIO/nubeio-rubix-lib-mqtt-go/pkg/mqtt_lib"
)

type TSerialConnection struct {
	Port           serial.Port
	Enable         bool
	Connected      bool
	Error          bool
	ActivePortList []string
}

var SerialConn TSerialConnection
var MqttConnection *mqtt_lib.MqttConnection

func (c *TSerialConnection) Disconnect() error {
	return c.Port.Close()
}

func SerialOpenAndRead(mqttConn *mqtt_lib.MqttConnection) {
	MqttConnection = mqttConn
	NewSerialConnection()
	SerialReadForever()
}

func NewSerialConnection() {
	c := GetSerialConfig()
	portName := c.Port
	baudRate := c.BaudRate
	parity := c.Parity
	stopBits := c.StopBits
	dataBits := c.DataBits
	if SerialConn.Connected {
		log.Println("Existing serial port connection by this app is open So! close existing connection")
		err := SerialConn.Disconnect()
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
	ports, _ := serial.GetPortsList()
	SerialConn.ActivePortList = ports
	log.Println("SerialPort try and connect to", portName)
	log.Println(ports)
	port, err := serial.Open(portName, m)
	SerialConn.Port = port

	if err != nil {
		SerialConn.Error = true
		log.Fatal("ERROR ", err)
	}
	SerialConn.Connected = true
	log.Println("Connected to serial port", portName)
}

func SerialReadForever() {
	if SerialConn.Error || !SerialConn.Connected || SerialConn.Port == nil {
		return
	}
	count := 0
	scanner := bufio.NewScanner(SerialConn.Port)

	for scanner.Scan() {
		var data = scanner.Text()
		if decoder.CheckPayloadLength(data) {
			count = count + 1
			log.Printf("loop count %d", count)

			common_data, full_data := decoder.DecodePayload(data)
			points.PublishSensor(common_data, full_data, MqttConnection)
		} else {
			log.Printf("lora serial messsage size %d", len(data))
		}
	}
}
