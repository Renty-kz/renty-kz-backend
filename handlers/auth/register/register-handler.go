package registerHandler

import (
	"fmt"
	"net/http"

	"github.com/KadirbekSharau/rentykz-backend/controllers/auth-controllers/register"
	"github.com/KadirbekSharau/rentykz-backend/util"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	ActiveUserRegisterHandler(ctx *gin.Context)
	InactiveUserRegisterHandler(ctx *gin.Context)
	AdminRegisterHandler(ctx *gin.Context)
	OrganizationRegisterHandler(ctx *gin.Context)
	ModeratorRegisterHandler(ctx *gin.Context)
}

type handler struct {
	service registerAuthController.Service
}

func NewHandlerRegister(service registerAuthController.Service) *handler {
	return &handler{service: service}
}

/* Active User Register Handler */
func (h *handler) ActiveUserRegisterHandler(ctx *gin.Context) {

	var input registerAuthController.InputUserRegister
	ctx.ShouldBindJSON(&input)
	fmt.Printf(input.Password + "\n")
	errResponse, errCount := util.GoValidator(input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultRegister, errRegister := h.service.InactiveUserRegisterService(&input)
	fmt.Printf(resultRegister.Password + "\n")
	ErrUserRegisterHandler(resultRegister, ctx, errRegister)
}

/* Inactive User Register Handler */
func (h *handler) InactiveUserRegisterHandler(ctx *gin.Context) {

	var input registerAuthController.InputUserRegister
	ctx.ShouldBindJSON(&input)

	errResponse, errCount := util.GoValidator(input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultRegister, errRegister := h.service.AdminRegisterService(&input)

	ErrUserRegisterHandler(resultRegister, ctx, errRegister)
}

/* Admin Register Handler */
func (h *handler) AdminRegisterHandler(ctx *gin.Context) {

	var input registerAuthController.InputUserRegister
	ctx.ShouldBindJSON(&input)

	errResponse, errCount := util.GoValidator(input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultRegister, errRegister := h.service.ActiveUserRegisterService(&input)

	ErrUserRegisterHandler(resultRegister, ctx, errRegister)
}

/* Organization Register Handler */
func (h *handler) OrganizationRegisterHandler(ctx *gin.Context) {

	var input registerAuthController.InputOrganizationRegister
	ctx.ShouldBindJSON(&input)

	errResponse, errCount := util.GoValidator(input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultRegister, errRegister := h.service.OrganizationRegisterService(&input)

	ErrOrganizationRegisterHandler(resultRegister, ctx, errRegister)
}

/* Moderator Register Handler */
func (h *handler) ModeratorRegisterHandler(ctx *gin.Context) {

	var input registerAuthController.InputModeratorRegister
	ctx.ShouldBindJSON(&input)

	errResponse, errCount := util.GoValidator(input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultRegister, errRegister := h.service.ModeratorRegisterService(&input)

	ErrModeratorRegisterHandler(resultRegister, ctx, errRegister)
}