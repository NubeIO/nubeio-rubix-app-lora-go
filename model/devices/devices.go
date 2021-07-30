package modeldevices

import (
	"github.com/NubeIO/nubeio-rubix-app-lora-go/model/common"
	modelpoints "github.com/NubeIO/nubeio-rubix-app-lora-go/model/points"
)

type CommonDevice struct {
	Manufacture 	string `json:"manufacture"` // nube
	DeviceType		string `json:"device_type"` // droplet
	Model 			string `json:"model"` // thml

}

type Device struct {
	modelcommon.CommonUUID
	modelcommon.CommonName
	modelcommon.Common
	modelcommon.Created
	NetworkUuid     	string  `json:"network_uuid" gorm:"TYPE:varchar(255) REFERENCES networks;not null;default:null"`
	Point 				[]modelpoints.Point `json:"points" gorm:"constraint:OnDelete:CASCADE"`

}

