package routes

import (
	"github.com/OdairPianta/julia/http/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	router.GET("/users", controllers.FindUsers)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id", controllers.FindUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.PUT("/users/:id/update_fcm_token", controllers.UpdateFcmToken)
	router.DELETE("/users/:id", controllers.DeleteUser)
	router.GET("/users_select", controllers.SelectUsers)
}
