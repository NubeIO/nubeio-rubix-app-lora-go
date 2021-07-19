package serial

import (
	"bufio"
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


func NewSerialConnection(mqttConn *mqtt_lib.MqttConnection) {

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
	//topic := "test"

	for scanner.Scan() {
		var data = scanner.Text()
		if decoder.CheckPayloadLength(data) {
			count = count + 1
			log.Println("loop count", count)
			s := decoder.CheckSensorType(data)

			log.Println("Raw serial messages", data)
			me := string(decoder.SensorNames.ME)
			thml := string(decoder.SensorNames.THML)
			//var message interface{}
			if s == me {
				d := decoder.MicroEdge(data, me)
				log.Println(d)
				points.PublishMicro(me, d, mqttConn)
				//message = decoder.TMicroEdge{Sensor: me, Id: d.Id, Rssi: d.Rssi, Voltage: d.Voltage}
			} else if s == thml {
				d := decoder.Droplet(data, thml)
				//message = decoder.TDroplet{Sensor: thml, Id: d.Id, Rssi: d.Rssi, Voltage: d.Voltage}
				points.PublishDroplet(thml, d, mqttConn)
			}

			//var message interface{}
			//if s == me {
			//	d := decoder.MicroEdge(data, me)
			//	log.Println(d)
			//	message = decoder.TMicroEdge{Sensor: me, Id: d.Id, Rssi: d.Rssi, Voltage: d.Voltage}
			//} else if s == thml {
			//	d := decoder.Droplet(data, thml)
			//	log.Println(d)
			//	message = decoder.TDroplet{Sensor: thml, Id: d.Id, Rssi: d.Rssi, Voltage: d.Voltage}
			//	points.PublishPoints(d, mqttConn)
			//}
			//jsonValue, _ := json.Marshal(message)
			//log.Println("MQTT messages, topic:", topic, " ", "data:", message)
			//mqttConn.Publish(string(jsonValue), topic)

		} else {
			log.Println(len(data), "size")
		}

	}



}
