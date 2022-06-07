package getOrganizationsHandler

import (
	"github.com/KadirbekSharau/rentykz-backend/controllers/organization/get-organizations"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetActiveOrganizationsHandler(ctx *gin.Context)
	GetInactiveOrganizationsHandler(ctx *gin.Context)
	GetAllOrganizationsHandler(ctx *gin.Context)
}

type handler struct {
	service getOrganizationsController.Service
}

func NewHandler(service getOrganizationsController.Service) *handler {
	return &handler{service: service}
}

/* Get All Organizations Handler */
func (h *handler) GetAllOrganizationsHandler(ctx *gin.Context) {

	organizations, err := h.service.GetAllOrganizationsService()

	ErrAllOrganizationsHandler(organizations, ctx, err)
}

/* Get Active Organizations Handler */
func (h *handler) GetActiveOrganizationsHandler(ctx *gin.Context) {

	organizations, err := h.service.GetActiveOrganizationsService()

	ErrActiveOrganizationsHandler(organizations, ctx, err)
}

/* Get Inactive Organizations Handler */
func (h *handler) GetInactiveOrganizationsHandler(ctx *gin.Context) {

	organizations, err := h.service.GetInactiveOrganizationsService()

	ErrInactiveOrganizationsHandler(organizations, ctx, err)
}