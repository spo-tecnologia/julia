package routes

import (
	"github.com/OdairPianta/julia/http/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", controllers.Login)
	router.POST("/forgot_password", controllers.ForgotPassword)
	router.POST("/recover_password", controllers.RecoverPassword)
}
