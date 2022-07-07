package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default();

	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	serverPort := os.Getenv("SERVER_PORT")

	err := r.Run(fmt.Sprintf(":%s", serverPort))
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error %v", err))
	}
}