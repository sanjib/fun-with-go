package main

import (
	"fmt"
	"net/http"
	"strings"
)

// handles /
func handleAll(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/contacts" {
		fmt.Fprint(w, "<h1>Contacts Page</h1>")
	} else if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Home Page</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>404 Page Not Found</h1>")
	}
}

// handles /api/movie/
func handleAPIMovie(w http.ResponseWriter, r *http.Request) {
	pathUnits := strings.Split(r.URL.Path, "/")
	movieName := "Not specified"
	if pathUnits[3] != "" {
		movieName = pathUnits[3]
	}
	w.Header().Set("Content-Type", "text/json")
	fmt.Fprintf(w, "{ \"title\": \"%s\" }", movieName)
}

// handles /api/movies
func handleAPIMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json")
	fmt.Fprint(w, "[{ \"title\": \"Indiana Jones and the Temple of Doom\" }, { \"title\": \"The Secret of My Success\" }]")
}

// handles /hello/
func handleHello(w http.ResponseWriter, r *http.Request) {
	subPath := strings.Split(r.URL.Path, "/")
	name := ""
	if subPath[2] != "" {
		name = subPath[2]
	}
	fmt.Fprintf(w, "<h1>Hello %s</h1>", name)
}
