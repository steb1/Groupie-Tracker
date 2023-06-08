package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("http://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var responseObjectArtist ResponseArtist

	json.Unmarshal(responseData, &responseObjectArtist)
	fmt.Println(responseObjectArtist[0])
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
	ID             int `json:"id"`
	DatesLocations []string
}
