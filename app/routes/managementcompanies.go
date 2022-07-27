package routes

import (
	"github.com/Jon1701/property-reviews/app/controllers"
	"github.com/Jon1701/property-reviews/app/middleware"
	"github.com/gin-gonic/gin"
)

// Registers the routes under the /api/management group.
func RegisterManagementCompaniesRoutes(appCtx *controllers.AppContext, rg *gin.RouterGroup) {
	rg.POST("/", appCtx.CreateManagementCompany)
	rg.GET("/", middleware.SanitizePaginationParameters(), appCtx.GetManagementCompanies)
	rg.GET("/:managementID", appCtx.GetManagementCompanyByID)
	rg.PATCH("/:managementID", appCtx.UpdateManagementCompany)
}
