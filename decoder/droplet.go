package decoder

import (
	"strconv"
)


type TDroplet struct {
	Sensor                  string `json:"sensor"`
	Id                      string `json:"id"`
	Rssi 					int `json:"rssi"`
	Voltage           		int `json:"voltage"`
	Temperature             float64 `json:"temperature"`
	Humidity 				int `json:"humidity"`
	Light 					int `json:"light"`
	Motion 					int `json:"motion"`
}

func Droplet(data string, sensor string) TDroplet {
	d := Common(data)
	_id := d.id
	_rssi := d.rssi
	_temperature := dropletTemp(data)
	_humidity := dropletHumidity(data)
	_voltage := dropletVoltage(data)
	_light := dropletLight(data)
	_motion := dropletMotion(data)
	_v := TDroplet{Sensor: sensor, Id: _id, Rssi: _rssi, Voltage: _voltage, Temperature: _temperature, Humidity: _humidity, Light: _light, Motion: _motion}
	return _v
}

func dropletTemp(data string) float64 {
	v, _ := strconv.ParseInt(data[10:12]+data[8:10], 16, 0)
	v_ := float64(v) / 100
	return v_
}

func dropletHumidity(data string) int {
	v, _ := strconv.ParseInt(data[16:18], 16, 0)
	v_ := v & 127
	return int(v_)
}

func dropletVoltage(data string) int {
	v, _ := strconv.ParseInt(data[22:24], 16, 0)
	v_ := v / 50
	return int(v_)
}

func dropletLight(data string) int {
	v := data[20:22] + data[18:20]
	v_, _ := strconv.ParseInt(v, 16, 0)
	return int(v_)
}

func dropletMotion(data string) int {
	v_, _ := strconv.ParseInt(data[16:18], 16, 0)
	if v_ > 127 {
		return 1
	} else {
		return 0
	}
}
