package getFieldController

import (
	"net/http"

	"github.com/KadirbekSharau/rentykz-backend/models"
	"github.com/KadirbekSharau/rentykz-backend/handlers/field/get-field"
	"github.com/KadirbekSharau/rentykz-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FieldController interface {
	FindAll() []models.Field
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type fieldController struct {
	service getFieldService.FieldService
}

var validate *validator.Validate

func New(service getFieldService.FieldService) FieldController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", util.ValidateTitle)
	return &fieldController{
		service: service,
	}
}

func (c *fieldController) FindAll() []models.Field {
	return c.service.FindAll()
}

func (c *fieldController) Save(ctx *gin.Context) error {
	var field models.Field
	err := ctx.ShouldBindJSON(&field)
	if err != nil {
		return err
	}
	err = validate.Struct(field)
	if err != nil {
		return err
	}
	c.service.Save(field)
	return nil
}

func (c *fieldController) ShowAll(ctx *gin.Context) {
	fields := c.service.FindAll()
	data := gin.H{
		"Title" : "Renty",
		"fields" : fields,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}