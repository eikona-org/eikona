package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/imgProcessing/backend/v2/service"
	"net/http"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) <= 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.JWTAuthService().ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
