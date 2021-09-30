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

func SerialOpenAndRead(port string, baud int, mqttConn *mqtt_lib.MqttConnection) {
	MqttConnection = mqttConn
	settings := GetSerialConfig()
	settings.Port = port
	settings.BaudRate = baud
	NewSerialConnection(settings)
	SerialReadForever()
}

func NewSerialConnection(settings TSerialSettings) {
	if SerialConn.Connected {
		log.Println("Existing serial port connection by this app is open So! close existing connection")
		err := SerialConn.Disconnect()
		if err != nil {
			return
		}
	}
	m := &serial.Mode{
		BaudRate: settings.BaudRate,
		Parity:   settings.Parity,
		DataBits: settings.DataBits,
		StopBits: settings.StopBits,
	}
	ports, _ := serial.GetPortsList()
	SerialConn.ActivePortList = ports
	log.Println("SerialPort try and connect to", settings.Port)
	// log.Println(ports)
	port, err := serial.Open(settings.Port, m)
	SerialConn.Port = port

	if err != nil {
		SerialConn.Error = true
		log.Fatal("ERROR ", err)
	}
	SerialConn.Connected = true
	log.Println("Connected to serial port", settings.Port)
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
			log.Printf("message count %d\n", count)

			common_data, full_data := decoder.DecodePayload(data)
			points.PublishSensor(common_data, full_data, MqttConnection)
		} else {
			log.Printf("INVALID SERIAL! size: %d\n", len(data))
		}
	}
}
