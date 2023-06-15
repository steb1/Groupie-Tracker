package controllers

import "net/http"

func ServeCSS(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path
	http.ServeFile(w, r, "template"+filename)
}
