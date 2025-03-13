package routes

import (
	"github.com/OdairPianta/julia/http/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterStatusRoutes(router *gin.RouterGroup) {
	router.GET("/status", controllers.FindStatus)
}
