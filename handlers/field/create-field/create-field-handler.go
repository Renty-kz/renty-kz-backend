package createFieldHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	createField "github.com/KadirbekSharau/rentykz-backend/controllers/field-controllers/create-field"
	util "github.com/KadirbekSharau/rentykz-backend/util"
)

type Handler interface {
	CreateFieldHandler(ctx *gin.Context)
}

type handler struct {
	service createField.Service
}

func NewCreateFieldHandler(service createField.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateFieldHandler(ctx *gin.Context) {

	var input createField.InputCreateField
	ctx.ShouldBindJSON(&input)

	config := util.ErrorConfig{
		Options: []util.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "Name",
				Message: "name is required on body",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	_, errCreateField := h.service.CreateFieldService(&input)

	switch errCreateField {
	case "CREATE_FIELD_CONFLICT_409":
		util.APIResponse(ctx, "Name field already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "CREATE_FIELD_FAILED_403":
		util.APIResponse(ctx, "Create new field account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		util.APIResponse(ctx, "Create new field account successfully", http.StatusCreated, http.MethodPost, nil)
	}
}