package controllers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func error(w http.ResponseWriter, code string) {
	url := "./template/" + code + ".html"
	t, err := template.ParseFiles(url)
	if err != nil {
		fmt.Fprintf(w, "Internal Server error")
		//w.WriteHeader(500)
		log.Fatal(err)
	}
	p := ""
	t.Execute(w, p)
}
