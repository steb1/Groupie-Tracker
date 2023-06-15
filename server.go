package main

import (
	"fmt"
	"groupieTracker/controllers"
	"log"
	"net/http"
	"text/template"
)

func main() {
	//fileServer := http.FileServer(http.Dir("./templates"))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/Artiste", controllers.ArtisteHandler)
	http.HandleFunc("/css/", controllers.ServeCSS)
	http.HandleFunc("/ArtisteDetail", controllers.ArtisteDetailHandler)

	//on demarre le serveur grace a ListenAndServe en renseignant un numero de port
	fmt.Printf("Starting server at port 8080\nhttp://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

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
