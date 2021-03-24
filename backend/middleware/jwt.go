package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/imgProcessing/backend/v2/service"
	"log"
	"net/http"
)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		log.Println(authHeader[len(BEARER_SCHEMA):])
		tokenString := authHeader[len(BEARER_SCHEMA):]
		log.Println(len(authHeader))

		log.Println("ok")
		token, err := service.JWTAuthService().ValidateToken(tokenString)
		log.Println("nok")
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[User]: ", claims["user"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])
		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
