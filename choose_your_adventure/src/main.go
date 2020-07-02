package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

var arcs map[string]Arc

func main() {

	byteValue, err := ioutil.ReadFile("../assets/gopher.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(byteValue, &arcs)

	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	foundArc := arcs[path[1:]]
	t, _ := template.ParseFiles("../templates/index.html")
	t.Execute(w, foundArc)
}
