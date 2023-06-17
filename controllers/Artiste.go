package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"
)

func ArtisteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Artiste" {
		t, err := template.ParseFiles("./template/404.html")
		if err != nil {
			error(w, "404")
			log.Fatal(err)
		}
		p := ""
		t.Execute(w, p)
		return
	}

	t, err := template.ParseFiles("./template/artistes.html")

	if err != nil {
		error(w, "500")
		log.Fatal(err)
	}
	response, _ := http.Get("http://groupietrackers.herokuapp.com/api/artists")
	responseData, _ := io.ReadAll(response.Body)
	var responseObjectArtist ResponseArtist
	json.Unmarshal(responseData, &responseObjectArtist)

	//fmt.Println(responseObjectArtist)

	t.Execute(w, responseObjectArtist)
}
