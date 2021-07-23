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
	ErrorData = 0
	StaticData = 1
	WriteData  = 2
	PollData   = 3
)

func ZipHydrotap(data string, sensor TSensorType) (TZipHydrotapBase, interface{}) {
	switch pl := ZHtGetPayloadType(data); pl {
	case StaticData:
		payloadFull := ZHtStaticPayloadDecoder(data)
		commonData := Common(data, sensor)
		payloadFull.CommonValues = commonData
		return payloadFull.TZipHydrotapBase, payloadFull
	case WriteData:
		payloadFull := ZHtWritePayloadDecoder(data)
		commonData := Common(data, sensor)
		payloadFull.CommonValues = commonData
		return payloadFull.TZipHydrotapBase, payloadFull
	case PollData:
		payloadFull := ZHtPollPayloadDecoder(data)
		commonData := Common(data, sensor)
		payloadFull.CommonValues = commonData
		return payloadFull.TZipHydrotapBase, payloadFull
	}

	return TZipHydrotapBase{}, nil
}

func ZHtGetPayloadType(data string) TZHTPayloadType {
	plId, _ := strconv.ParseInt(data[12:14], 16, 0)
	return TZHTPayloadType(plId)
}

func ZHtCheckPayloadLength(data string) bool {
	payloadLength := len(data) - 10 // removed addr, nonce and MAC
	payloadLength /= 2
	payloadType := ZHtGetPayloadType(data)
	dataLength, _ := strconv.ParseInt(data[14:16], 16, 0)

	return (payloadType == StaticData && dataLength == 93) ||
		(payloadType == WriteData && dataLength == 23) ||
		(payloadType == PollData && dataLength == 40)
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
