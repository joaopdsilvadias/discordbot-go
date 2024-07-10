package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(env string) string {
	err := godotenv.Load(`.env`)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(env)
}
