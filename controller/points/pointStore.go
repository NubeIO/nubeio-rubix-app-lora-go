package points

import (
	"encoding/json"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/decoder"
	"github.com/NubeIO/nubeio-rubix-lib-mqtt-go/pkg/mqtt_lib"
	"math"
	"reflect"
)
//func test(v interface{}) {
//	st := reflect.ValueOf(v)
//	fmt.Println(666)
//	fmt.Println(666)
//
//	fmt.Println(st.Interface())
//	fmt.Println(666)
//
//}



type ReturnValue struct {
	Status string
	CustomStruct interface{}
}





func GetReturn(status string, class interface{}){
	//var result = ReturnValue {Status : status, CustomStruct: class}
	//fmt.Println(8888)
	//fmt.Println(result.Status)
	//fmt.Println(result.CustomStruct)
	//fmt.Println(88888)

	//msg, ok := result.CustomStruct.(decoder.TDroplet)
	//
	//if ok {
	//	fmt.Println(44444, msg)
	//} else {
	//	fmt.Println(555555, result.CustomStruct)
	//}
	//msg, ok := result.CustomStruct.(Message1)
	//if ok {
	//	fmt.Printf("Message1 is %s\n", msg.message)
	//}
}

func PublishDroplet(sensor string, d  decoder.TDroplet, mqttConn *mqtt_lib.MqttConnection){
	topic := "aaaaa"
	//_v := TDroplet{Sensor: sensor, Id: _id, Rssi: _rssi, Voltage: _voltage, Temperature: _temperature, Humidity: _humidity, Light: _light, Motion: _motion}
	//message := decoder.TDroplet{Sensor: sensor, Id: d.Id, Rssi: d.Rssi, Voltage: d.Voltage, Temperature: d.Temperature, Humidity: d.Humidity, Light:  d.Light, Motion:  d.Motion}
	//fmt.Println(666, sensor)

	v := reflect.ValueOf(d)
	v.Type().Name()
	jsonValue, _ := json.Marshal(d)
	//for i, x := range jsonValue {
	//	fmt.Println(i, x)
	//
	//}
	//jsonValue, _ := json.Marshal(message)
	mqttConn.Publish(string(jsonValue), topic)

}

func PublishMicro(sensor string, d decoder.TMicroEdge, mqttConn *mqtt_lib.MqttConnection){
	topic := "aaaaa"
	message := decoder.TMicroEdge{Sensor: sensor, Id: d.Id, Rssi: d.Rssi, Voltage: d.Voltage, Pulse: d.Pulse}


	jsonValue, _ := json.Marshal(message)
	//for i, x := range jsonValue {
	//	fmt.Println(i, x)
	//
	//}

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
