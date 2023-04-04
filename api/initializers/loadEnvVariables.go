package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEndVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
