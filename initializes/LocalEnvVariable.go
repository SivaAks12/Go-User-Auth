package initializes

import (
	"log"

	"github.com/joho/godotenv"
)

func LocalEnvVariable() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
