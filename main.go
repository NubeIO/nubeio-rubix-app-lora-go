package main

import (
	"github.com/NubeIO/nubeio-rubix-app-lora-go/controller/networks"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/setup"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/rest"
	"log"
)


//func init() {
//	database.Setup()
//	helpers.DisableLogging(false)
//}
// @title GO Nube API
// @version 1.0
// @description nube api docs
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
func main() {

	err := setup.InitMQTT();if err != nil {
		log.Println(err)
		return
	}
	db, err := setup.InitDB();if err != nil {
		log.Println(err)
		return
	}

	app := rest.New(3)
	app.Controller(networks.New(db))
	err = app.Run(":1920");if err != nil {
		log.Println(err)
		return
	}


}
