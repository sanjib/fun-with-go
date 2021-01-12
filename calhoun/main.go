package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/api/movies", handleAPIMovies)
	http.HandleFunc("/api/movie/", handleAPIMovie)
	http.HandleFunc("/hello/", handleHello)
	http.HandleFunc("/", handleAll)
	server := http.Server{Addr: ":3000"}
	server.ListenAndServe()
}
