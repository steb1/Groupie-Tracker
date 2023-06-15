package controllers

import (
	"net/http"
	"text/template"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		t, _ := template.ParseFiles("./template/404.html")
		p := ""
		t.Execute(w, p)
		return
	}

	t, _ := template.ParseFiles("./template/index.html")
	p := ""
	t.Execute(w, p)
}
