package controllers

import (
	"net/http"
	"strconv"
	"text/template"
)

func ArtisteDetailHandler(w http.ResponseWriter, r *http.Request) {

	// Recup the ID passed through the URL

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

	// Get request to the API
	responseObjectArtist := getArtistes("http://groupietrackers.herokuapp.com/api/artists")
	responseObjectLocations := getLocations("http://groupietrackers.herokuapp.com/api/locations")
	responseObjectDates := getDates("http://groupietrackers.herokuapp.com/api/dates")
	responseObjectRelations := getRelations("http://groupietrackers.herokuapp.com/api/relation")

	//Creating a struct which will contains all the data we want to display in the template

	data := &data{
		Artistes:  responseObjectArtist,
		Locations: responseObjectLocations,
		Dates:     responseObjectDates,
		Relation:  responseObjectRelations,
		ID:        intID,
	}

	// Call the template

	t, err := template.ParseFiles("./template/artisteDetail.html")

	if err != nil {
		error(w)
		return
	}

	// Execute the template
	t.Execute(w, data)
}
