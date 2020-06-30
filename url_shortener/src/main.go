package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"url_shortener/src/urlshort"
)

var yaml string
var jsonFile string

func main() {

	parseFlags()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	dat, err := ioutil.ReadFile(yaml)
	if err != nil {
		panic(err)
	}

	yamlHandler, err := urlshort.YAMLHandler(dat, mapHandler)
	if err != nil {
		panic(err)
	}

	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	jsonHandler, err := urlshort.JSONHandler(jsonData, yamlHandler)
	if err != nil {
		panic(err)
	}

	dbHandler, err := urlshort.DBHandler(jsonHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", dbHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func parseFlags() {
	flag.StringVar(&yaml, "yaml", "urls.yaml", "YAML file with the paths and urls")
	flag.StringVar(&jsonFile, "json", "urls.json", "JSON file with paths and urls")
}
