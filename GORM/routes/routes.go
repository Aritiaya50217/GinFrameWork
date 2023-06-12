package routes

import (
	"example/gorm-api/controller"

	"github.com/gin-gonic/gin"
)

// set up  router
func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/user-api")
	{
		api.GET("user", controller.GetUsers)
		api.POST("user", controller.CreateUser)
		api.GET("user/:id", controller.GetUserById)
		api.PUT("user/:id", controller.UpdateUser)
		api.DELETE("user/:id", controller.DeleteUser)
	}
	return r
}
