package model

import (
	modelcommon "github.com/NubeIO/nubeio-rubix-app-lora-go/model/common"
	modeldevices "github.com/NubeIO/nubeio-rubix-app-lora-go/model/devices"
)

//https://project-haystack.org/doc/appendix/protocol

type Network struct {
	Uuid			string 		`json:"uuid"  gorm:"type:varchar(255);unique;primaryKey"`
	modelcommon.Common
	Manufacture 	string `json:"manufacture"`
	Model 			string `json:"model"`
	Device 			[]modeldevices.Device `json:"devices" gorm:"constraint:OnDelete:CASCADE;"`
}

