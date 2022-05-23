package getCitiesHandler

import (
	"net/http"

	"github.com/KadirbekSharau/rentykz-backend/controllers/city-controllers/get-cities"
	"github.com/KadirbekSharau/rentykz-backend/util"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetCitiesHandler(ctx *gin.Context)
}

type handler struct {
	service getCitiesController.Service
}

func NewHandler(service getCitiesController.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetCitiesHandler(ctx *gin.Context) {

	fields, err := h.service.GetCitiesService()

	switch err {

	case "RESULTS_CITY_NOT_FOUND_404":
		util.APIResponse(ctx, "Cities data is not exists", http.StatusConflict, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Results Cities data successfully", http.StatusOK, http.MethodPost, fields)
	}
}