package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"log"
	"net/http"
	//"os"
)

func main() {
	response, _ := http.Get("http://groupietrackers.herokuapp.com/api/artists")
	response1, _ := http.Get("http://groupietrackers.herokuapp.com/api/locations")
	response2, _:= http.Get("http://groupietrackers.herokuapp.com/api/dates")
	response3, _ := http.Get("http://groupietrackers.herokuapp.com/api/relation")

	//////////////////////////////////////////

	responseData, _ := ioutil.ReadAll(response.Body)
	responseData1, _ := ioutil.ReadAll(response1.Body)
	responseData2, _ := ioutil.ReadAll(response2.Body)
	responseData3, _ := ioutil.ReadAll(response3.Body)

	////////////////////////////////////////

	var responseObjectArtist ResponseArtist
	var responseObjectLocations ResponseLocations
	var responseObjectDates ResponseDates
	var responseObjectRelations ResponseRelation

	json.Unmarshal(responseData, &responseObjectArtist)

	json.Unmarshal(responseData1, &responseObjectLocations)

	json.Unmarshal(responseData2, &responseObjectDates)

	json.Unmarshal(responseData3, &responseObjectRelations)

	/* fmt.Println(responseObjectArtist[0])
	fmt.Println(responseObjectLocations.Index[0])
	fmt.Println(responseObjectDates.Index[0]) */
	fmt.Println(responseObjectRelations.Index[1] )

}

type ResponseArtist []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type ResponseLocations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type ResponseDates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type ResponseRelation struct {
	Index []struct {
	ID             int              `json:"id"`            
	DatesLocations map[string][]string `json:"datesLocations"`
	}	
}
