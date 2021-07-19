package model

import (
	modelcommon "github.com/NubeIO/nubeio-rubix-app-lora-go/model/common"
	modeldevices "github.com/NubeIO/nubeio-rubix-app-lora-go/model/devices"
)



type Network struct {
	Uuid			string 		`json:"uuid"  gorm:"type:varchar(255);unique;primaryKey"`
	modelcommon.Common
	Device 			[]modeldevices.Device `json:"devices" gorm:"constraint:OnDelete:CASCADE;"`
}

