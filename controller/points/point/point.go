package point

import (
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

var db *gorm.DB

var deviceModel []modeldevices.Device
var pointModel []modelpoints.Point
//var pointStore []modelpoints.PointStore
var childTable = "PointStores"


func New(_db *gorm.DB) rest.IController {
	c := rest.Controller("api/point")
	c.POST("/", create)
	c.SUB("/id")
		c.GET("/", get)
		c.PATCH("/", update)
		c.DELETE("/", _delete)
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



func resolveParameter(ctx *gin.Context) (query string, parameter string){
	_uuid := resolveUUID(ctx)
	id := resolveID(ctx)
	if _uuid != "" {
		return  "uuid = ? ", _uuid
	} else if id != ""{
		return  "id = ? ", id
	}
	return  "uuid = ? ", _uuid
}

func get(ctx *gin.Context) rest.IResponse {
	query, id := resolveParameter(ctx)

	withChildren := withChild(ctx)
	if withChildren {
		_query := db.Where(query, id).Preload(childTable).First(&pointModel);if _query.Error != nil {
			fmt.Println(3333)
			return response.NotFound("not found")
		}
		fmt.Println(444444)
		return response.Data(pointModel)
	} else {
		fmt.Println(55555, "query", query)
		_query := db.Where(query, id).First(&pointModel);if _query.Error != nil {
			fmt.Println(6666)
			return response.NotFound("not found")
		}
		fmt.Println(77777)
		return response.Data(pointModel)
	}
}

//func get2(ctx *gin.Context, query string, field string, withChildren bool) (pointModel,  error) {
//
//	if withChildren {
//		_query := db.Where(query, field).Preload(childTable).First(&deviceModel);if _query.Error != nil {
//			return pointModel, _query.Error
//		}
//		return pointModel, nil
//	} else {
//		_query := db.Where(query, field).First(&pointModel);if _query.Error != nil {
//			return response.BadEntity(_query.Error.Error())
//		}
//		return response.Data(pointModel)
//	}
//}


func update (ctx *gin.Context) rest.IResponse {
	body, _ := getBODY(ctx)
	_uuid := resolveUUID(ctx)
	query := db.Where("uuid = ?", _uuid).First(&pointModel);if query.Error != nil {
		return response.BadEntity(query.Error.Error())
	}
	query = db.Model(&pointModel).Updates(body);if query.Error != nil {
		return response.BadEntity(query.Error.Error())
	}
	return response.Data(pointModel)
}


func _delete(ctx *gin.Context) rest.IResponse {
	_uuid := resolveUUID(ctx)
	query := db.Where("uuid = ? ", _uuid).Unscoped().Delete(&pointModel) ;if query.Error != nil {
		return response.NotFound("point now found")
	}
	r := query.RowsAffected
	if r == 0 {
		return response.NotFound("point now found")
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


func resolveUUID(ctx *gin.Context) string {
	id := ctx.Query("uuid")
	return id
}

func resolveID(ctx *gin.Context) string {
	id := ctx.Query("id")
	return id
}