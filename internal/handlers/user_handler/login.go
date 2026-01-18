package user_handler

import (
	"main/internal/dto"
	"main/internal/servicves/user_services"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func LoginView(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials dto.LoginRequest

		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx := c.Request.Context()
		user, err := user_services.LoginUser(ctx, pool, credentials)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := generateJWT(user.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not authorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func generateJWT(email string) (string, error) {
	var jwtSecretKey = []byte(os.Getenv("SECRET_KEY"))

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenObj.SignedString(jwtSecretKey)
	return token, err
}
