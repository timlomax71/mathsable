package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var templates = template.Must(template.ParseGlob("www/templates/*"))

func handler(w http.ResponseWriter, r *http.Request) {
	render(w, "layout")
}

func render(w http.ResponseWriter, template string) {
	err := templates.ExecuteTemplate(w, template, nil)
	if err != nil {
		log.Print("Failed to render template: ", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8081"
	}

	log.Print("Listening on port: " + port)

	r := mux.NewRouter().StrictSlash(true)
	css := http.StripPrefix("/css/", http.FileServer(http.Dir("./www/static/css/")))
	r.PathPrefix("/css/").Handler(css)
	fa := http.StripPrefix("/fontawesome/", http.FileServer(http.Dir("./www/static/fontawesome/")))
	r.PathPrefix("/fontawesome/").Handler(fa)
	f := http.StripPrefix("/fonts/", http.FileServer(http.Dir("./www/static/fonts/")))
	r.PathPrefix("/fonts/").Handler(f)
	img := http.StripPrefix("/img/", http.FileServer(http.Dir("./www/static/img/")))
	r.PathPrefix("/img/").Handler(img)
	js := http.StripPrefix("/js/", http.FileServer(http.Dir("./www/static/js/")))
	r.PathPrefix("/js/").Handler(js)

	r.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":"+port, r))

	// http.HandleFunc(STATIC_URL, StaticHandler)
}
