package routes

import (
	createFieldController "github.com/KadirbekSharau/rentykz-backend/controllers/field-controllers/create-field"
	getFieldController "github.com/KadirbekSharau/rentykz-backend/controllers/field-controllers/get-field"
	getFieldsController "github.com/KadirbekSharau/rentykz-backend/controllers/field-controllers/get-fields"
	createFieldHandler "github.com/KadirbekSharau/rentykz-backend/handlers/field/create-field"
	getFieldsHandler "github.com/KadirbekSharau/rentykz-backend/handlers/field/get-fields"
	getFieldHandler "github.com/KadirbekSharau/rentykz-backend/handlers/field/get-field"
	"github.com/KadirbekSharau/rentykz-backend/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

	groupRoute := route.Group("/api/v1")
	groupRoute.GET("/field", handlerGetFields.GetFieldsHandler)
	groupRoute.GET("/sporttype/:id/fields", handlerGetFields.GetFieldsBySportTypeIdHandler)
	groupRoute.GET("/organization/:id/fields", handlerGetFields.GetFieldsByOrganizationIdHandler)/* needs middleware */
	groupRoute.GET("/moderator/:id/fields", handlerGetFields.GetFieldsByModeratorIdHandler)/* needs middleware */
	groupRoute.GET("/city/:id/fields", handlerGetFields.GetFieldsByCityIdHandler)
	groupRoute.GET("/field/:id", handlerGetField.GetFieldByIdHandler).Use(middlewares.Auth())
	groupRoute.POST("/field", handlerCreateField.CreateFieldHandler).Use(middlewares.Auth())
}