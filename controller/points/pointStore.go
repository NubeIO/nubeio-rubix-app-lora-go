package points

import (
	"encoding/json"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/decoder"
	"github.com/NubeIO/nubeio-rubix-lib-mqtt-go/pkg/mqtt_lib"
	"log"
)


func PublishDroplet(sensor string, d  decoder.TDroplet, mqttConn *mqtt_lib.MqttConnection){
	topic := "aaaaa"
	//fmt.Println(c)
	//d := decoder.Droplet(data, thml)
	message := decoder.TDroplet{Sensor: sensor, Id: d.Id, Rssi: d.Rssi, Voltage: d.Voltage}
	jsonValue, _ := json.Marshal(message)


	log.Println("MQTT messages, topic:", topic, " ", "data:", jsonValue)
	mqttConn.Publish(string(jsonValue), topic)

}

func PublishMicro(sensor string, d decoder.TMicroEdge, mqttConn *mqtt_lib.MqttConnection){
	topic := "aaaaa"
	message := decoder.TDroplet{Sensor: sensor, Id: d.Id, Rssi: d.Rssi, Voltage: d.Voltage}
	jsonValue, _ := json.Marshal(message)
	log.Println("MQTT messages, topic:", topic, " ", "data:", message)
	mqttConn.Publish(string(jsonValue), topic)

}