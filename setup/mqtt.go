package setup

import (
	"github.com/NubeIO/nubeio-rubix-lib-mqtt-go/mqtt_config"
	"log"
)

func InitMQTT() error {
	var mqtt mqtt_config.Params
	mqtt.UseConfigFile = false
	var br mqtt_config.Broker
	br.Host = "0.0.0.0"
	br.Port = "1883"
	err := mqtt_config.SetMqttConfig(br, mqtt); if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
