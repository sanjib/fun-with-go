package main

import (
	"fmt"
	"net/http"
)

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

func handleAPIMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json")
	fmt.Fprint(w, "{ \"title\": \"Indiana Jones and the Temple of Doom\" }")
}

func handleAPIMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json")
	fmt.Fprint(w, "[{ \"title\": \"Indiana Jones and the Temple of Doom\" }, { \"title\": \"The Secret of My Success\" }]")
}

func main() {
	http.HandleFunc("/api/movies", handleAPIMovies)
	http.HandleFunc("/api/movie", handleAPIMovie)
	http.HandleFunc("/", handleAll)
	server := http.Server{Addr: ":3000"}
	server.ListenAndServe()
}
