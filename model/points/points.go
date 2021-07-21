package modelpoints

import modelcommon "github.com/NubeIO/nubeio-rubix-app-lora-go/model/common"


type CommonPoint struct {
	Writeable 		bool   `json:"writeable"`
	Cov  			float64 `json:"cov"`
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
	//EquipRef 					[]EquipRef `json:"equip_ref" gorm:"constraint:OnDelete:CASCADE"`
	PriorityArrayModel 			PriorityArrayModel `json:"priority_array" gorm:"constraint:OnDelete:CASCADE"`
	PointStore 					PointStore `json:"point_store" gorm:"constraint:OnDelete:CASCADE"`

}


type PointStore struct {
	PointUuid     	string `json:"point_uuid" gorm:"REFERENCES points;not null;default:null;primaryKey"`
	Value  			float64 `json:"value"`
}

type PriorityArrayModel struct {
	PointUuid     	string `json:"point_uuid" gorm:"REFERENCES points;not null;default:null;primaryKey"`
	P1  			float64 `json:"_1" gorm:"default:null"`
	P2  			float64 `json:"_2" gorm:"default:null"`
	P3  			float64 `json:"_3" gorm:"default:null"`
	P4  			float64 `json:"_4" gorm:"default:null"`
	P5  			float64 `json:"_5" gorm:"default:null"`
	P6  			float64 `json:"_6" gorm:"default:null"`
	P7  			float64 `json:"_7" gorm:"default:null"`
	P8  			float64 `json:"_8" gorm:"default:null"`
	P9  			float64 `json:"_9" gorm:"default:null"`
	P10  			float64 `json:"_10" gorm:"default:null"`
	P11  			float64 `json:"_11" gorm:"default:null"`
	P12  			float64 `json:"_12" gorm:"default:null"`
	P13  			float64 `json:"_13" gorm:"default:null"`
	P14  			float64 `json:"_14" gorm:"default:null"`
	P15  			float64 `json:"_15" gorm:"default:null"`
	P16  			float64 `json:"_16" gorm:"default:null"`
	P17  			float64 `json:"_17" gorm:"default:null"`
}