package loginAuthController

import (
	"github.com/KadirbekSharau/rentykz-backend/dto"
	"github.com/KadirbekSharau/rentykz-backend/service/auth/jwt_service"
	"github.com/KadirbekSharau/rentykz-backend/service/auth/login"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService loginAuthService.LoginService
	jWtService   jwt_service.JWTService
}

func NewLoginController(loginService loginAuthService.LoginService,
	jWtService jwt_service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := controller.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return controller.jWtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
