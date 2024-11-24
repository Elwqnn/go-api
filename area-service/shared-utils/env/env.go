package env

import (
	"github.com/joho/godotenv"
	"log"
)

// LoadEnv loads the .env file
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}
}
