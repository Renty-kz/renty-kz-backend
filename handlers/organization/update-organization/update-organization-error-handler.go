package updateOrganizationHandler

import (
	"net/http"

	"github.com/KadirbekSharau/rentykz-backend/models"
	"github.com/KadirbekSharau/rentykz-backend/util"
	"github.com/gin-gonic/gin"
)


var ActivationOrganizationConfig = util.ErrorConfig{
	Options: []util.ErrorMetaConfig{
		{
			Tag:     "required",
			Field:   "ID",
			Message: "id is required on param",
		},
	},
}

/* Activate Organization Error Handler Function */
func ErrActivateOrganizationHandler(res *models.EntityOrganizations, ctx *gin.Context, errUpdateOrganization string) {
	switch errUpdateOrganization {
	case "RESULTS_ORGANIZATION_NOT_FOUND_404":
		util.APIResponse(ctx, "Organization not found", http.StatusNotFound, http.MethodPut, nil)
		return

	default:
		util.APIResponse(ctx, "Updated Organization successfully", http.StatusCreated, http.MethodPut, res)
	}
}