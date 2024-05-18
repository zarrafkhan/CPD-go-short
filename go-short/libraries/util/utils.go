package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
