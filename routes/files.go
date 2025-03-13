package routes

import (
	"github.com/OdairPianta/julia/http/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterFilesRoutes(router *gin.RouterGroup) {
	router.GET("/files", controllers.FindFiles)
	router.POST("/files", controllers.CreateFile)
	router.DELETE("/files/:id", controllers.DeleteFile)
}
