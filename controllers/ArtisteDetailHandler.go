package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"text/template"
)

func ArtisteDetailHandler(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query()["ID"]
	IDD := ID[0]
	intID, err := strconv.Atoi(IDD)

	if err != nil {
		error(w)
		return
	}

	if IDD == "0" || len(ID) == 0 || intID > 57 || IDD == "" {
		t, err := template.ParseFiles("./template/404.html")
		if err != nil {
			error(w)
			return
		}
		p := ""
		t.Execute(w, p)
		return
	}

	intID -= 1

	response, _ := http.Get("http://groupietrackers.herokuapp.com/api/artists")
	responseData, _ := io.ReadAll(response.Body)
	var responseObjectArtist ResponseArtist
	json.Unmarshal(responseData, &responseObjectArtist)

	response1, _ := http.Get("http://groupietrackers.herokuapp.com/api/locations")
	responseData1, _ := io.ReadAll(response1.Body)
	var responseObjectLocations ResponseLocations
	json.Unmarshal(responseData1, &responseObjectLocations)

	data := &data{
		Artistes:  responseObjectArtist,
		Locations: responseObjectLocations,
		ID:        intID,
	}

	fmt.Println(data.Locations)

	t, err := template.ParseFiles("./template/artisteDetail.html")

	if err != nil {
		error(w)
		return
	}

	t.Execute(w, data)
}
