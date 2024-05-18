package internals

import (
	utils "example/go-short/internals/util"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var route = mux.NewRouter().StrictSlash(true).UseEncodedPath()
var Client, Collection = SetupMongo()
var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("temps/index.html"))
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}

func printLocal() {
	u, err := url.Parse("http://localhost:8080/")
	utils.Check(err)
	fmt.Printf("%+v \n", u)
}

func Start_Server() int {

	// route.HandleFunc("/go-sh/{url}", GetURL_Server)
	// route.HandleFunc("/add/{url}", InsertURL_Server)
	// route.HandleFunc("/get/{url}", GetURL_No_Redirect_Server)
	// route.HandleFunc("/rm/{url}", RemoveURL)

	printLocal()
	utils.LoadEnv()

	//serve css stylesheet
	fileServer := http.FileServer(http.Dir("styles"))
	http.Handle("../styles/", http.StripPrefix("../styles", fileServer))

	route.HandleFunc("/", rootHandle)

	srv := &http.Server{
		Handler: route,
		Addr:    "localhost:" + os.Getenv("SERVER_PORT"),

		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// serving error handle
	if srv.ListenAndServe() != nil {
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

func GrabFormPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	r.PostForm.Get(vars["url"])

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
