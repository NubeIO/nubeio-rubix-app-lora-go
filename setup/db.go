package setup

import (
	modeldevices "github.com/NubeIO/nubeio-rubix-app-lora-go/model/devices"
	modelnetworks "github.com/NubeIO/nubeio-rubix-app-lora-go/model/networks"
	modelpoints "github.com/NubeIO/nubeio-rubix-app-lora-go/model/points"
	"github.com/NubeIO/nubeio-rubix-lib-sqlite-go/pkg/database"
	"github.com/NubeIO/nubeio-rubix-lib-sqlite-go/sql_config"
	"gorm.io/gorm"
	"log"
)

func InitDB(DisableLogging bool) (*gorm.DB, error) {
	var args sql_config.Params
	args.UseConfigFile = false
	var config sql_config.Database
	config.DbName = "test.db"
	config.DbPath = "./"
	config.Logging = !DisableLogging

	err := sql_config.SetSqliteConfig(config, args); if err != nil {
		log.Println(err)
		return nil, err
	}
	var network []modelnetworks.Network
	var device []modeldevices.Device
	var point []modelpoints.Point
	var pointStore []modelpoints.PointStore
	var priorityArrayModel []modelpoints.PriorityArrayModel
	var models = []interface{}{
		&network,  &device,  &point, &pointStore, &priorityArrayModel,
	}

	err = database.SetupDB(models); if err != nil {
		log.Println( err)
		return nil, err
	}
	var db = database.GetDB()
	return db, err


}