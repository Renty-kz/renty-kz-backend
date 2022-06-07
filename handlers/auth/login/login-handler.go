package loginHandler

import (
	"net/http"

	util "github.com/KadirbekSharau/rentykz-backend/util"
	"github.com/gin-gonic/gin"
	loginAuth "github.com/KadirbekSharau/rentykz-backend/controllers/auth/login"

)

type Handler interface {
	UserLoginHandler(ctx *gin.Context)
	OrganizationLoginHandler(ctx *gin.Context)
	ModeratorLoginHandler(ctx *gin.Context)
	AdminLoginHandler(ctx *gin.Context)
}

type handler struct {
	service loginAuth.Service
}

func NewLoginHandler(service loginAuth.Service) *handler {
	return &handler{service: service}
}

/* User Login Handler */
func (h *handler) UserLoginHandler(ctx *gin.Context) {

	var input loginAuth.InputLogin
	ctx.ShouldBindJSON(&input)

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultLogin, errLogin := h.service.UserLoginService(&input)

	UserLoginTokenHandler(ctx, errLogin, resultLogin)
}

/* Admin Login Handler */
func (h *handler) AdminLoginHandler(ctx *gin.Context) {

	var input loginAuth.InputLogin
	ctx.ShouldBindJSON(&input)

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultLogin, errLogin := h.service.AdminLoginService(&input)

	AdminLoginTokenHandler(ctx, errLogin, resultLogin)
}

/* Organization Login Handler */
func (h *handler) OrganizationLoginHandler(ctx *gin.Context) {

	var input loginAuth.InputLogin
	ctx.ShouldBindJSON(&input)

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultLogin, errLogin := h.service.OrganizationLoginService(&input)

	OrganizationLoginTokenHandler(ctx, errLogin, resultLogin)
}

/* Moderator Login Handler */
func (h *handler) ModeratorLoginHandler(ctx *gin.Context) {

	var input loginAuth.InputLogin
	ctx.ShouldBindJSON(&input)

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultLogin, errLogin := h.service.ModeratorLoginService(&input)

	ModeratorLoginTokenHandler(ctx, errLogin, resultLogin)
}