// routes/routes.go
package routes

import (
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(controllers.AuthMiddleware())
	{
		authorized.GET("/users", controllers.GetUsers)
		authorized.GET("/users/:id", controllers.GetUser)
		authorized.PUT("/users/:id", controllers.UpdateUser)
		authorized.DELETE("/users/:id", controllers.DeleteUser)
		authorized.GET("/fetch-countries", controllers.FetchCountries)
		authorized.GET("/countries", controllers.GetCountries)
	}

	return r
}
