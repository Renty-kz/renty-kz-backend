package main

import (
	"io"
	"net/http"
	"os"

	"github.com/KadirbekSharau/rentykz-backend/controller"
	"github.com/KadirbekSharau/rentykz-backend/middlewares"
	"github.com/KadirbekSharau/rentykz-backend/service"
	"github.com/gin-gonic/gin"
)

var (
	fieldService service.FieldService = service.New()
	fieldController controller.FieldController = controller.New(fieldService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	server.Use(
		gin.Recovery(), 
		middlewares.Logger(), 
		middlewares.BasicAuth(),
	)

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/fields", func(ctx *gin.Context) {
			ctx.JSON(200, fieldController.FindAll())
		})
	
		apiRoutes.POST("/fields", func(ctx *gin.Context) {
			err := fieldController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message" : "Field input found"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", fieldController.ShowAll)
	}


	server.Run(":8080")
}