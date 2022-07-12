package main

import (
	"fmt"
	"os"

	"github.com/Jon1701/property-reviews/app/controllers"
	"github.com/Jon1701/property-reviews/app/routes"
	"github.com/Jon1701/property-reviews/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	// Check for Postgres Application Connection string.
	dbConnString := os.Getenv("POSTGRES_APP_CONNSTRING")
	if len(dbConnString) == 0 {
		panic("Missing Environment Variable: POSTGRES_APP_CONNSTRING")
	}

	// Check for JWT Signing Key.
	jwtSigningKey := os.Getenv("JWT_SIGNING_KEY")
	if len(jwtSigningKey) == 0 {
		panic("Missing Environment Variable: JWT_SIGNING_KEY")
	}

	// Check for Server Port.
	serverPort := os.Getenv("SERVER_PORT")
	if len(serverPort) == 0 {
		panic("Missing Environment Variable: SERVER_PORT")
	}

	// Create database connection.
	db, err := storage.NewConnection(dbConnString)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v\n", err))
	}

	// Define application context.
	appCtx := controllers.New(db)

	// Router.
	r := gin.Default()

	// Subrouters.
	api := r.Group("/api")
	users := api.Group("/users")
	management := api.Group("/management")

	// Register /api/users/*.
	routes.RegisterUserRoutes(&appCtx, users)

	// Register /api/management/*.
	routes.RegisterManagementCompaniesRoutes(&appCtx, management)

	// Run server.
	err = r.Run(fmt.Sprintf(":%s", serverPort))
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server: %v\n", err))
	}
}
