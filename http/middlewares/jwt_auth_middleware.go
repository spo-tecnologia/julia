package middlewares

import (
	"net/http"
	"strconv"

	"github.com/OdairPianta/julia/repository"
	"github.com/OdairPianta/julia/services/token"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := token.TokenValid(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não está logado"})
			context.Abort()
			return
		}
		var userId, errToken = token.ExtractTokenID(context)
		if errToken != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não está logado"})
			context.Abort()
			return
		}
		var stringId = strconv.FormatUint(uint64(userId), 10)
		var user, errUser = repository.FindUserByID(stringId)
		if errUser != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não está logado"})
			context.Abort()
			return
		}
		context.Set("user", *user)
		context.Next()
	}
}
