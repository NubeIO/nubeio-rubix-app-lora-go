package template

import modeldevices "github.com/NubeIO/nubeio-rubix-app-lora-go/model/devices"

type Template struct {
	Uuid					string `json:"uuid" gorm:"type:varchar(255);unique;not null;default:null;primaryKey"`
	modeldevices.Device

}

type Device struct {
	Uuid					string `json:"uuid" gorm:"type:varchar(255);unique;not null;default:null;primaryKey"`
	Name        			string `json:"name" validate:"min=1,max=255"  gorm:"type:varchar(255);unique;not null"`
	Description 			string `json:"description"`
	Id						string `json:"id"`

}

type Point struct {
	Uuid					string `json:"uuid" gorm:"type:varchar(255);unique;not null;default:null;primaryKey"`
	Name        			string `json:"name" validate:"min=1,max=255"  gorm:"type:varchar(255);unique;not null"`
	Description 			string `json:"description"`
	Id						string `json:"id"`

}
