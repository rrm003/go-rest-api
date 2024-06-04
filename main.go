// main.go
package main

import (
	"go-rest-api/database"
	"go-rest-api/routes"
	"log"

	_ "go-rest-api/docs"

	"github.com/joho/godotenv"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go REST API
// @version 1.0
// @description This is a sample REST API with JWT authentication.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @Schemes http
func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.InitDatabase()
	r := routes.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
