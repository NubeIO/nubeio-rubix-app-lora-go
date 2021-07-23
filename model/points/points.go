package modelpoints

import modelcommon "github.com/NubeIO/nubeio-rubix-app-lora-go/model/common"

type CommonPoint struct {
	Writeable 		bool   `json:"writeable"`
	Cov  			float64 `json:"cov"`
	ObjectType		string `json:"object_type"`
}

type EquipRef struct {
	PointUuid     	string `json:"point_uuid" gorm:"REFERENCES points;not null;default:null;primaryKey"`
	Ref  			string `json:"equip_ref"`
}

type Association struct {
	PointUuid     	string `json:"point_uuid" gorm:"REFERENCES points;not null;default:null;primaryKey"`
	Association  	string `json:"equip_ref"`
}


type Point struct {
	Uuid						string `json:"uuid" gorm:"type:varchar(255);unique;not null;default:null;primaryKey"`
	modelcommon.Common
	DeviceUuid     				string `json:"device_uuid" gorm:"TYPE:string REFERENCES devices;not null;default:null"`
	CommonPoint
	EquipRef 					[]EquipRef `json:"equip_ref" gorm:"default:null"`
	PriorityArrayModel 			PriorityArrayModel `json:"priority_array" gorm:"constraint:OnDelete:CASCADE"`
	PointStore 					PointStore `json:"point_store" gorm:"constraint:OnDelete:CASCADE"`
}


type PointStore struct {
	PointUuid     	string `json:"point_uuid" gorm:"REFERENCES points;not null;default:null;primaryKey"`
	Value  			float64 `json:"value"`
}

type PriorityArrayModel struct {
	PointUuid     	string `json:"point_uuid" gorm:"REFERENCES points;not null;default:null;primaryKey"`
	P1  			string `json:"_1" gorm:"default:null"`
	P2  			string `json:"_2" gorm:"default:null"`
	P3  			string `json:"_3" gorm:"default:null"`
	P4  			string `json:"_4" gorm:"default:null"`
	P5  			string `json:"_5" gorm:"default:null"`
	P6  			string `json:"_6" gorm:"default:null"`
	P7  			string `json:"_7" gorm:"default:null"`
	P8  			string `json:"_8" gorm:"default:null"`
	P9  			string `json:"_9" gorm:"default:null"`
	P10  			string `json:"_10" gorm:"default:null"`
	P11  			string `json:"_11" gorm:"default:null"`
	P12  			string `json:"_12" gorm:"default:null"`
	P13  			string `json:"_13" gorm:"default:null"`
	P14  			string `json:"_14" gorm:"default:null"`
	P15  			string `json:"_15" gorm:"default:null"`
	P16  			string `json:"_16" gorm:"default:null"`
	P17  			string `json:"_17" gorm:"default:null"`
}