package points

import (
	"encoding/json"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/decoder"
	"github.com/NubeIO/nubeio-rubix-lib-mqtt-go/pkg/mqtt_lib"
	"math"
)


func PublishDroplet(sensor string, d  decoder.TDroplet, mqttConn *mqtt_lib.MqttConnection){
	topic := "aaaaa"
	message := decoder.TDroplet{Sensor: sensor, Id: d.Id, Rssi: d.Rssi, Voltage: d.Voltage, Temperature: d.Temperature}
	jsonValue, _ := json.Marshal(message)
	mqttConn.Publish(string(jsonValue), topic)

}

func PublishMicro(sensor string, d decoder.TMicroEdge, mqttConn *mqtt_lib.MqttConnection){
	topic := "aaaaa"
	message := decoder.TMicroEdge{Sensor: sensor, Id: d.Id, Rssi: d.Rssi, Voltage: d.Voltage, Pulse: d.Pulse}
	jsonValue, _ := json.Marshal(message)
	mqttConn.Publish(string(jsonValue), topic)

}



func cov(new float64, existingData float64, cov float64) (bool, float64) {
	c:= new - existingData
	if math.Abs(c) < cov {
		return false, existingData
	} else {
		return false, new
	}

}
