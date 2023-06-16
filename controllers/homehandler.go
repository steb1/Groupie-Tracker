package controllers

import (
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		t, err := template.ParseFiles("./template/404.html")
		if err != nil {
			w.WriteHeader(500)
			error(w)
			return
		}
		p := ""
		t.Execute(w, p)
		return
	}

	t, _ := template.ParseFiles("./template/index.html")
	p := ""
	t.Execute(w, p)
}
