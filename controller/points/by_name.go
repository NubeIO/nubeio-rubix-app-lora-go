package points

import (
	modeldevices "github.com/NubeIO/nubeio-rubix-app-lora-go/model/devices"
	modelnetworks "github.com/NubeIO/nubeio-rubix-app-lora-go/model/networks"
	modelpoints "github.com/NubeIO/nubeio-rubix-app-lora-go/model/points"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/response"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/rest"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ByName(_db *gorm.DB) rest.IController {
	c := rest.Controller("api/points/name/:network/:device/:point")
	c.GET("/", getByName)
	db = _db
	return c
}

var pointModel []modelpoints.Point
var networkModel []modelnetworks.Network
var deviceModel []modeldevices.Device



func getByName(ctx *gin.Context) rest.IResponse {
	network := ctx.Param("network")
	device := ctx.Param("device")
	point := ctx.Param("point")
	db.Table("points").Select("points.name AS name, points.uuid AS uuid, points.description AS description").
		Joins("JOIN devices").
		Joins("JOIN networks").
		Where("points.name = ? AND devices.name = ? AND networks.name = ? ",point, device, network).First(&pointModel)
	return response.Data(pointModel)

}