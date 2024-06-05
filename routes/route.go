// routes/routes.go
package routes

import (
	"go-rest-api/controllers"
	"go-rest-api/database"
	"go-rest-api/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *database.GormDatabase) *gin.Engine {
	r := gin.Default()

	userService := services.NewUserService(db)
	userController := controllers.NewUserController(userService)

	r.POST("/signup", userController.SignUp)
	r.POST("/login", userController.Login)

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(controllers.AuthMiddleware())
	{
		authorized.GET("/users", userController.GetUsers)
		authorized.GET("/users/:id", userController.GetUser)
		authorized.PUT("/users/:id", userController.UpdateUser)
		authorized.DELETE("/users/:id", userController.DeleteUser)
		authorized.GET("/fetch-countries", controllers.FetchCountries)
		authorized.GET("/countries", userController.GetCountries)
	}

	return r
}
