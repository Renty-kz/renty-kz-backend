package routes

import (
	loginAuthController "github.com/KadirbekSharau/rentykz-backend/controllers/auth-controllers/login"
	LoginHandler "github.com/KadirbekSharau/rentykz-backend/handlers/auth/login"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Auth routes */
func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		LoginRepository loginAuthController.Repository = loginAuthController.NewRepositoryLogin(db)
		loginService loginAuthController.Service = loginAuthController.NewServiceLogin(LoginRepository)
		loginHandler LoginHandler.LoginHandler = LoginHandler.NewLoginHandler(loginService)
	)

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/login", loginHandler.LoginHandler)
}