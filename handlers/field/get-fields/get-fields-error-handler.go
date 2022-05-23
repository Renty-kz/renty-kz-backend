package getFieldsHandler

import (
	"net/http"

	"github.com/KadirbekSharau/rentykz-backend/models"
	"github.com/KadirbekSharau/rentykz-backend/util"
	"github.com/gin-gonic/gin"
)

var configCityId = util.ErrorConfig{
	Options: []util.ErrorMetaConfig{
		{
			Tag:     "required",
			Field:   "ID",
			Message: "id is required on param",
		},
	},
}

var configSportTypeId = util.ErrorConfig{
	Options: []util.ErrorMetaConfig{
		{
			Tag:     "required",
			Field:   "ID",
			Message: "id is required on param",
		},
	},
}

/* Fields By City Id Error Handler Function */
func ErrFieldsByCityIdHandler(fields *[]models.EntityFields, ctx *gin.Context, err string) {
	switch err {

	case "RESULTS_FIELD_NOT_FOUND_404":
		util.APIResponse(ctx, "Fields in this city do not exist", http.StatusConflict, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Results Fields data successfully", http.StatusOK, http.MethodPost, fields)
	}
}

/* Fields By Sport Type Id Error Handler Function */
func ErrFieldsBySportTypeIdHandler(fields *[]models.EntityFields, ctx *gin.Context, err string) {
	switch err {

	case "RESULTS_FIELD_NOT_FOUND_404":
		util.APIResponse(ctx, "Fields in this sport type do not exist", http.StatusConflict, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Results Fields data successfully", http.StatusOK, http.MethodPost, fields)
	}
}

/* Fields By Organization Id Error Handler Function */
func ErrFieldsByOrganizationIdHandler(fields *[]models.EntityFields, ctx *gin.Context, err string) {
	switch err {

	case "RESULTS_FIELD_NOT_FOUND_404":
		util.APIResponse(ctx, "Fields in this organization do not exist", http.StatusConflict, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Results Fields data successfully", http.StatusOK, http.MethodPost, fields)
	}
}

/* Fields By Moderator Id Error Handler Function */
func ErrFieldsByModeratorIdHandler(fields *[]models.EntityFields, ctx *gin.Context, err string) {
	switch err {

	case "RESULTS_FIELD_NOT_FOUND_404":
		util.APIResponse(ctx, "Fields in this moderator do not exist", http.StatusConflict, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Results Fields data successfully", http.StatusOK, http.MethodPost, fields)
	}
}