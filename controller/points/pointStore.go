package points

import (
	"encoding/json"
	"log"
	"math"

	"github.com/NubeIO/nubeio-rubix-app-lora-go/decoder"
	"github.com/NubeIO/nubeio-rubix-lib-mqtt-go/pkg/mqtt_lib"
)

func PublishSensor(common_sensor_data decoder.CommonValues, sensor_struct interface{}, mqttConn *mqtt_lib.MqttConnection) {
	jsonValue, _ := json.Marshal(sensor_struct)
	PublishJSON(common_sensor_data, jsonValue, mqttConn)
}

func PublishJSON(common_sensor_data decoder.CommonValues, jsonValue []byte, mqttConn *mqtt_lib.MqttConnection) {
	topic := "test-topic/" + string(common_sensor_data.Id)

	log.Printf("MQTT PUB: {\"topic\": \"%s\", \"payload\": \"%s\"}", topic, string(jsonValue))
	mqttConn.Publish(string(jsonValue), topic)
}

func cov(new float64, existingData float64, cov float64) (bool, float64) {
	c := new - existingData
	if math.Abs(c) < cov {
		return false, existingData
	} else {
		return false, new
	}
}
