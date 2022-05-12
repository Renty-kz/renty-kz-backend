package routes

import (
	getFieldsController "github.com/KadirbekSharau/rentykz-backend/controllers/field-controllers/get-fields"
	getFieldsHandler "github.com/KadirbekSharau/rentykz-backend/handlers/field/get-fields"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All field routes */
func InitFieldRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		getFieldsRepository getFieldsController.Repository = getFieldsController.NewGetFieldsRepository(db)
		getFieldsService getFieldsController.Service = getFieldsController.NewGetFieldsService(getFieldsRepository)
		handler getFieldsHandler.GetFieldsHandler = getFieldsHandler.NewGetFieldsHandler(getFieldsService)
	)

	groupRoute := route.Group("/api/v1")
	groupRoute.GET("/field", handler.GetFieldsHandler)
}