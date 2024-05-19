package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// simple err log
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// loads .env file in root dir
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
