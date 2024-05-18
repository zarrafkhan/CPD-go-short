package libraries

import (
	utils "example/go-short/libraries/util"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

// var route = mux.NewRouter()
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
	if err != nil {
		fmt.Println("local fail")
		log.Fatal(err)
	}
	fmt.Printf("%+v \n", u)
}

func Start_Server() int {
	router := httprouter.New()

	printLocal()
	utils.LoadEnv()

	//serve css stylesheet
	fileServer := http.FileServer(http.Dir("styles"))
	router.Handler(http.MethodGet, "/styles/", http.StripPrefix("/styles", fileServer))

	router.HandlerFunc(http.MethodGet, "/", rootHandle)
	router.HandlerFunc(http.MethodPost, "/", HandleNewLinkSubmit)
	router.HandlerFunc(http.MethodGet, "/gs/:short", HandleRedirect)

	// serving error handle
	if http.ListenAndServe(":8080", router) != nil {
		return -1
	}
	return 0
}

func HandleNewLinkSubmit(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	urls := r.Form.Get("links")
	if urls == "" {
		log.Fatal("Did not parse")
		return
	}

	full, sh := AddURL(Collection, urls)
	fmt.Println(full, " ", sh)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	url := params.ByName("short")
	result, _ := GetLinkFromShort(Collection, url)
	http.Redirect(w, r, result, http.StatusSeeOther)
	// println(url)
}

// func InsertURL_Server(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Println(vars["url"])
// 	result := AddURL(Collection, vars["url"])
// 	fmt.Fprintf(w, "Short: %v\n", result)
// }

// func GetURL_Server(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Println(vars["url"])
// 	result, _ := GetLinkFromShort(Collection, vars["url"])
// 	result = "https://" + result
// 	http.Redirect(w, r, result, http.StatusSeeOther)
// 	fmt.Println(result)
// }

// func GetURL_No_Redirect_Server(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Println(vars["url"])
// 	result, _ := GetLinkFromShort(Collection, vars["url"])
// 	result = "https://" + result
// 	fmt.Fprintf(w, "Full Link: %v\n", result)
// 	fmt.Println(result)
// }

// func RemoveURL(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Println(vars["url"])
// 	DeletURL(Collection, vars["url"])
// 	fmt.Fprintf(w, "'%v' has been removed.\n", vars["url"])
// }
