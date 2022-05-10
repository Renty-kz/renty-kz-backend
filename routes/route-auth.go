package routes

import (
	"net/http"

	loginAuthController "github.com/KadirbekSharau/rentykz-backend/controller/auth-controllers/login"
	"github.com/KadirbekSharau/rentykz-backend/service/auth/jwt_service"
	loginAuthService "github.com/KadirbekSharau/rentykz-backend/service/auth/login"
	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"
)

/* @description All Auth routes */
func InitAuthRoutes(route *gin.Engine) {
	var (
		loginService loginAuthService.LoginService = loginAuthService.NewLoginService()
		jwtService   jwt_service.JWTService		   = jwt_service.NewJWTService()
		loginController loginAuthController.LoginController = loginAuthController.NewLoginController(loginService, jwtService)
	)

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
}