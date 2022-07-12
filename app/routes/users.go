package routes

import (
	"github.com/Jon1701/property-reviews/app/controllers"
	"github.com/gin-gonic/gin"
)

// Registers the routes under the /api/users group.
func RegisterUserRoutes(appCtx *controllers.AppContext, rg *gin.RouterGroup) {
	rg.POST("/signup", appCtx.CreateUser)
	rg.POST("/login", appCtx.UserLogin)
}
