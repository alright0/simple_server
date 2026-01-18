package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

type DbConnection struct {
	user             string
	password         string
	host             string
	port             string
	name             string
	ConnectionString string
}

func GetDbConnectionString() *DbConnection {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, name)

	return &DbConnection{
		user:             user,
		password:         password,
		host:             host,
		port:             port,
		name:             name,
		ConnectionString: connectionString,
	}
}
