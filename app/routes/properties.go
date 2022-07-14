package routes

import (
	"github.com/Jon1701/property-reviews/app/controllers"
	"github.com/gin-gonic/gin"
)

// Registers the routes under the /api/property group.
func RegisterPropertiesRoutes(appCtx *controllers.AppContext, rg *gin.RouterGroup) {
	rg.POST("/", appCtx.CreateProperty)
	rg.PATCH("/:propertyID", appCtx.UpdateProperty)
}
