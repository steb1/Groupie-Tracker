package main

import (
	"fmt"
	"groupieTracker/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.HomeHandler)
	http.HandleFunc("/Artiste", controllers.ArtisteHandler)
	http.HandleFunc("/css/", controllers.ServeCSS)
	http.HandleFunc("/ArtisteDetail", controllers.ArtisteDetailHandler)

	//on demarre le serveur grace a ListenAndServe en renseignant un numero de port
	fmt.Printf("Starting server at port 8080\nhttp://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
