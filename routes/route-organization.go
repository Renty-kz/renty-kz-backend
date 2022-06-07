package routes

import (
	getOrganizationsController "github.com/KadirbekSharau/rentykz-backend/controllers/organization/get-organizations"
	getOrganizationsHandler "github.com/KadirbekSharau/rentykz-backend/handlers/organization/get-organizations"
	updateOrganizationController "github.com/KadirbekSharau/rentykz-backend/controllers/organization/update-organization"
	updateOrganizationHandler "github.com/KadirbekSharau/rentykz-backend/handlers/organization/update-organization"
	"github.com/KadirbekSharau/rentykz-backend/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Organization Routes */
func InitOrganizationRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		getOrganizationsRepository getOrganizationsController.Repository = getOrganizationsController.NewRepository(db)
		getOrganizationsService getOrganizationsController.Service = getOrganizationsController.NewService(getOrganizationsRepository)
		handlerGetOrganizations getOrganizationsHandler.Handler = getOrganizationsHandler.NewHandler(getOrganizationsService)

		updateOrganizationRepository updateOrganizationController.Repository = updateOrganizationController.NewRepository(db)
		updateOrganizationsService updateOrganizationController.Service = updateOrganizationController.NewService(updateOrganizationRepository)
		handlerUpdateOrganization updateOrganizationHandler.Handler = updateOrganizationHandler.NewHandler(updateOrganizationsService)
	)

	groupRoute := route.Group("/api/v1/organization")
	groupRoute.GET("", middlewares.Auth(2), handlerGetOrganizations.GetAllOrganizationsHandler)
	groupRoute.GET("/active", middlewares.Auth(2), handlerGetOrganizations.GetActiveOrganizationsHandler)
	groupRoute.GET("/inactive", middlewares.Auth(2), handlerGetOrganizations.GetInactiveOrganizationsHandler)
	groupRoute.PUT("/:id/activate", middlewares.Auth(2), handlerUpdateOrganization.ActivateOrganizationHandler)
	groupRoute.PUT("/:id/inactivate", middlewares.Auth(2), handlerUpdateOrganization.InactivateOrganizationHandler)
}