package routes

import (
	deleteAuthController "github.com/KadirbekSharau/rentykz-backend/controllers/auth-controllers/delete"
	loginAuthController "github.com/KadirbekSharau/rentykz-backend/controllers/auth-controllers/login"
	registerAuthController "github.com/KadirbekSharau/rentykz-backend/controllers/auth-controllers/register"
	deleteHandler "github.com/KadirbekSharau/rentykz-backend/handlers/auth/delete"
	LoginHandler "github.com/KadirbekSharau/rentykz-backend/handlers/auth/login"
	registerHandler "github.com/KadirbekSharau/rentykz-backend/handlers/auth/register"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Auth routes */
func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		loginRepository loginAuthController.Repository = loginAuthController.NewRepositoryLogin(db)
		loginService    loginAuthController.Service    = loginAuthController.NewServiceLogin(loginRepository)
		loginHandler    LoginHandler.Handler           = LoginHandler.NewLoginHandler(loginService)

		registerRepository registerAuthController.Repository = registerAuthController.NewRepositoryRegister(db)
		registerService    registerAuthController.Service    = registerAuthController.NewServiceRegister(registerRepository)
		registerHandler    registerHandler.Handler           = registerHandler.NewHandlerRegister(registerService)

		deleteRepository deleteAuthController.Repository = deleteAuthController.NewRepositoryDelete(db)
		deleteService    deleteAuthController.Service    = deleteAuthController.NewServiceDelete(deleteRepository)
		deleteHandler    deleteHandler.Handler           = deleteHandler.NewHandlerDelete(deleteService)
	)

	groupRoute := route.Group("/api/v1/auth")
	groupRoute.POST("/login", loginHandler.LoginHandler)
	groupRoute.POST("/register", registerHandler.ActiveUserRegisterHandler)
	groupRoute.DELETE("/delete", deleteHandler.DeleteHandler)
}
