package getOrganizationsHandler

import (
	"net/http"

	"github.com/KadirbekSharau/rentykz-backend/models"
	"github.com/KadirbekSharau/rentykz-backend/util"
	"github.com/gin-gonic/gin"
)

/* All Organizations Error Handler Function */
func ErrAllOrganizationsHandler(organizations *[]models.EntityOrganizations, ctx *gin.Context, err string) {
	switch err {

	case "RESULTS_ORGANIZATION_NOT_FOUND_404":
		util.APIResponse(ctx, "All Organizations data is not exists", http.StatusNotFound, http.MethodGet, nil)

	default:
		util.APIResponse(ctx, "Results All Organizations data successfully", http.StatusOK, http.MethodGet, organizations)
	}
}

/* Active Organizations Error Handler Function */
func ErrActiveOrganizationsHandler(organizations *[]models.EntityOrganizations, ctx *gin.Context, err string) {
	switch err {

	case "RESULTS_ORGANIZATION_NOT_FOUND_404":
		util.APIResponse(ctx, "Active Organizations data is not exists", http.StatusNotFound, http.MethodGet, nil)

	default:
		util.APIResponse(ctx, "Results Active Organizations data successfully", http.StatusOK, http.MethodGet, organizations)
	}
}

/* Inactive Organizations Error Handler Function */
func ErrInactiveOrganizationsHandler(organizations *[]models.EntityOrganizations, ctx *gin.Context, err string) {
	switch err {

	case "RESULTS_ORGANIZATION_NOT_FOUND_404":
		util.APIResponse(ctx, "Inactive Organizations data is not exists", http.StatusNotFound, http.MethodGet, nil)

	default:
		util.APIResponse(ctx, "Results Inactive Organizations data successfully", http.StatusOK, http.MethodGet, organizations)
	}
}