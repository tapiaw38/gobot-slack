package env

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv loads the .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
