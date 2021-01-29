package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello world...")
}

func bye(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Goodbye...")
}

func home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Home...")
}

func auth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// authenticate
		log.Println("authenticate")
		h(w, r)
	}
}

func logging(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// authenticate
		log.Println("log")
		h(w, r)
	}
}

func main() {
	http.HandleFunc("/hello/", auth(logging(hello)))
	http.HandleFunc("/bye/", logging(bye))
	http.HandleFunc("/", home)

	s := http.Server{}
	err := s.ListenAndServeTLS("chang-cert.pem", "chang-server.key")
	//err := http.ListenAndServe(":http", nil)
	if err != nil {
		log.Fatal(err)
	}
}
