package model

import (
	modelcommon "github.com/NubeIO/nubeio-rubix-app-lora-go/model/common"
	modeldevices "github.com/NubeIO/nubeio-rubix-app-lora-go/model/devices"
)

//https://project-haystack.org/doc/appendix/protocol

type Network struct {
	modelcommon.CommonUUID
	modelcommon.CommonNameUnique
	modelcommon.Common
	modelcommon.Created
	Manufacture 	string `json:"manufacture"`
	Model 			string `json:"model"`
	NetworkType		string `json:"network_type"`
	Device 			[]modeldevices.Device `json:"devices" gorm:"constraint:OnDelete:CASCADE;"`
}

