package main

import (
	"fmt"
	"os"

	"github.com/Jon1701/property-reviews/app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default();
	
	routes.RegisterRoutes(r)

	serverPort := os.Getenv("SERVER_PORT")

	err := r.Run(fmt.Sprintf(":%s", serverPort))
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error %v", err))
	}
}