package main

import (
	"log"
	"main/internal/config"
	"main/internal/handlers/user_handler"
	"main/internal/infrastructure/db"
	"main/internal/middleware"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitEnv()

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

	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	pool := conn.GetPool()

	//router.GET("/", middleware.AuthMiddleware(), handlers.IndexView)
	//router.GET("/t", http.HandlerFunc(handlers.IndexView))
	router.POST("/create_user", middleware.AuthMiddleware(), user_handler.CreateUser(pool))
	router.POST("/login", user_handler.LoginView(pool))

	err = router.Run(":8080")
	if err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
