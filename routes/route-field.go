package routes

import (
	createFieldController "github.com/KadirbekSharau/rentykz-backend/controllers/field/create-field"
	getFieldController "github.com/KadirbekSharau/rentykz-backend/controllers/field/get-field"
	getFieldsController "github.com/KadirbekSharau/rentykz-backend/controllers/field/get-fields"
	createFieldHandler "github.com/KadirbekSharau/rentykz-backend/handlers/field/create-field"
	getFieldHandler "github.com/KadirbekSharau/rentykz-backend/handlers/field/get-field"
	getFieldsHandler "github.com/KadirbekSharau/rentykz-backend/handlers/field/get-fields"
	"github.com/KadirbekSharau/rentykz-backend/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Roles: { "user": 1, "admin": 2, "organization":3, "moderator":4 } 

/* @description All field routes */
func InitFieldRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		getFieldsRepository getFieldsController.Repository = getFieldsController.NewGetFieldsRepository(db)
		getFieldsService getFieldsController.Service = getFieldsController.NewGetFieldsService(getFieldsRepository)
		handlerGetFields getFieldsHandler.Handler = getFieldsHandler.NewGetFieldsHandler(getFieldsService)

		createFieldRepository createFieldController.Repository = createFieldController.NewRepositoryCreate(db)
		createFieldService createFieldController.Service = createFieldController.NewCreateFieldService(createFieldRepository)
		handlerCreateField createFieldHandler.Handler = createFieldHandler.NewCreateFieldHandler(createFieldService)

		getFieldRepository getFieldController.Repository = getFieldController.NewRepository(db)
		getFieldService getFieldController.Service = getFieldController.NewService(getFieldRepository)
		handlerGetField getFieldHandler.Handler = getFieldHandler.NewHandler(getFieldService)
	)

	groupRoute := route.Group("/api/v1/field")
	groupRoute.GET("", handlerGetFields.GetFieldsHandler)
	groupRoute.GET("/sporttype/:id/fields", handlerGetFields.GetFieldsBySportTypeIdHandler)
	groupRoute.GET("/organization/:id/fields", middlewares.Auth(3), handlerGetFields.GetFieldsByOrganizationIdHandler)
	groupRoute.GET("/moderator/:id/fields", middlewares.Auth(4), handlerGetFields.GetFieldsByModeratorIdHandler)
	groupRoute.GET("/city/:id/fields", handlerGetFields.GetFieldsByCityIdHandler)
	groupRoute.GET("/:id", handlerGetField.GetFieldByIdHandler)
	groupRoute.POST("", handlerCreateField.CreateFieldHandler)
}