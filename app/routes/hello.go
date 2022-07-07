package routes

import (
	"github.com/Jon1701/property-reviews/app/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterRoutes = func (r *gin.Engine) {
	r.GET("/", controllers.RootHandlerFunc)
}