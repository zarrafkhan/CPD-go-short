package internals

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()
var Client, Collection = SetupMongo()

func Start_Server() int {

	router.HandleFunc("/go-sh/{url}", InsertURL_Server)
	router.HandleFunc("/goto/{url}", GetURL_Server)
	router.HandleFunc("/gets/{url}", GetURL_No_Redirect_Server)
	router.HandleFunc("/rm/{url}", RemoveURL)

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
	result := AddURL(Collection, vars["url"])
	fmt.Fprintf(w, "Short: %v\n", result)
}

func GetURL_Server(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["url"])
	result, _ := GetLinkFromShort(Collection, vars["url"])
	result = "https://" + result
	http.Redirect(w, r, result, http.StatusSeeOther)
	fmt.Println(result)
}

func GetURL_No_Redirect_Server(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["url"])
	result, _ := GetLinkFromShort(Collection, vars["url"])
	result = "https://" + result
	fmt.Fprintf(w, "Full Link: %v\n", result)
	fmt.Println(result)
}

func RemoveURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["url"])
	DeletURL(Collection, vars["url"])
	fmt.Fprintf(w, "'%v' has been removed.\n", vars["url"])
}
