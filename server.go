package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/Artiste", ArtisteHandler)
	http.HandleFunc("/css/", ServeCSS)
	http.HandleFunc("/ArtisteDetail", ArtisteDetailHandler)

	//http.HandleFunc("/css/", ServeCSS)

	//on demarre le serveur grace a ListenAndServe en renseignant un numero de port
	fmt.Printf("Starting server at port 8080\nhttp://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	response, _ := http.Get("http://groupietrackers.herokuapp.com/api/artists")
	response1, _ := http.Get("http://groupietrackers.herokuapp.com/api/locations")
	response2, _ := http.Get("http://groupietrackers.herokuapp.com/api/dates")
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

	/* fmt.Println(responseObjectArtist[0].FirstAlbum)
	fmt.Println(responseObjectLocations.Index[0])
	fmt.Println(responseObjectDates.Index[0]) */
	//fmt.Println(responseObjectRelations.Index[1].DatesLocations["noumea-new_caledonia"])

}

func ArtisteDetailHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query()["ID"]
	IDD := ID[0]
	intID, err := strconv.Atoi(IDD) 

	if err != nil {
		t, _ := template.ParseFiles("./template/404.html")
		p := ""
		t.Execute(w, p)
		return
	}

	if IDD == "0" || len(ID) == 0 || intID > 57 || IDD == "" {
		t, _ := template.ParseFiles("./template/404.html")
		p := ""
		t.Execute(w, p)
		return
	}

	t, err := template.ParseFiles("./template/artisteDetail.html")

	if err != nil {
		log.Fatal(err)
	}



	response, _ := http.Get("http://groupietrackers.herokuapp.com/api/artists")
	responseData, _ := ioutil.ReadAll(response.Body)
	var responseObjectArtist ResponseArtist
	json.Unmarshal(responseData, &responseObjectArtist)



	t.Execute(w, responseObjectArtist[intID-1])
}

func ServeCSS(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path
	http.ServeFile(w, r, "template"+filename)
}

func ArtisteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Artiste" {
		t, _ := template.ParseFiles("./template/404.html")
		p := ""
		t.Execute(w, p)
		return
	}

	t, err := template.ParseFiles("./template/artistes.html")

	if err != nil {
		log.Fatal(err)
	}
	response, _ := http.Get("http://groupietrackers.herokuapp.com/api/artists")
	responseData, _ := ioutil.ReadAll(response.Body)
	var responseObjectArtist ResponseArtist
	json.Unmarshal(responseData, &responseObjectArtist)

	//fmt.Println(responseObjectArtist)

	t.Execute(w, responseObjectArtist)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		t, _ := template.ParseFiles("./template/404.html")
		p := ""
		t.Execute(w, p)
		return
	}

	t, _ := template.ParseFiles("./template/index.html")
	p := ""
	t.Execute(w, p)
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
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	}
}
