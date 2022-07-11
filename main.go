package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Jon1701/property-reviews/app/controllers"
	"github.com/Jon1701/property-reviews/app/routes"
	"github.com/Jon1701/property-reviews/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := storage.NewConnection()
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}

	appCtx := controllers.New(db)

	r := gin.Default()
	routes.RegisterRoutes(r)

	api := r.Group("/api")
	users := api.Group("/users")
	routes.RegisterUserRoutes(&appCtx, users)

	serverPort := os.Getenv("SERVER_PORT")

	err = r.Run(fmt.Sprintf(":%s", serverPort))
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error %v", err))
	}
}
