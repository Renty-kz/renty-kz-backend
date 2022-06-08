package delete

import (
	deleteAuth "github.com/KadirbekSharau/rentykz-backend/controllers/auth-controllers/delete"
	"github.com/KadirbekSharau/rentykz-backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	DeleteHandler(ctx *gin.Context)
}

type handler struct {
	service deleteAuth.Service
}

func NewHandlerDelete(service deleteAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) DeleteHandler(ctx *gin.Context) {

	var input deleteAuth.InputDelete
	ctx.ShouldBindJSON(&input)
	resultDelete, errDelete := h.service.DeleteService(&input)

	switch errDelete {

	case "ACCOUNT_NOT_FOUND_404":
		util.APIResponse(ctx, "User account is not registered", http.StatusNotFound, http.MethodDelete, nil)
		return

	default:
		util.APIResponse(ctx, "Delete successfully", http.StatusOK, http.MethodDelete, map[string]uint{"Deleted user's id": resultDelete})
		return
	}
}
