package routes

import (
	createCityController "github.com/KadirbekSharau/rentykz-backend/controllers/city/create-city"
	getCitiesController "github.com/KadirbekSharau/rentykz-backend/controllers/city/get-cities"
	getCityController "github.com/KadirbekSharau/rentykz-backend/controllers/city/get-city"
	createCityHandler "github.com/KadirbekSharau/rentykz-backend/handlers/city/create-city"
	getCitiesHandler "github.com/KadirbekSharau/rentykz-backend/handlers/city/get-cities"
	getCityHandler "github.com/KadirbekSharau/rentykz-backend/handlers/city/get-city"
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

		getCityRepository getCityController.Repository = getCityController.NewRepository(db)
		getCityService getCityController.Service = getCityController.NewService(getCityRepository)
		handlerGetCity getCityHandler.Handler = getCityHandler.NewHandler(getCityService)
	)

	groupRoute := route.Group("/api/v1/city")
	groupRoute.GET("", handlerGetCities.GetCitiesHandler)
	groupRoute.GET("/:id", handlerGetCity.GetCityByIdHandler)
	groupRoute.POST("", middlewares.Auth(2), handlerCreateCity.CreateCityHandler)
}