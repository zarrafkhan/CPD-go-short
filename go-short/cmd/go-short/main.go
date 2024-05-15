package main

import (
	"example/go-short/internals/db"
	"example/go-short/models"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.GET("/", getURL)
	r.Run("localhost:5050")
	fmt.Println("https://localhost:5050")

}

func getURL(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Shorten("https://example.com/"))
}

// handle error
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
