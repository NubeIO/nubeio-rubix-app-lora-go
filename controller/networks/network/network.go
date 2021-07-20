package network

import (
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
var childTable = "Device"

func New(_db *gorm.DB) rest.IController {
	c := rest.Controller("api/network")
	c.POST("/", create)
	c.SUB("/uuid")
		c.GET("/", get)
		c.PATCH("/", update)
		c.DELETE("/", _delete)
	db = _db
	return c
}

func create(ctx *gin.Context) rest.IResponse {
	dto, err := getBODY(ctx)
	dto, err = addNetworkValidate(dto)
	if err != nil {
		return response.BadEntity(err.Error())
	}
	if err = db.Create(&dto).Error; err != nil {
		return response.BadEntity(err.Error())
	}
	return response.Created(dto.Uuid)
}


func get(ctx *gin.Context) rest.IResponse {
	_uuid := resolveID(ctx)
	withChildren := withChild(ctx)
	if withChildren {
		query := db.Where("uuid = ? ", _uuid).Preload(childTable).First(&networkModel);if query.Error != nil {
			return response.BadEntity(query.Error.Error())
		}
		return response.Data(networkModel)
	} else {
		query := db.Where("uuid = ? ", _uuid).First(&networkModel);if query.Error != nil {
			return response.BadEntity(query.Error.Error())
		}
		return response.Data(networkModel)
	}
}

func update (ctx *gin.Context) rest.IResponse {
	body, _ := getBODY(ctx)
	_uuid := resolveID(ctx)
	query := db.Where("uuid = ?", _uuid).First(&networkModel);if query.Error != nil {
		return response.BadEntity(query.Error.Error())
	}
	query = db.Model(&networkModel).Updates(body);if query.Error != nil {
		return response.BadEntity(query.Error.Error())
	}
	return response.Data(networkModel)
}

func _delete(ctx *gin.Context) rest.IResponse {
	_uuid := resolveID(ctx)
	query := db.Where("uuid = ? ", _uuid).Unscoped().Delete(&networkModel) ;if query.Error != nil {
		return response.NotFound("network now found")
	}
	r := query.RowsAffected
	if r == 0 {
		return response.NotFound("network now found")
	} else {
		return response.OKWithMessage("deleted network")
	}

}


func getBODY(ctx *gin.Context) (dto *modelnetworks.Network, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func addNetworkValidate(data *modelnetworks.Network) (*modelnetworks.Network, error) {
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