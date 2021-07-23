package modelcommon

import "time"

type Common struct {
	Name        			string `json:"name" validate:"min=1,max=255"  gorm:"type:varchar(255);unique;not null"`
	Description 			string `json:"description"`
	Id						string `json:"id"`
	IdRef					string `json:"id_ref"`
	Type					string `json:"type"`
	Enable 					bool `json:"enable"`
	Fault					bool  `json:"fault"`
	FaultMessage 			string  `json:"fault_message"`
	EnableHistory 			bool   `json:"history_enable"`
	CreatedAt 				time.Time `json:"created_on"`
	UpdatedAt 				time.Time  `json:"updated_on"`

}
