package thing

import modelcommon "github.com/NubeIO/nubeio-rubix-app-lora-go/model/common"

type Thing struct {
	Uuid						string `json:"uuid" gorm:"type:varchar(255);unique;not null;default:null;primaryKey"`
	modelcommon.Common

}
