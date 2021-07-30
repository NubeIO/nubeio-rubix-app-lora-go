package decoder

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"strconv"
)

type TZipHydrotapBase struct {
	CommonValues
	PayloadType string `json:"payload_type"`
}

type TZipHydrotapStatic struct {
	TZipHydrotapBase
	SerialNumber            string `json:"serial_number"`
	ModelNumber             string `json:"model_number"`
	ProductNumber           string `json:"product_number"`
	FirmwareVersion         string `json:"firmware_version"`
	CalibrationDate         string `json:"calibration_date"`
	First50LitresData       string `json:"first_50_litres_data"`
	FilterLogDateInternal   string `json:"filter_log_date_internal"`
	FilterLogLitresInternal int    `json:"filter_log_litres_internal"`
	FilterLogDateExternal   string `json:"filter_log_date_external"`
	FilterLogLitresExternal int    `json:"filter_log_litres_external"`
}

type TZipHydrotapWrite struct {
	TZipHydrotapBase
	Time                   string  `json:"time"`
	DispenseTimeBoiling    int     `json:"dispense_time_boiling"`
	DispenseTimeChilled    int     `json:"dispense_time_chilled"`
	DispenseTimeSparkling  int     `json:"dispense_time_sparkling"`
	TemperatureSPBoiling   float32 `json:"temperature_sp_boiling"`
	TemperatureSPChilled   float32 `json:"temperature_sp_chilled"`
	TemperatureSPSparkling float32 `json:"temperature_sp_sparkling"`
	// TODO: timers
	SleepModeSetting              int8   `json:"sleep_mode_setting"`
	FilterInfoWriteLitresInternal int    `json:"filter_info_write_litres_internal"`
	FilterInfoWriteMonthsInternal int    `json:"filter_info_write_months_internal"`
	FilterInfoWriteLitresExternal int    `json:"filter_info_write_litres_external"`
	FilterInfoWriteMonthsExternal int    `json:"filter_info_write_months_external"`
	SafetyAllowTapChanges         bool   `json:"safety_allow_tap_changes"`
	SafetyLock                    bool   `json:"safety_lock"`
	SafetyHotIsolation            bool   `json:"safety_hot_isolation"`
	SecurityEnable                bool   `json:"security_enable"`
	SecurityPin                   string `json:"security_pin"`
}

type TZipHydrotapPoll struct {
	TZipHydrotapBase
	StaticCOVFlag                     bool    `json:"static_cov_flag"`
	WriteCOVFlag                      bool    `json:"write_cov_flag"`
	SleepModeStatus                   int8    `json:"sleep_mode_status"`
	TemperatureNTCBoiling             float32 `json:"temperature_ntc_boiling"`
	TemperatureNTCChilled             float32 `json:"temperature_ntc_chilled"`
	TemperatureNTCStream              float32 `json:"temperature_ntc_stream"`
	TemperatureNTCCondensor           float32 `json:"temperature_ntc_condensor"`
	UsageEnergyKWh                    float32 `json:"usage_energy_kwh"`
	UsageWaterDeltaDispensesBoiling   int     `json:"usage_water_delta_dispenses_boiling"`
	UsageWaterDeltaDispensesChilled   int     `json:"usage_water_delta_dispenses_chilled"`
	UsageWaterDeltaDispensesSparkling int     `json:"usage_water_delta_dispenses_sparkling"`
	UsageWaterDeltaLitresBoiling      int     `json:"usage_water_delta_litres_boiling"`
	UsageWaterDeltaLitresChilled      int     `json:"usage_water_delta_litres_chilled"`
	UsageWaterDeltaLitresSparkling    int     `json:"usage_water_delta_litres_sparkling"`
	Fault1                            uint8   `json:"fault_1"`
	Fault2                            uint8   `json:"fault_2"`
	Fault3                            uint8   `json:"fault_3"`
	Fault4                            uint8   `json:"fault_4"`
	FilterWarningInternal             bool    `json:"filter_warning_internal"`
	FilterWarningExternal             bool    `json:"filter_warning_external"`
	FilterInfoUsageLitresInternal     int     `json:"filter_info_usage_litres_internal"`
	FilterInfoUsageDaysInternal       int     `json:"filter_info_usage_days_internal"`
	FilterInfoUsageLitresExternal     int     `json:"filter_info_usage_litres_external"`
	FilterInfoUsageDaysExternal       int     `json:"filter_info_usage_days_external"`
}

type TZHTPayloadType int

const (
	ErrorData  = 0
	StaticData = 1
	WriteData  = 2
	PollData   = 3
)

const ZHT_HEX_STR_DATA_START = 14

