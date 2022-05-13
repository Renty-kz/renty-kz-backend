package routes

import (
	createFieldController "github.com/KadirbekSharau/rentykz-backend/controllers/field-controllers/create-field"
	getFieldsController "github.com/KadirbekSharau/rentykz-backend/controllers/field-controllers/get-fields"
	createFieldHandler "github.com/KadirbekSharau/rentykz-backend/handlers/field/create-field"
	getFieldsHandler "github.com/KadirbekSharau/rentykz-backend/handlers/field/get-fields"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All field routes */
func InitFieldRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		getFieldsRepository getFieldsController.Repository = getFieldsController.NewGetFieldsRepository(db)
		getFieldsService getFieldsController.Service = getFieldsController.NewGetFieldsService(getFieldsRepository)
		handlerGetFields getFieldsHandler.GetFieldsHandler = getFieldsHandler.NewGetFieldsHandler(getFieldsService)

		createFieldRepository createFieldController.Repository = createFieldController.NewRepositoryCreate(db)
		createFieldService createFieldController.Service = createFieldController.NewServiceCreate(createFieldRepository)
		handlerCreateField createFieldHandler.CreateFieldHandler = createFieldHandler.NewCreateFieldHandler(createFieldService)
	)

	groupRoute := route.Group("/api/v1")
	groupRoute.GET("/field", handlerGetFields.GetFieldsHandler)
	groupRoute.POST("/field", handlerCreateField.CreateFieldHandler)
}