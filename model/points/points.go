package modelpoints

import modelcommon "github.com/NubeIO/nubeio-rubix-app-lora-go/model/common"


type CommonPoint struct {
	Writeable bool   			`json:"writeable"`
}


type Point struct {
	Uuid			string `json:"uuid" gorm:"type:varchar(255);unique;not null;default:null;primaryKey"`
	modelcommon.Common
	DeviceUuid     	string `json:"device_uuid" gorm:"TYPE:string REFERENCES devices;not null;default:null"`
	CommonPoint
	//PointStore 	PointStore `json:"point_store"`
	PointStore 	PointStore `json:"point_store" gorm:"constraint:OnDelete:CASCADE"`

}

type PointStore struct {
	Uuid			string `json:"uuid" gorm:"type:varchar(255);unique;not null;default:null;primaryKey"`
	PointUuid     	string `gorm:"TYPE:string REFERENCES points;not null;default:null"`
}
