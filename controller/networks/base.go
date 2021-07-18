package networks

import "gorm.io/gorm"

type Controller struct {
	DB *gorm.DB
}
