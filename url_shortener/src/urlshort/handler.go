package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url_shortener/src/postgresconnection"

	"github.com/go-yaml/yaml"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	fmt.Println(pathsToUrls)
	return func(w http.ResponseWriter, r *http.Request) {
		url := pathsToUrls[r.URL.Path]
		if url != "" {
			http.Redirect(w, r, url, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

type YAMLMap struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var ymlMap []YAMLMap
	err := yaml.Unmarshal(yml, &ymlMap)
	if err != nil {
		return nil, err
	}

	paths := make(map[string]string)
	for _, y := range ymlMap {
		paths[y.Path] = y.Url
	}

	return MapHandler(paths, fallback), nil
}

type JSONMap struct {
	Path string `json:"path"`
	URL  string `json:"URL"`
}

// JSONHandler handles json files
func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var jsonMap []JSONMap
	err := json.Unmarshal(jsonBytes, &jsonMap)

	if err != nil {
		return nil, err
	}

	paths := make(map[string]string)
	for _, y := range jsonMap {
		paths[y.Path] = y.URL
	}

	return MapHandler(paths, fallback), nil
}

// DBHandler handles database data
func DBHandler(fallback http.Handler) (http.HandlerFunc, error) {
	db, err := postgresconnection.Getconn()

	if err != nil {
		return nil, err
	}

	paths := make(map[string]string)
	sqlStatement := `SELECT id, path, url FROM urls;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var path, url string
		err = rows.Scan(&id, &path, &url)
		if err != nil {
			// handle this error
			panic(err)
		}
		paths[path] = url
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return MapHandler(paths, fallback), nil
}
