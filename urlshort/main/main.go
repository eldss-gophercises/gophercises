package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/eldss/gophercises/urlshort"
)

func main() {
	ymlPath := flag.String("yaml", "", "Path to a yaml file with an array of 'path' to 'url' maps.")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	fmt.Println("Starting the server on :8080")

	if *ymlPath == "" {
		http.ListenAndServe(":8080", mapHandler)
	} else {
		// Read YAML
		yaml, err := ioutil.ReadFile(*ymlPath)
		if err != nil {
			panic(err)
		}
		// Make handler
		yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
		if err != nil {
			panic(err)
		}
		// Serve
		http.ListenAndServe(":8080", yamlHandler)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
