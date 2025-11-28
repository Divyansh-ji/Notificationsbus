package db

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVar() {
	err := godotenv.Load() // loads .env from project root
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}
}
