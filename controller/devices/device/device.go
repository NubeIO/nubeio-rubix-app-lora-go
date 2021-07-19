package device

import (
	modeldevices "github.com/NubeIO/nubeio-rubix-app-lora-go/model/devices"
	modelnetworks "github.com/NubeIO/nubeio-rubix-app-lora-go/model/networks"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/response"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/rest"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

var db *gorm.DB
var networkModel []modelnetworks.Network
var deviceModel []modeldevices.Device
var childTable = "Points"

func New(_db *gorm.DB) rest.IController {
	c := rest.Controller("api/device")
	c.POST("/", create)
	c.SUB("/uuid")
		c.GET("/", get)
		c.PUT("/", update)
		c.DELETE("/", _delete)
	db = _db
	return c
}

func create(ctx *gin.Context) rest.IResponse {
	body, err := getBODY(ctx); if err != nil {
		return response.BadEntity(err.Error())
	}
	networkUUID := body.NetworkUuid
	query := db.Where("uuid = ? ", networkUUID).First(&networkModel);if query.Error != nil {
		return response.BadEntity(query.Error.Error())
	}
	body, err = validate(body); if err != nil {
		return response.BadEntity(err.Error())
	}
	if err = db.Create(&body).Error; err != nil {
		return response.BadEntity(err.Error())
	}
	return response.Created(body.Uuid)
}


func get(ctx *gin.Context) rest.IResponse {
	_uuid := resolveID(ctx)
	withChildren := withChild(ctx)
	if withChildren {
		query := db.Where("uuid = ? ", _uuid).Preload(childTable).First(&networkModel);if query.Error != nil {
			return response.BadEntity(query.Error.Error())
		}
		return response.Data(deviceModel)
	} else {
		query := db.Where("uuid = ? ", _uuid).First(&deviceModel);if query.Error != nil {
			return response.BadEntity(query.Error.Error())
		}
		return response.Data(deviceModel)
	}
}

func update(ctx *gin.Context) rest.IResponse {
	body, _ := getBODY(ctx)
	_uuid := resolveID(ctx)
	query := db.Where("uuid = ?", _uuid).First(&deviceModel);if query.Error != nil {
		return response.BadEntity(query.Error.Error())
	}
	query = db.Model(&deviceModel).Updates(body);if query.Error != nil {
		return response.BadEntity(query.Error.Error())
	}
	return response.Data(deviceModel)
}

func _delete(ctx *gin.Context) rest.IResponse {
	_uuid := resolveID(ctx)
	query := db.Where("uuid = ? ", _uuid).Unscoped().Delete(&deviceModel) ;if query.Error != nil {
		return response.NotFound("device now found")
	}
	r := query.RowsAffected
	if r == 0 {
		return response.NotFound("device now found")
	} else {
		return response.OKWithMessage("device network")
	}

}

func getBODY(ctx *gin.Context) (dto *modeldevices.Device, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func validate(data *modeldevices.Device) (*modeldevices.Device, error) {
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



func resolveID(ctx *gin.Context) string {
	id := ctx.Query("uuid")
	return id
}
