package controllers

import (
	"net/http"
	"text/template"
)

func ServeCSS(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path
	_, err := template.ParseFiles("./template" + filename)

	if err != nil {
		//w.WriteHeader(404)
		return
	}
	http.ServeFile(w, r, "template"+filename)
}
