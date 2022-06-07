package getFieldHandler

import (
	"net/http"

	getFieldController "github.com/KadirbekSharau/rentykz-backend/controllers/field/get-field"
	"github.com/KadirbekSharau/rentykz-backend/util"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetFieldByIdHandler(ctx *gin.Context)
}

type handler struct {
	service getFieldController.Service
}

func NewHandler(service getFieldController.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetFieldByIdHandler(ctx *gin.Context) {
	var input getFieldController.InputField

	input.ID = ctx.Params.ByName("id")

	config := util.ErrorConfig{
		Options: []util.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "ID",
				Message: "id is required on param",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	resultField, errResultField := h.service.GetFieldByIdService(&input)

	switch errResultField {

	case "RESULT_STUDENT_NOT_FOUND_404":
		util.APIResponse(ctx, "Field data is not exist or deleted", http.StatusNotFound, http.MethodGet, nil)
		return

	default:
		util.APIResponse(ctx, "Result Field data successfully", http.StatusOK, http.MethodGet, resultField)
	}
}