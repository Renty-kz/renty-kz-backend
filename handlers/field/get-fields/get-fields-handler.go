package getFieldsHandler

import (
	"net/http"

	"github.com/KadirbekSharau/rentykz-backend/util"
	"github.com/KadirbekSharau/rentykz-backend/controllers/field/get-fields"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetFieldsHandler(ctx *gin.Context)
	GetFieldsByCityIdHandler(ctx *gin.Context)
	GetFieldsBySportTypeIdHandler(ctx *gin.Context)
	GetFieldsByOrganizationIdHandler(ctx *gin.Context)
	GetFieldsByModeratorIdHandler(ctx *gin.Context)
}

type handler struct {
	service getFieldsController.Service
}

func NewGetFieldsHandler(service getFieldsController.Service) *handler {
	return &handler{service: service}
}

/* Get Fields Handler */
func (h *handler) GetFieldsHandler(ctx *gin.Context) {

	fields, err := h.service.GetFieldsService()

	switch err {

	case "RESULTS_FIELD_NOT_FOUND_404":
		util.APIResponse(ctx, "Fields data is not exists", http.StatusConflict, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Results Fields data successfully", http.StatusOK, http.MethodPost, fields)
	}
}

/* Get Fields By City ID Handler */
func (h *handler) GetFieldsByCityIdHandler(ctx *gin.Context) {
	var input getFieldsController.InputFieldByCityId

	input.CityID = ctx.Params.ByName("id")

	errResponse, errCount := util.GoValidator(&input, configCityId.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	fields, err := h.service.GetFieldsByCityIdService(&input)
	
	ErrFieldsByCityIdHandler(fields, ctx, err)
}

/* Get Fields By Sport Type ID Handler */
func (h *handler) GetFieldsBySportTypeIdHandler(ctx *gin.Context) {
	var input getFieldsController.InputFieldsBySportTypeId

	input.SportTypeID = ctx.Params.ByName("id")

	errResponse, errCount := util.GoValidator(&input, configSportTypeId.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	fields, err := h.service.GetFieldsBySportTypeIdService(&input)
	
	ErrFieldsBySportTypeIdHandler(fields, ctx, err)
}

/* Get Fields By Organization ID Handler */
func (h *handler) GetFieldsByOrganizationIdHandler(ctx *gin.Context) {
	var input getFieldsController.InputFieldsByOrganizationId

	input.OrganizationID = ctx.Params.ByName("id")

	errResponse, errCount := util.GoValidator(&input, configSportTypeId.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	fields, err := h.service.GetFieldsByOrganizationIdService(&input)
	
	ErrFieldsByOrganizationIdHandler(fields, ctx, err)
}

/* Get Fields By Moderator and Organization ID Handler */
func (h *handler) GetFieldsByModeratorIdHandler(ctx *gin.Context) {
	var input getFieldsController.InputFieldsByModeratorId

	input.ModeratorID = ctx.Params.ByName("id")

	errResponse, errCount := util.GoValidator(&input, configSportTypeId.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	fields, err := h.service.GetFieldsByModeratorIdService(&input)
	
	ErrFieldsByOrganizationIdHandler(fields, ctx, err)
}