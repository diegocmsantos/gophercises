package server

import (
	"gophercises/choose_your_adventure/src/models"
	"html/template"
	"log"
	"net/http"
)

var arcs map[string]models.Arc

const templatesPath = "server/templates/index.html"

// StartHTTPServer starts a HTTP server for the game
func StartHTTPServer(arcsMap map[string]models.Arc) {
	arcs = arcsMap
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	foundArc := arcs[path[1:]]
	t, err := template.ParseFiles(templatesPath)
	if err != nil {
		panic(err)
	}
	t.Execute(w, foundArc)
}
