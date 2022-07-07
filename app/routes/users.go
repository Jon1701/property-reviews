package routes

import (
	"github.com/Jon1701/property-reviews/app/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	rg.POST("/signup", controllers.CreateUser)
}