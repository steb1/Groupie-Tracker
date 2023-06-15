package main

import (
	"fmt"
	"log"
	"net/http"
	"groupieTracker/controllers"
)

func main() {
	fileServer := http.FileServer(http.Dir("./template"))

	http.Handle("/", fileServer)
	http.HandleFunc("/Artiste", controllers.ArtisteHandler)
	http.HandleFunc("/css/", controllers.ServeCSS)
	http.HandleFunc("/ArtisteDetail", controllers.ArtisteDetailHandler)

	//on demarre le serveur grace a ListenAndServe en renseignant un numero de port
	fmt.Printf("Starting server at port 8080\nhttp://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
