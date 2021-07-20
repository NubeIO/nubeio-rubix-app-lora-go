package modelpoints

import modelcommon "github.com/NubeIO/nubeio-rubix-app-lora-go/model/common"


type CommonPoint struct {
	Writeable 		bool   `json:"writeable"`
	Cov  			float64 `json:"cov"`
	Fallback 		float64 `json:"fallback"`
}

type Point struct {
	Uuid						string `json:"uuid" gorm:"type:varchar(255);unique;not null;default:null;primaryKey"`
	modelcommon.Common
	DeviceUuid     				string `json:"device_uuid" gorm:"TYPE:string REFERENCES devices;not null;default:null"`
	CommonPoint
	PriorityArrayModel 			PriorityArrayModel `json:"priority_array" gorm:"constraint:OnDelete:CASCADE"`
	PointStore 					PointStore `json:"point_store" gorm:"constraint:OnDelete:CASCADE"`

}

type PointStore struct {
	PointUuid     	string `json:"point_uuid" gorm:"REFERENCES points;not null;default:null;primaryKey"`
	Value  			float64 `json:"value"`
}

type PriorityArrayModel struct {
	PointUuid     	string `json:"point_uuid" gorm:"REFERENCES points;not null;default:null;primaryKey"`
	P1  			float64 `json:"_1"`
	P2  			float64 `json:"_2"`
	P3  			float64 `json:"_3"`
	P4  			float64 `json:"_4"`
	P5  			float64 `json:"_5"`
	P6  			float64 `json:"_6"`
	P7  			float64 `json:"_7"`
	P8  			float64 `json:"_8"`
	P9  			float64 `json:"_9"`
	P10  			float64 `json:"_10"`
	P11  			float64 `json:"_11"`
	P12  			float64 `json:"_12"`
	P13  			float64 `json:"_13"`
	P14  			float64 `json:"_14"`
	P15  			float64 `json:"_15"`
	P16  			float64 `json:"_16"`
}