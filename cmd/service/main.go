package main

import (
	"log"
	"main/internal/config"
	"main/internal/infrastructure/db"
	"main/internal/router"
)

func main() {
	config.InitEnv()

	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	pool := conn.GetPool()

	r := router.SetupRouter(pool)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
