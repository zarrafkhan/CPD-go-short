package main

import (
	"html/template"
	"log"
	"net/http"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("temps/index.html"))
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	http.HandleFunc("/", rootHandle)
	http.ListenAndServe(":5050", nil)
}

// handle error
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
