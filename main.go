package main

import (
	"net/http"

	"main/config"
	"main/middleware"
	"main/views"
)

func main() {
	config.InitEnv()

	mux := http.NewServeMux()

	mux.Handle("GET /", middleware.AuthMiddleware(http.HandlerFunc(views.IndexView)))
	mux.Handle("GET /t", http.HandlerFunc(views.IndexView))
	mux.Handle("POST /login", http.HandlerFunc(views.LoginView))

	http.ListenAndServe(":8080", mux)
}
