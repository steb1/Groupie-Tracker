package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"text/template"
)

func ArtisteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Artiste" {
		t, err := template.ParseFiles("./template/404.html")
		if err != nil {
			error(w, "500")
			return
		}
		p := ""
		t.Execute(w, p)
		return
	}

	t, err := template.ParseFiles("./template/artistes.html")

	if err != nil {
		error(w, "500")
		return
	}
	response, err := http.Get("http://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		error(w, "500")
		return
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		error(w, "500")
		return
	}
	var responseObjectArtist ResponseArtist
	json.Unmarshal(responseData, &responseObjectArtist)

	//fmt.Println(responseObjectArtist)

	t.Execute(w, responseObjectArtist)
}
