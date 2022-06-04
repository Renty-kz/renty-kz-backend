package routes

import (
	createCityController "github.com/KadirbekSharau/rentykz-backend/controllers/city-controllers/create-city"
	getCitiesController "github.com/KadirbekSharau/rentykz-backend/controllers/city-controllers/get-cities"
	createCityHandler "github.com/KadirbekSharau/rentykz-backend/handlers/city/create-city"
	getCitiesHandler "github.com/KadirbekSharau/rentykz-backend/handlers/city/get-cities"
	"github.com/KadirbekSharau/rentykz-backend/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All field routes */
func InitCityRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		createCityRepository createCityController.Repository = createCityController.NewCreateCityRepository(db)
		createCityService createCityController.Service = createCityController.NewCreateCityService(createCityRepository)
		handlerCreateCity createCityHandler.Handler = createCityHandler.NewCreateCityHandler(createCityService)

		getCitiesRepository getCitiesController.Repository = getCitiesController.NewRepository(db)
		getCitiesService getCitiesController.Service = getCitiesController.NewService(getCitiesRepository)
		handlerGetCities getCitiesHandler.Handler = getCitiesHandler.NewHandler(getCitiesService)
	)

	groupRoute := route.Group("/api/v1").Use(middlewares.Auth(2))
	groupRoute.GET("/city", handlerGetCities.GetCitiesHandler)
	groupRoute.POST("/city", handlerCreateCity.CreateCityHandler)
}