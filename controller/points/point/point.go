package point

import (
	"errors"
	"fmt"
	modeldevices "github.com/NubeIO/nubeio-rubix-app-lora-go/model/devices"
	modelpoints "github.com/NubeIO/nubeio-rubix-app-lora-go/model/points"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/response"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/rest"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)


const (
	ParameterNotSupported = "search parameter is not supported"
)

var db *gorm.DB

var deviceModel []modeldevices.Device
var pointModel []modelpoints.Point
//var pointStore []modelpoints.PointStore
var childTable = "PointStores"


func New(_db *gorm.DB) rest.IController {
	c := rest.Controller("api/point")
	c.POST("/", create)
	c.SUB("/:uuid").
		GET("/", get).
		PATCH("/", update).
		DELETE("/", _delete)
	db = _db
	return c
}

func create(ctx *gin.Context) rest.IResponse {
	body, err := getBODY(ctx); if err != nil {
		return response.BadEntity(err.Error())
	}
	deviceUUID := body.DeviceUuid
	query := db.Where("uuid = ? ", deviceUUID).First(&deviceModel);if query.Error != nil {
		return response.BadEntity(query.Error.Error())
	}
	body, err = validate(body); if err != nil {
		return response.BadEntity(err.Error())
	}
	if err = db.Create(&body).Error; err != nil {
		return response.BadEntity(err.Error())
	}
	_pointModel := new(modelpoints.PointStore)
	_priorityArrayModel := new(modelpoints.PriorityArrayModel)
	_pointModel.PointUuid = body.Uuid
	_priorityArrayModel.PointUuid = body.Uuid

	if err = db.Create(&_pointModel).Error; err != nil {
		return response.BadEntity(err.Error())
	}
	if err = db.Create(&_priorityArrayModel).Error; err != nil {
		return response.BadEntity(err.Error())
	}
	return response.Created(body.Uuid)
}


func get(ctx *gin.Context) rest.IResponse {
	qString, id, err := resolveParameter(ctx); if err != nil {
		return response.NotFound("query not parameter supported")
	}
	withChildren := withChild(ctx)
	if withChildren {
		query := db.Where(qString, id).Preload(childTable).First(&pointModel);if query.Error != nil {
			return response.NotFound("not found")
		}
		return response.Data(pointModel)
	} else {
		query := db.Where(qString, id).First(&pointModel);if query.Error != nil {
			return response.NotFound(query.Error.Error())
		}
		return response.Data(pointModel)
	}
}


func update (ctx *gin.Context) rest.IResponse {
	body, _ := getBODY(ctx)
	uid := resolveUUID(ctx)
	query := db.Where("uuid = ?", uid).First(&pointModel);if query.Error != nil {
		return response.BadEntity(query.Error.Error())
	}
	query = db.Model(&pointModel).Updates(body);if query.Error != nil {
		return response.BadEntity(query.Error.Error())
	}
	return response.Data(pointModel)
}


func _delete(ctx *gin.Context) rest.IResponse {
	uid := resolveUUID(ctx)
	query := db.Where("uuid = ? ", uid).Unscoped().Delete(&pointModel) ;if query.Error != nil {
		return response.NotFound(query.Error.Error())
	}
	r := query.RowsAffected
	if r == 0 {
		return response.NotFound(query.Error.Error())
	} else {
		return response.OKWithMessage("point deleted")
	}
}


func getBODY(ctx *gin.Context) (dto *modelpoints.Point, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func validate(data *modelpoints.Point) (*modelpoints.Point, error) {
	if err := validator.Validate(data); err != nil {
		return nil, err
	}
	data.Uuid, _ = uuid.MakeUUID()
	return data, nil
}

func withChild(ctx *gin.Context) bool {
	var args rest.Args
	var at = rest.ArgsType
	var ad = rest.ArgsDefault
	args.WithChildren = ctx.DefaultQuery(at.WithChildren, ad.WithChildren)
	withChildren, _ := rest.WithChildren(args.WithChildren)
	return withChildren
}


func resolveParameter(ctx *gin.Context) (query string, parameter string, err error){
	id := resolveID(ctx)
	uid := resolveUUID(ctx)
	name := resolveName(ctx)
	objectType := resolveObject(ctx)
	deviceName := resolveDeviceName(ctx)
	networkName := resolveNetworkName(ctx)
	fmt.Println(deviceName, networkName, 99999999)
	if id != ""{
		return  "id = ? ", id, nil
	} else if id != ""{
		return  "uuid = ? ", uid, nil
	} else if id != ""{
		return  "name = ? ", name, nil
	} else if id != ""{
		return  "object_type = ? ", objectType, nil
	}
	return  "", "", errors.New(ParameterNotSupported)
}

func resolveObject(ctx *gin.Context) string {
	id := ctx.Query("objectType")
	return id
}

func resolveUUID(ctx *gin.Context) string {
	id := ctx.Query("uuid")
	//id := ctx.Query("uuid")
	//network := ctx.Param("network")
	fmt.Println(666666, ctx.Query("uuid"), 666666 ,ctx.Param("uuid"))

	return id
}

func resolveName(ctx *gin.Context) string {
	id := ctx.Query("name")
	return id
}
func resolveDeviceName(ctx *gin.Context) string {
	id := ctx.Query("deviceName")
	return id
}

func resolveNetworkName(ctx *gin.Context) string {
	id := ctx.Query("networkName")
	return id
}

func resolveID(ctx *gin.Context) string {
	id := ctx.Query("id")
	return id
}
