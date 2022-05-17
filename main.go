package main

import (
	"io"
	"os"
	//"time"

	"github.com/KadirbekSharau/rentykz-backend/configs/db"
	"github.com/KadirbekSharau/rentykz-backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
			AllowOrigins:  []string{"*"},
			AllowMethods:  []string{"*"},
			AllowHeaders:  []string{"*"},
			AllowWildcard: true,
		}),
		//middlewares.BasicAuth(),
	)

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")
	routes.InitAuthRoutes(db, server)
	routes.InitFieldRoutes(db, server)
	routes.InitCityRoutes(db, server)

	server.Run(":8080")
}
