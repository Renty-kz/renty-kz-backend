package updateOrganizationHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/KadirbekSharau/rentykz-backend/controllers/organization/update-organization"
	util "github.com/KadirbekSharau/rentykz-backend/util"
)

type Handler interface {
	ActivateOrganizationHandler(ctx *gin.Context)
	InactivateOrganizationHandler(ctx *gin.Context)
}

type handler struct {
	service updateOrganizationController.Service
}

func NewHandler(service updateOrganizationController.Service) *handler {
	return &handler{service: service}
}

/* Activate Organization Handler */
func (h *handler) ActivateOrganizationHandler(ctx *gin.Context) {

	var input updateOrganizationController.InputActivateOrganization
	ctx.ShouldBindJSON(&input)

	input.ID = ctx.Params.ByName("id")

	errResponse, errCount := util.GoValidator(&input, ActivationOrganizationConfig.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	res, errUpdateOrganization := h.service.ActivateOrganizationByIdService(&input)

	ErrActivateOrganizationHandler(res, ctx, errUpdateOrganization)
}

/* Inactivate Organization Handler */
func (h *handler) InactivateOrganizationHandler(ctx *gin.Context) {

	var input updateOrganizationController.InputActivateOrganization
	ctx.ShouldBindJSON(&input)

	input.ID = ctx.Params.ByName("id")

	errResponse, errCount := util.GoValidator(&input, ActivationOrganizationConfig.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	res, errUpdateOrganization := h.service.InactivateOrganizationByIdService(&input)

	ErrActivateOrganizationHandler(res, ctx, errUpdateOrganization)
}