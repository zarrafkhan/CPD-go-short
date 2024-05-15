package main

import (
	"example/go-short/internals/db"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// loads .env
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()
	e := db.Init(os.Getenv("MONGO_KEY"), "cpd")
	Check(e)

	fmt.Println("Mongo Bongo")

	//returns till main() returns smth
	defer func() {
		Check(db.Disc())
	}()
}

// handle error
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
