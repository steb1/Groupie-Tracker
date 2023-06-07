package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/1")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject ResponseArtist

	json.Unmarshal(responseData, &responseObject)
	//fmt.Println(responseObject)

	 for _, val := range responseObject.Membres {
		fmt.Println(val)
	 }
}

type ResponseArtist struct {
	ID int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Membres []string `json:"members"`
	CreationDate int `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`
	Locations ResponseLocations `json:"locations"`
	ConcertDates ResponseDates `json:"concerDates"`
	Relations ResponseRelation `json:"relations"`
}

type ResponseLocations struct {
	ID int `json:"id"`
	Locations []string `json:"locations"`
	Dates ResponseDates `json:"dates"`
}

type ResponseDates struct {
	ID int `json:"id"`
	Dates []string `json:"dates"`
}

type ResponseRelation struct {
	ID int `json:"id"`
	DatesLocations []string
}