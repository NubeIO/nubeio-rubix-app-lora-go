package decoder

import (
	"strconv"
)

type TZipHydrotapBase struct {
	CommonValues
}

type TZipHydrotapStatic struct {
	TZipHydrotapBase
	Data int `json:"data"`
}

type TZipHydrotapWrite struct {
	TZipHydrotapBase
	Data int `json:"data"`
}

type TZipHydrotapPoll struct {
	TZipHydrotapBase
	Data int `json:"data"`
}

type TZHTPayloadType int

const (
	ERROR_DATA  = 0
	STATIC_DATA = 1
	WRITE_DATA  = 2
	POLL_DATA   = 3
)

func ZipHydrotap(data string, sensor TSensorType) (TZipHydrotapBase, interface{}) {
	switch pl := ZHtGetPayloadType(data); pl {
	case STATIC_DATA:
		payload_full := ZHtStaticPayloadDecoder(data)
		common_data := Common(data, sensor)
		payload_full.CommonValues = common_data
		return payload_full.TZipHydrotapBase, payload_full
	case WRITE_DATA:
		payload_full := ZHtWritePayloadDecoder(data)
		common_data := Common(data, sensor)
		payload_full.CommonValues = common_data
		return payload_full.TZipHydrotapBase, payload_full
	case POLL_DATA:
		payload_full := ZHtPollPayloadDecoder(data)
		common_data := Common(data, sensor)
		payload_full.CommonValues = common_data
		return payload_full.TZipHydrotapBase, payload_full
	}

	return TZipHydrotapBase{}, nil
}

func ZHtGetPayloadType(data string) TZHTPayloadType {
	pl_id, _ := strconv.ParseInt(data[12:14], 16, 0)
	return TZHTPayloadType(pl_id)
}

func ZHtCheckPayloadLength(data string) bool {
	payload_length := len(data) - 10 // removed addr, nonce and MAC
	payload_length /= 2
	payload_type := ZHtGetPayloadType(data)
	data_length, _ := strconv.ParseInt(data[14:16], 16, 0)

	return (payload_type == STATIC_DATA && data_length == 93) ||
		(payload_type == WRITE_DATA && data_length == 23) ||
		(payload_type == POLL_DATA && data_length == 40)
}

func ZHtStaticPayloadDecoder(data string) TZipHydrotapStatic {
	return TZipHydrotapStatic{Data: 1}
}

func ZHtWritePayloadDecoder(data string) TZipHydrotapWrite {
	return TZipHydrotapWrite{Data: 1}
}

func ZHtPollPayloadDecoder(data string) TZipHydrotapPoll {
	return TZipHydrotapPoll{Data: 1}
}
