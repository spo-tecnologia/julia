package routes

import (
	"github.com/OdairPianta/julia/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisteSwaggerRoutes(router *gin.RouterGroup) {
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
