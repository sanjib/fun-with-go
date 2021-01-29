package main

import "net/http"

func main() {
	s := http.Server{}

	err := s.ListenAndServeTLS("chang-cert.pem", "chang-server.key")
	if err != nil {
		panic(err)
	}
}
