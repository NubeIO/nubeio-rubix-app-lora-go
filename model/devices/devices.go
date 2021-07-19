package modeldevices

import (
	"github.com/NubeIO/nubeio-rubix-app-lora-go/model/common"
	modelpoints "github.com/NubeIO/nubeio-rubix-app-lora-go/model/points"
)

type CommonDevice struct {
	Name        	string `json:"name" validate:"min=1,max=255"  gorm:"type:varchar(255);unique;not null"`
	Description 	string `json:"description"`
}

type Device struct {
	Uuid				string `json:"uuid"  gorm:"type:varchar(255);unique;primaryKey"`
	modelcommon.Common
	NetworkUuid     	string  `json:"network_uuid" gorm:"TYPE:varchar(255) REFERENCES networks;not null;default:null"`
	Point 				[]modelpoints.Point `json:"points" gorm:"constraint:OnDelete:CASCADE"`

}

