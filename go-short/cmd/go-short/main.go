package main

import (
	"example/go-short/internals/db"
	// "example/go-short/internals/models"

	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
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

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("temps/index.html"))
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}

func printLocal() {
	u, err := url.Parse("http://localhost:5050/")
	Check(err)
	fmt.Printf("%+v \n", u)
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

	printLocal()

	http.HandleFunc("/", rootHandle)
	//serve css
	fileServer := http.FileServer(http.Dir("styles"))
	http.Handle("/styles/", http.StripPrefix("/styles", fileServer))
	http.ListenAndServe(":5050", nil)

}

// handle error
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
