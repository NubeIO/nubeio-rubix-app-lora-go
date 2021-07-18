package networks

import (
	"github.com/NubeIO/nubeio-rubix-app-lora-go/model/networks"
	"github.com/NubeIO/nubeio-rubix-app-lora-go/response"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/rest"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

func New(db *gorm.DB) rest.IController {
	c := rest.Controller("api/networks")
	c.GET("/", getTODOs)
	c.POST("/", createTODO)
	//c.GET("/", getTODOs)
	//c.POST("/", createTODO)
	//c.SUB("/:id").
	//	GET("/", getTODO).
	//	PUT("/", updateTODO).
	//	DELETE("/", deleteTODO)
	DB = db
	return c
}

func createTODO(ctx *gin.Context) rest.IResponse {
	dto, err := getDTO(ctx)
	dto, err = checkAddNetwork(dto)
	if err != nil {
		return response.BadEntity(err.Error())
	}
	if err = DB.Create(&dto).Error; err != nil {
		return response.BadEntity(err.Error())
	}
	return response.Created(dto.Uuid)
}


func getTODOs(c *gin.Context) rest.IResponse {
	var args rest.Args
	var at = rest.ArgsType
	var ad = rest.ArgsDefault
	args.WithChildren = c.DefaultQuery(at.WithChildren, ad.WithChildren)
	withChildren, _ := rest.WithChildren(args.WithChildren)
	//aa := c.Query(at.Order)
	var items []model_networks.Network
	if withChildren { // drop child to reduce json size
		query := DB.Preload("Device").Find(&items)
		if query.Error != nil {
			return response.Data(items)
		}
		return response.Data(items)
	} else {
		query := DB.Find(&items)
		if query.Error != nil {
			return response.Data(items)
		}
		return response.Data(items)
	}
}


func getDTO(ctx *gin.Context) (dto *model_networks.Network, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func checkAddNetwork(data *model_networks.Network) (*model_networks.Network, error) {
	if err := validator.Validate(data); err != nil {
		return nil, err
	}
	data.Uuid, _ = uuid.MakeUUID()
	return data, nil
}
