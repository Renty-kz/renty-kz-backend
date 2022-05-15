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
}

type handler struct {
	service registerAuthController.Service
}

func NewHandlerRegister(service registerAuthController.Service) *handler {
	return &handler{service: service}
}

/* Active User Register Handler */
func (h *handler) ActiveUserRegisterHandler(ctx *gin.Context) {

	var input registerAuthController.InputRegister
	ctx.ShouldBindJSON(&input)
	fmt.Printf(input.Password + "\n")
	errResponse, errCount := util.GoValidator(input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultRegister, errRegister := h.service.InactiveUserRegisterService(&input)
	fmt.Printf(resultRegister.Password + "\n")
	ErrRegisterHandler(resultRegister, ctx, errRegister)
}

/* Inactive User Register Handler */
func (h *handler) InactiveUserRegisterHandler(ctx *gin.Context) {

	var input registerAuthController.InputRegister
	ctx.ShouldBindJSON(&input)

	errResponse, errCount := util.GoValidator(input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultRegister, errRegister := h.service.AdminRegisterService(&input)

	ErrRegisterHandler(resultRegister, ctx, errRegister)
}

/* Admin Register Handler */
func (h *handler) AdminRegisterHandler(ctx *gin.Context) {

	var input registerAuthController.InputRegister
	ctx.ShouldBindJSON(&input)

	errResponse, errCount := util.GoValidator(input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultRegister, errRegister := h.service.ActiveUserRegisterService(&input)

	ErrRegisterHandler(resultRegister, ctx, errRegister)
}