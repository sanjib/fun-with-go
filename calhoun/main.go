package main

import (
	"fmt"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello cal...1,2,3...!!</h1>")
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
	http.HandleFunc("/", handleHome)
	server := http.Server{Addr: ":3000"}
	server.ListenAndServe()
}
