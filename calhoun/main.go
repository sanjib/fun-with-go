package main

import (
	"calhoun/handlers"
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/api/movies", handlers.HandleAPIMovies)
	http.HandleFunc("/api/movie/", handlers.HandleAPIMovie)
	http.HandleFunc("/hello/", handlers.HandleHello)
	http.HandleFunc("/contacts/", handlers.HandleContacts)
	http.HandleFunc("/faq/", handlers.HandleFAQ)
	http.HandleFunc("/", handlers.HandleDefault)

	port := "3000"
	server := http.Server{Addr: ":" + port}
	fmt.Println("--> Starting server at port", port)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
