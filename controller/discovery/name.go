package discovery


import (
	modelpoints "github.com/NubeIO/nubeio-rubix-app-lora-go/model/points"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/response"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/rest"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func New(_db *gorm.DB) rest.IController {
	c := rest.Controller("api/all/name")
	c.GET("/", get)
	db = _db
	return c
}

func get(c *gin.Context) rest.IResponse {
	var args rest.Args
	var at = rest.ArgsType
	var ad = rest.ArgsDefault
	args.WithChildren = c.DefaultQuery(at.WithChildren, ad.WithChildren)
	withChildren, err := rest.WithChildren(args.WithChildren); if err != nil {
		return response.Data(err)
	}
	var items []modelpoints.Point
	if withChildren { // drop child to reduce json size
		query := db.Preload("PointStore").Preload("PriorityArrayModel").Find(&items); if query.Error != nil {
			return response.Data(items)
		}
		return response.Data(items)
	} else {
		query := db.Find(&items);if query.Error != nil {
			return response.Data(items)
		}
		return response.Data(items)
	}
}

