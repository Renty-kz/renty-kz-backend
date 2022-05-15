package main

import (
	"io"
	"os"
	"time"

	"github.com/KadirbekSharau/rentykz-backend/configs/db"
	getFieldController "github.com/KadirbekSharau/rentykz-backend/controllers/field-controllers/get-field"
	getFieldService "github.com/KadirbekSharau/rentykz-backend/handlers/field/get-field"
	"github.com/KadirbekSharau/rentykz-backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	fieldService getFieldService.FieldService   	   = getFieldService.New()
	fieldController getFieldController.FieldController = getFieldController.New(fieldService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	db := db.NewDatabaseConnection()
	server := gin.New()

	server.Use(
		gin.Recovery(),
		gin.Logger(),
		cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"Content-Length", "text/plain", "Authorization", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
		//middlewares.BasicAuth(),
	)

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")
	routes.InitAuthRoutes(db, server)
	routes.InitFieldRoutes(db, server)
	routes.InitCityRoutes(db, server)

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", fieldController.ShowAll)
	}

	server.Run(":8080")
}
