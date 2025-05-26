package middleware

import (
	"net/http"
	"strings"

	"github.com/Auxesia23/url_shortener/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JwtAuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			return 
		}
		
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
			return
		}

		token, err := auth.VerifyJWT(tokenString) 
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token otorisasi tidak valid atau kedaluwarsa"})
			return
		}
		
		claims := token.Claims.(jwt.MapClaims)
		
		user := claims["email"].(string)
		
		c.Set("user", user)
		c.Next()
	}
}