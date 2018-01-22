package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template



func main() {
	
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("starting-files/public"))))
	http.HandleFunc("/", dogs)
	http.ListenAndServe(":8080", nil)
}

func dogs(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))
	e := tpl.Execute(w, nil)
	if e != nil {
		log.Fatalln("template didn't execute: ", e)
	}
}
