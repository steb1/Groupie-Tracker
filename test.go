package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"
)

func test(w http.ResponseWriter, r *http.Request) {

	response, _ := http.Get("http://groupietrackers.herokuapp.com/api/artists")
	responseData, _ := io.ReadAll(response.Body)
	var responseObjectArtist ResponseArtist
	json.Unmarshal(responseData, &responseObjectArtist)

	response1, _ := http.Get("http://groupietrackers.herokuapp.com/api/locations")
	responseData1, _ := io.ReadAll(response1.Body)
	var responseObjectLocations ResponseLocations
	json.Unmarshal(responseData1, &responseObjectLocations)

	data := &data {
		Artistes: responseObjectArtist,
		Locations: responseObjectLocations,
	}


	fmt.Println(data.Locations)
	

	t, _ := template.ParseFiles("./template/test.html")
	
	t.Execute(w, data)
}
