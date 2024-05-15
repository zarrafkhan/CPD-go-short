package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	//"github.com/gin-gonic/gin"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

// func testShorten() {
// 	l1 := models.Shorten("https://example.com/")
// 	fmt.Println(l1.SH)
// }

func main() {
	// printLocal()

	// http.HandleFunc("/", rootHandle)
	// //serve css
	// fileServer := http.FileServer(http.Dir("styles"))
	// http.Handle("/styles/", http.StripPrefix("/styles", fileServer))
	// http.ListenAndServe(":5050", nil)

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://testUser:OCMBJtQstHm9dhYr@cluster-go-short.c4w8kgt.mongodb.net/?retryWrites=true&w=majority&appName=Cluster-go-short").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

}

// handle error
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
