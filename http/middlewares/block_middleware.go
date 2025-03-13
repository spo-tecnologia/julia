package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BlockMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.URL.Path[0:5] == "/.git" ||
			context.Request.URL.Path[0:5] == "/.env" {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão para acessar este recurso"})
			context.Abort()
			return
		}
		context.Next()
	}
}
