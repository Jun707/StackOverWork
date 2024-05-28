package database

import (
	"github.com/joho/godotenv"
	"log"
    "os"
)

// database connection options
func NewDBOptions() string {
	err := godotenv.Load()
	if err != nil {
        log.Fatal("Error loading .env file")
    }
	password := os.Getenv("password")
    endpoint := os.Getenv("endpoint")
    username := os.Getenv("username")
	dsn := username + ":" + password + "@tcp(" + endpoint + ":3306)/Job_Database?charset=utf8mb4"
	return dsn
}
