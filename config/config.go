package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(value string) string {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	return os.Getenv(value)
}
