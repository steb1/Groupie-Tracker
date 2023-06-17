package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"text/template"
)

func ArtisteDetailHandler(w http.ResponseWriter, r *http.Request) {

	// Recup the ID passed through the URL

	fmt.Println(r.URL.Path)

	ID := r.URL.Query().Get("ID")

	if len(ID) == 0 {
		w.WriteHeader(400)
		error(w)
		return
	}

	intID, err := strconv.Atoi(ID)

	if err != nil {
		w.WriteHeader(404)
		error(w)
		return
	}

	if ID == "0" || len(ID) == 0 || intID > 52 || ID == "" {
		t, err := template.ParseFiles("./template/404.html")
		if err != nil {
			w.WriteHeader(500)
			error(w)
			return
		}
		w.WriteHeader(404)
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

	responseObjectLocations.Index[intID].Locations = Title(responseObjectLocations.Index[intID].Locations)
	responseObjectDates.Index[intID].Dates = DateFormat(responseObjectDates.Index[intID].Dates)
	responseObjectRelations.Index[intID].DatesLocations = RelationDates(responseObjectRelations.Index[intID].DatesLocations)

	//fmt.Println(responseObjectDates.Index[intID].Dates)

	//fmt.Println(responseObjectRelations.Index[intID].DatesLocations)

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
		w.WriteHeader(500)
		error(w)
		return
	}

	// Execute the template
	t.Execute(w, data)
}
