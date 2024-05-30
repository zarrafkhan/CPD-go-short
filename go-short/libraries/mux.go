package libraries

import (
	"context"
	utils "example/go-short/libraries/util"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/julienschmidt/httprouter"
)

var (
	Client, Collection = SetupMongo()
	temp               *template.Template
)

type Row struct {
	ShortList []string
	IDList    []string
}

func init() {
	temp = template.Must(template.ParseGlob("temps/index.html"))
}

func printLocal() {
	u, err := url.Parse("http://localhost:8080/")
	if err != nil {
		fmt.Println("local fail")
		log.Fatal(err)
	}
	fmt.Printf("%+v \n", u)
}

func rootHandle(w http.ResponseWriter, r *http.Request) {

	var shorts, ogs []string

	for _, l := range GetShorts(Collection) {
		shorts = append(shorts, l.ShortLink)
		ogs = append(ogs, l.ID)
	}

	data := Row{
		ShortList: shorts,
		IDList:    ogs,
	}

	err := temp.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Start_Server() int {
	utils.LoadEnv()

	// setup router and serving css stylesheet
	router := httprouter.New()
	printLocal()
	router.ServeFiles("/styles/*filepath", http.Dir("styles"))

	//http handlers for redirecting
	router.HandlerFunc(http.MethodGet, "/", rootHandle)
	router.HandlerFunc(http.MethodPost, "/", HandleNewLinkSubmit)
	router.HandlerFunc(http.MethodGet, "/gs/:short", HandleRedirect)

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

	if !VerifyLink(urls) {
		fmt.Println("Please retry with a valid url")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	//added context w/ timeout to ensure it handles cancellations for IO blocking operations
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	full, sh := AddURL(Collection, urls, ctx)
	fmt.Println(full, " ", sh)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	url := params.ByName("short")
	result, _ := GetLinkFromShort(Collection, url)
	println(result)
	http.Redirect(w, r, result, http.StatusSeeOther)
}

func RemoveURL(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	urls := r.Form.Get("links")
	if urls == "" {
		log.Fatal("Did not parse")
		return
	}
	DeletURL(Collection, urls)
	fmt.Fprintf(w, "'%v' has been removed.\n", urls)
}

// func GetList() {
// 	for _, l := range GetShorts(Collection) {
// 		IDList = append(IDList, l.ID)
// 		ShortList = append(ShortList, l.ShortLink)
// 	}
// }

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
