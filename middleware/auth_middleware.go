package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized"})
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]
		var jwtSecretKey = []byte(os.Getenv("SECRET_KEY"))
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not valid"})
			}
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Error decoding token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		email := claims["email"].(string)
		c.Set("email", string(email))

		c.Next()
	}
}
