package routes

import (
	"net/http"

	getFieldController "github.com/KadirbekSharau/rentykz-backend/controller/field-controllers/get-field"
	getFieldService "github.com/KadirbekSharau/rentykz-backend/service/field/get-field"
	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"
)

/* @description All field routes */
func InitFieldRoutes(route *gin.Engine) {
	var (
		fieldService getFieldService.FieldService   	   = getFieldService.New()	
		fieldController getFieldController.FieldController = getFieldController.New(fieldService)
	)

	groupRoute := route.Group("/api/v1")
	groupRoute.GET("/fields", func(ctx *gin.Context) {
		ctx.JSON(200, fieldController.FindAll())
	})

	groupRoute.POST("/fields", func(ctx *gin.Context) {
		err := fieldController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Field input found"})
		}
	})
}