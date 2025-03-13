package routes

import (
	"os"

	"github.com/OdairPianta/julia/http/middlewares"
	"github.com/OdairPianta/julia/http/requests"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	ginMode := os.Getenv("GIN_MODE")
	if ginMode != "" {
		gin.SetMode(ginMode)
	}
	r := gin.Default()
	InitRoutes(r)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

func InitRoutes(r *gin.Engine) *gin.Engine {
	requests.RegisterCustomValidators()

	r.Use(middlewares.JSONLogMiddleware(), middlewares.BlockMiddleware(), middlewares.CorsMiddleware())

	public := r.Group("/api")
	RegisteSwaggerRoutes(public)
	RegisterStatusRoutes(public)
	RegisterAuthRoutes(public)
	RegisterPoliciesRoutes(public)

	protected := r.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())
	RegisterFilesRoutes(protected)
	RegisterSampleRoutes(protected)
	RegisterUserRoutes(protected)
	return r
}
