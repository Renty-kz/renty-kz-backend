package main

import (
	"io"
	"os"
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
	server := gin.Default()

	server.Use(
		cors.New(cors.Config{
			AllowOrigins:  []string{"*"},
			AllowMethods:  []string{"*"},
			AllowHeaders:  []string{"*"},
			AllowWildcard: true,
		}),
	)


	routes.InitAuthRoutes(db, server)
	routes.InitFieldRoutes(db, server)
	routes.InitCityRoutes(db, server)
	routes.InitOrganizationRoutes(db, server)

	server.Run(":8080")
}
