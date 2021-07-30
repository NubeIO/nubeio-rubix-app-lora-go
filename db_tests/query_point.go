package main

import (
	"fmt"
	modeldevices "github.com/NubeIO/nubeio-rubix-app-lora-go/model/devices"
	modelnetworks "github.com/NubeIO/nubeio-rubix-app-lora-go/model/networks"
	modelpoints "github.com/NubeIO/nubeio-rubix-app-lora-go/model/points"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}


var pointModel []modelpoints.Point
var networkModel []modelnetworks.Network
var deviceModel []modeldevices.Device
var priorityArrayModel []modelpoints.PriorityArrayModel
var pointStore []modelpoints.PointStore
var loadChildTable = "PointStore"
var loadPriorityArrayModel = "PriorityArrayModel"

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	uuid := "3884d433d2de495f"

	// QUERY
	//db.Where("uuid = ?", uuid).Preload(childTable).Preload(PriorityArrayModel).First(&pointModel)
	//log.Printf("%+v",pointModel)

	// UPDATE NAME
	point := new(modelpoints.Point)
	point.Name = "bbb"
	query := db.Where("uuid = ?", uuid).Preload(loadChildTable).Preload(loadPriorityArrayModel).Updates(point);if query.Error != nil {
		fmt.Println(err)

	}

	log.Printf("%+v",point.PriorityArrayModel)

	//for i, s := range point.PriorityArrayModel {
	//	fmt.Println(i, s)
	//}

	// UPDATE Priority
	pri := new(modelpoints.PriorityArrayModel)
	pri.P1 = "12.4"
	db.Where("point_uuid = ?", uuid).Preload(loadChildTable).Preload(loadPriorityArrayModel).Updates(pri);if query.Error != nil {
		fmt.Println(err)

	}
	// UPDATE PointStore
	store := new(modelpoints.PointStore)
	store.Value = "11"
	db.Where("point_uuid = ?", uuid).Preload(loadChildTable).Preload(loadPriorityArrayModel).Updates(store);if query.Error != nil {
		fmt.Println(err)

	}
	log.Printf("%+v",pointModel)




}