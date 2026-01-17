package main

import (
	"log"
	"main/db"
	"net/http"
	"os"
	"strings"
	"time"

	"main/config"
	//"main/middleware"
	"main/views"

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

	//router.GET("/", middleware.AuthMiddleware(), views.IndexView)
	//router.GET("/t", http.HandlerFunc(views.IndexView))
	//router.POST("/login", views.LoginView(pool))
	router.POST("/login", views.LoginView(pool))

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
