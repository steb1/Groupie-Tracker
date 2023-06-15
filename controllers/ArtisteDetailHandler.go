package controllers

import (
	"net/http"
	"strconv"
	"text/template"
)

func ArtisteDetailHandler(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query()["ID"]

	if len(ID[0]) == 0 {
		error(w)
		return
	}
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

	responseObjectArtist := getArtistes("http://groupietrackers.herokuapp.com/api/artists")
	responseObjectLocations := getLocations("http://groupietrackers.herokuapp.com/api/locations")
	responseObjectRelations := getRelations("https://groupietrackers.herokuapp.com/api/dates")
	responseObjectDates := getDates("https://groupietrackers.herokuapp.com/api/relation")

	data := &data{
		Artistes:  responseObjectArtist,
		Locations: responseObjectLocations,
		Relation:  responseObjectRelations,
		Dates:     responseObjectDates,
		ID:        intID,
	}

	t, err := template.ParseFiles("./template/artisteDetail.html")

	if err != nil {
		error(w)
		return
	}

	t.Execute(w, data)
}
