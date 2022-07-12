package routes

import (
	"github.com/Jon1701/property-reviews/app/controllers"
	"github.com/gin-gonic/gin"
)

// Registers the routes under the /api/management group.
func RegisterManagementCompaniesRoutes(appCtx *controllers.AppContext, rg *gin.RouterGroup) {
	rg.POST("/", appCtx.CreateManagementCompany)
	rg.PATCH("/:managementID", appCtx.UpdateManagementCompany)
}