func ZipHydrotap(data string, sensor TSensorType) (TZipHydrotapBase, interface{}) {
	bytes := ZHtGetPayloadBytes(data)
	log.Printf("ZHT BYTES: %d", len(bytes))
	switch pl := ZHtGetPayloadType(data); pl {
	case StaticData:
		payload_full := ZHtStaticPayloadDecoder(bytes)
		common_data := Common(data, sensor)
		payload_full.CommonValues = common_data
		payload_full.PayloadType = "static"
		return payload_full.TZipHydrotapBase, payload_full
	case WriteData:
		payload_full := ZHtWritePayloadDecoder(bytes)
		common_data := Common(data, sensor)
		payload_full.CommonValues = common_data
		payload_full.PayloadType = "write"
		return payload_full.TZipHydrotapBase, payload_full
	case PollData:
		payload_full := ZHtPollPayloadDecoder(bytes)
		common_data := Common(data, sensor)
		payload_full.CommonValues = common_data
		payload_full.PayloadType = "poll"
		return payload_full.TZipHydrotapBase, payload_full
	}

	return TZipHydrotapBase{}, nil
}

func ZHtGetPayloadType(data string) TZHTPayloadType {
	pl_id, _ := strconv.ParseInt(data[14:16], 16, 0)
	return TZHTPayloadType(pl_id)
}

func ZHtCheckPayloadLength(data string) bool {
	payload_length := len(data) - 10 // removed addr, nonce and MAC
	payload_length /= 2
	payload_type := ZHtGetPayloadType(data)
	data_length, _ := strconv.ParseInt(data[12:14], 16, 0)
	log.Printf("ZHT data_length: %d\n", data_length)

	return (payload_type == StaticData && data_length == 93 && payload_length > 93) ||
		(payload_type == WriteData && data_length == 23 && payload_length > 23) ||
		(payload_type == PollData && data_length == 40 && payload_length > 40)
}

func ZHtGetPayloadBytes(data string) []byte {
	length, _ := strconv.ParseInt(data[12:14], 16, 0)
	bytes, _ := hex.DecodeString(data[14 : 14+length*2])
	return bytes
}

func ZHtBytesToString(bytes []byte) string {
	str := ""
	// fmt.Println()
	for _, b := range bytes {
		if b == 0 {
			break
		}
		str += string(b)
		// fmt.Printf("%d,", b)
	}
	// fmt.Println()
	return str
}

func ZHtBytesToDate(bytes []byte) string {
	return fmt.Sprintf("%d/%d/%d", bytes[0], bytes[1], bytes[2])
}

func ZHtStaticPayloadDecoder(data []byte) TZipHydrotapStatic {
	index := 1
	sn := ZHtBytesToString(data[index : index+15])
	index += 15
	mn := ZHtBytesToString(data[index : index+20])
	index += 20
	pn := ZHtBytesToString(data[index : index+20])
	index += 20
	fw := ZHtBytesToString(data[index : index+20])
	index += 20
	cal_date := ZHtBytesToDate(data[index : index+3])
	index += 3
	f50l_date := ZHtBytesToDate(data[index : index+3])
	index += 3
	filt_log_date_int := ZHtBytesToDate(data[index : index+3])
	index += 3
	filt_log_litres_int := int(binary.LittleEndian.Uint16(data[index : index+2]))
	index += 2
	filt_log_date_ext := ZHtBytesToDate(data[index : index+3])
	index += 3
	filt_log_litres_ext := int(binary.LittleEndian.Uint16(data[index : index+2]))
	index += 2
	return TZipHydrotapStatic{
		SerialNumber:            sn,
		ModelNumber:             mn,
		ProductNumber:           pn,
		FirmwareVersion:         fw,
		CalibrationDate:         cal_date,
		First50LitresData:       f50l_date,
		FilterLogDateInternal:   filt_log_date_int,
		FilterLogLitresInternal: filt_log_litres_int,
		FilterLogDateExternal:   filt_log_date_ext,
		FilterLogLitresExternal: filt_log_litres_ext,
	}
}

func ZHtWritePayloadDecoder(data []byte) TZipHydrotapWrite {
	index := 1
	time := fmt.Sprintf("%d", binary.LittleEndian.Uint32(data[index:index+4]))
	index += 4
	disp_b := int(data[index])
	index += 1
	disp_c := int(data[index])
	index += 1
	disp_s := int(data[index])
	index += 1
	temp_sp_b := float32(binary.LittleEndian.Uint16(data[index:index+2])) / 10
	index += 2
	temp_sp_c := int(data[index])
	index += 1
	temp_sp_s := int(data[index])
	index += 1
	// TODO: timers
	sm := int8(data[index])
	index += 1
	fil_lyf_ltr_int := binary.LittleEndian.Uint16(data[index : index+2])
	index += 2
	fil_lyf_mnth_int := int(data[index])
	index += 1
	fil_lyf_ltr_ext := binary.LittleEndian.Uint16(data[index : index+2])
	index += 2
	fil_lyf_mnth_ext := int(data[index])
	index += 1
	sf_tap := (data[index]>>2)&1 == 1
	sf_l := (data[index]>>1)&1 == 1
	sf_hi := (data[index]>>0)&1 == 1
	index += 1
	secUI16 := binary.LittleEndian.Uint16(data[index : index+2])
	sec_en := (secUI16>>15)&1 == 1
	sec_pin := fmt.Sprintf("%.4d", secUI16&0x7FFF)
	return TZipHydrotapWrite{
		Time:                          time,
		DispenseTimeBoiling:           disp_b,
		DispenseTimeChilled:           disp_c,
		DispenseTimeSparkling:         disp_s,
		TemperatureSPBoiling:          temp_sp_b,
		TemperatureSPChilled:          float32(temp_sp_c),
		TemperatureSPSparkling:        float32(temp_sp_s),
		SleepModeSetting:              sm,
		FilterInfoWriteLitresInternal: int(fil_lyf_ltr_int),
		FilterInfoWriteMonthsInternal: fil_lyf_mnth_int,
		FilterInfoWriteLitresExternal: int(fil_lyf_ltr_ext),
		FilterInfoWriteMonthsExternal: fil_lyf_mnth_ext,
		SafetyAllowTapChanges:         sf_tap,
		SafetyLock:                    sf_l,
		SafetyHotIsolation:            sf_hi,
		SecurityEnable:                sec_en,
		SecurityPin:                   sec_pin,
	}
}

