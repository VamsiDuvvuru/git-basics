package middleware

import (
	"example/my-project-go/module/project2/rest/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Please pass the valid access token and try again!!"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Access token is expired or invalid , please try again!!"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
