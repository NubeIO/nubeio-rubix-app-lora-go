package setup

import (
	model_networks "github.com/NubeIO/nubeio-rubix-app-lora-go/model/networks"
	"github.com/NubeIO/nubeio-rubix-lib-sqlite-go/pkg/database"
	"github.com/NubeIO/nubeio-rubix-lib-sqlite-go/sql_config"
	"gorm.io/gorm"
	"log"
)

func InitDB() (*gorm.DB, error) {
	var args sql_config.Params
	args.UseConfigFile = false
	var config sql_config.Database
	config.DbName = "test.db"
	config.DbPath = "./"
	config.Logging = true

	err := sql_config.SetSqliteConfig(config, args); if err != nil {
		log.Println(err)
		return nil, err
	}
	var network []model_networks.Network
	var device []model_networks.Device
	var models = []interface{}{
		&network,  &device,
	}
	err = database.SetupDB(models); if err != nil {
		log.Println(2222, err)
		return nil, err
	}
	var db = database.GetDB()
	return db, err


}