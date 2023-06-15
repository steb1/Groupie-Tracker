package controllers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func error(w http.ResponseWriter) {
	t, err := template.ParseFiles("./template/404.html")
	if err != nil {
		fmt.Fprintf(w, "Internal Server error")
		w.WriteHeader(500)
		log.Fatal(err)
	}
	p := ""
	t.Execute(w, p)
}
