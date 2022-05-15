package routes

import (
	createCityController "github.com/KadirbekSharau/rentykz-backend/controllers/city-controllers/create-city"
	createCityHandler "github.com/KadirbekSharau/rentykz-backend/handlers/city/create-city"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All field routes */
func InitCityRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		createCityRepository createCityController.Repository = createCityController.NewCreateCityRepository(db)
		createCityService createCityController.Service = createCityController.NewCreateCityService(createCityRepository)
		handlerCreateCity createCityHandler.Handler = createCityHandler.NewCreateCityHandler(createCityService)
	)

	groupRoute := route.Group("/api/v1")
	//groupRoute.GET("/field", handlerGetFields.GetFieldsHandler)
	groupRoute.POST("/city", handlerCreateCity.CreateCityHandler)
}