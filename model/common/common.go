package modelcommon

import "time"


type CommonName struct {
	Name        			string `json:"name" validate:"min=1,max=255"  gorm:"type:varchar(255);unique;not null"`
	Description 			string `json:"description"`
	Id						string `json:"id"`
}

type CommonUUID struct {
	Uuid			string 		`json:"uuid"  gorm:"type:varchar(255);unique;primaryKey"`
}

type Common struct {
	CommonName
	Enable 					bool `json:"enable"`
	Fault					bool  `json:"fault"`
	FaultMessage 			string  `json:"fault_message"`
	EnableHistory 			bool   `json:"history_enable"`
	CreatedAt 				time.Time `json:"created_on"`
	UpdatedAt 				time.Time  `json:"updated_on"`
}

