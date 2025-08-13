package middlewares

import (
	"net/http"
	"the-redeemed/event-service/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	claims, err := utils.ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	context.Set("email", email)
	context.Set("userId", userId)
	context.Next()
}
