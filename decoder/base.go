package decoder

import (
	"strconv"
)

type SensorTypes string

var sensorCodes = struct {
	MicroAa   SensorTypes
	DropletAb SensorTypes
	DropletB0 SensorTypes
	DropletB1 SensorTypes
	DropletB2 SensorTypes
}{
	MicroAa:   "AA",
	DropletAb: "AB",
	DropletB0: "B0",
	DropletB1: "B1",
	DropletB2: "B2",
}

var SensorNames = struct {
	ME   SensorTypes
	TH   SensorTypes
	THL  SensorTypes
	THML SensorTypes
}{
	ME:   "ME",
	TH:   "TH",
	THL:  "THL",
	THML: "THML",
}

func CheckSensorType(data string) string {
	sensor := data[2:4]

	switch sensor {
	case string(sensorCodes.MicroAa):
		return string(SensorNames.ME)
	case string(sensorCodes.DropletB2):
		return string(SensorNames.THML)
	default:
		return "None"
	}

}

func CheckPayloadLength(data string) bool {
	dl := len(data)
	if dl == 36 || dl == 32 || dl == 44 {
		return true
	} else {
		return false
	}

}

type CommonValues struct {
	id   string
	rssi int
}

func Common(data string) CommonValues {
	id := decodeID(data)
	_rssi := rssi(data)
	_v := CommonValues{id: id, rssi: _rssi}
	return _v
}

func DataLength(data string) int {
	return len(data)
}

func decodeID(data string) string {
	return data[0:8]
}

func rssi(data string) int {
	_len := DataLength(data)
	v, _ := strconv.ParseInt(data[_len-4:_len-2], 16, 0)
	_v := v * -1
	return int(_v)
}
