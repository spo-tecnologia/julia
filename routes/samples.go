package routes

import (
	"github.com/OdairPianta/julia/http/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterSampleRoutes(router *gin.RouterGroup) {
	router.GET("/sample_models", controllers.FindSamples)
	router.GET("/sample_models/:id", controllers.FindSample)
	router.POST("/sample_models", controllers.CreateSampleModel)
	router.PUT("/sample_models/:id", controllers.UpdateSampleModel)
	router.DELETE("/sample_models/:id", controllers.DeleteSampleModel)
	router.GET("/sample_models_select", controllers.FindSamples)
}
