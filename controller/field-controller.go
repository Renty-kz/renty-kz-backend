package controller

import (
	"net/http"

	"github.com/KadirbekSharau/rentykz-backend/entity"
	"github.com/KadirbekSharau/rentykz-backend/service"
	"github.com/KadirbekSharau/rentykz-backend/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FieldController interface {
	FindAll() []entity.Field
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type fieldController struct {
	service service.FieldService
}

var validate *validator.Validate

func New(service service.FieldService) FieldController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateTitle)
	return &fieldController{
		service: service,
	}
}

func (c *fieldController) FindAll() []entity.Field {
	return c.service.FindAll()
}

func (c *fieldController) Save(ctx *gin.Context) error {
	var field entity.Field
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