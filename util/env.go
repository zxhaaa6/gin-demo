package util

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	value := os.Getenv(key)
	return value
}

func LoadEnv() {
	filename := ".env"
	err := godotenv.Load(filename)
	if err != nil {
		log.Println(".env file not loaded")
	}
}
