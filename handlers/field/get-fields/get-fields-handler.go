package getFieldsHandler

import (
	"net/http"

	"github.com/KadirbekSharau/rentykz-backend/util"
	"github.com/KadirbekSharau/rentykz-backend/controllers/field-controllers/get-fields"
	"github.com/gin-gonic/gin"
)

type GetFieldsHandler interface {
	GetFieldsHandler(ctx *gin.Context)
}

type handler struct {
	service getFieldsController.Service
}

func NewGetFieldsHandler(service getFieldsController.Service) *handler {
	return &handler{service: service}
}

func (h *handler)GetFieldsHandler(ctx *gin.Context) {

	fields, err := h.service.GetFieldsService()

	switch err {

	case "RESULTS_STUDENT_NOT_FOUND_404":
		util.APIResponse(ctx, "Fields data is not exists", http.StatusConflict, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Results Fields data successfully", http.StatusOK, http.MethodPost, fields)
	}
}