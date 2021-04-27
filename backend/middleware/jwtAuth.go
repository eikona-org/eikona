package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/eikona-org/eikona/v2/helper"
	"github.com/eikona-org/eikona/v2/service"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := jwtService.ValidateToken(tokenString)
		if !token.Valid {
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