func ZHtPollPayloadDecoder(data []byte) TZipHydrotapPoll {
	index := 1
	s_cov := (data[index]>>6)&1 == 1
	w_cov := (data[index]>>7)&1 == 1
	sms := int8((data[index]) & 0x3F)
	index += 1
	temp_b := float32(binary.LittleEndian.Uint16(data[index:index+2])) / 10
	index += 2
	temp_c := float32(binary.LittleEndian.Uint16(data[index:index+2])) / 10
	index += 2
	temp_s := float32(binary.LittleEndian.Uint16(data[index:index+2])) / 10
	index += 2
	temp_cond := float32(binary.LittleEndian.Uint16(data[index:index+2])) / 10
	index += 2
	f1 := data[index]
	index += 1
	f2 := data[index]
	index += 1
	f3 := data[index]
	index += 1
	f4 := data[index]
	index += 1
	kwh := math.Float32frombits(binary.LittleEndian.Uint32(data[index : index+4]))
	index += 4
	dlt_disp_b := binary.LittleEndian.Uint16(data[index : index+2])
	index += 2
	dlt_disp_c := binary.LittleEndian.Uint16(data[index : index+2])
	index += 2
	dlt_disp_s := binary.LittleEndian.Uint16(data[index : index+2])
	index += 2
	dlt_ltr_b := binary.LittleEndian.Uint16(data[index : index+2])
	index += 2
	dlt_ltr_c := binary.LittleEndian.Uint16(data[index : index+2])
	index += 2
	dlt_ltr_s := binary.LittleEndian.Uint16(data[index : index+2])
	index += 2
	fltr_wrn_int := (data[index]>>1)&1 == 1
	fltr_wrn_ext := (data[index]>>0)&1 == 1
	index += 1
	fltr_nfo_use_ltr_int := binary.LittleEndian.Uint16(data[index : index+2])
	index += 2
	fltr_nfo_use_day_int := binary.LittleEndian.Uint16(data[index : index+2])
	index += 2
	fltr_nfo_use_ltr_ext := binary.LittleEndian.Uint16(data[index : index+2])
	index += 2
	fltr_nfo_use_day_ext := binary.LittleEndian.Uint16(data[index : index+2])
	index += 2

	return TZipHydrotapPoll{
		StaticCOVFlag:                     s_cov,
		WriteCOVFlag:                      w_cov,
		SleepModeStatus:                   sms,
		TemperatureNTCBoiling:             temp_b,
		TemperatureNTCChilled:             temp_c,
		TemperatureNTCStream:              temp_s,
		TemperatureNTCCondensor:           temp_cond,
		UsageEnergyKWh:                    kwh,
		UsageWaterDeltaDispensesBoiling:   int(dlt_disp_b),
		UsageWaterDeltaDispensesChilled:   int(dlt_disp_c),
		UsageWaterDeltaDispensesSparkling: int(dlt_disp_s),
		UsageWaterDeltaLitresBoiling:      int(dlt_ltr_b),
		UsageWaterDeltaLitresChilled:      int(dlt_ltr_c),
		UsageWaterDeltaLitresSparkling:    int(dlt_ltr_s),
		Fault1:                            f1,
		Fault2:                            f2,
		Fault3:                            f3,
		Fault4:                            f4,
		FilterWarningInternal:             fltr_wrn_int,
		FilterWarningExternal:             fltr_wrn_ext,
		FilterInfoUsageLitresInternal:     int(fltr_nfo_use_ltr_int),
		FilterInfoUsageDaysInternal:       int(fltr_nfo_use_day_int),
		FilterInfoUsageLitresExternal:     int(fltr_nfo_use_ltr_ext),
		FilterInfoUsageDaysExternal:       int(fltr_nfo_use_day_ext),
	}
}
