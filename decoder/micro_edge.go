package decoder

import "strconv"

type TMicroEdge struct {
	Sensor        string
	Id            string
	Rssi, Voltage int
	Pulse         int
	AI1, AI2, AI3 float64
}

func MicroEdge(data string, sensor string) TMicroEdge {
	d := Common(data)
	_id := d.id
	_rssi := d.rssi
	_pulse := pulse(data)
	_ai1 := ai1(data)
	_ai2 := ai2(data)
	_ai3 := ai3(data)
	_voltage := voltage(data)
	_v := TMicroEdge{Sensor: sensor, Id: _id, Rssi: _rssi, Voltage: _voltage, Pulse: _pulse, AI1: _ai1, AI2: _ai2, AI3: _ai3}
	return _v
}

func pulse(data string) int {
	v, _ := strconv.ParseInt(data[8:16], 16, 0)
	return int(v)
}

func ai1(data string) float64 {
	v, _ := strconv.ParseInt(data[18:22], 16, 0)
	return float64(v)
}

func ai2(data string) float64 {
	v, _ := strconv.ParseInt(data[22:26], 16, 0)
	return float64(v)
}

func ai3(data string) float64 {
	v, _ := strconv.ParseInt(data[26:30], 16, 0)
	return float64(v)
}

func voltage(data string) int {
	v, _ := strconv.ParseInt(data[16:18], 16, 0)
	v_ := v / 50
	return int(v_)
}
