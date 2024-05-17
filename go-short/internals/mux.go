package internals

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const prefix = "go-sh/"

var router = mux.NewRouter()
var Client, collection = SetupMongo()

func Start_Server() int {

	router.HandleFunc("/add/{url}", InsertURL_Server)
	router.HandleFunc("/goto/{url}", GetURL_Server)
	router.HandleFunc("/gets/{url}", GetURL_No_Redirect_Server)

	http.Handle("/", router)
	port := ":" + os.Getenv("SERVER_PORT")

	if http.ListenAndServe(port, router) != nil {
		return -1
	}
	return 0
}
func InsertURL_Server(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["url"])
	result := AddURL(collection, vars["url"])
	fmt.Fprintf(w, "URL: %v\n", result)
}

func GetURL_Server(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["url"])
	result, _ := GetLinkFromShort(collection, vars["url"])
	result = "https://" + result
	http.Redirect(w, r, result, http.StatusSeeOther)
	fmt.Println(result)
}

func GetURL_No_Redirect_Server(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["url"])
	result, _ := GetLinkFromShort(collection, vars["url"])
	result = "https://" + result
	fmt.Fprintf(w, "Original: %v\n", result)
	fmt.Println(result)
}
