package router

import (
	"main/internal/handlers/user_handler"
	"main/internal/middleware"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(pool *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"), ",")
	router.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	router.SetTrustedProxies([]string{})

	router.POST("/create_user", middleware.AuthMiddleware(), user_handler.CreateUser(pool))
	router.POST("/delete_user", middleware.AuthMiddleware(), user_handler.DeleteUser(pool))
	router.POST("/login", user_handler.LoginView(pool))

	return router
}
