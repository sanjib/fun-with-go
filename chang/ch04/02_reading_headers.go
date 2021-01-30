package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, r.Header["Accept-Encoding"])
	_, _ = fmt.Fprintln(w, r.Header.Get("Accept-Encoding"))
	_, _ = fmt.Fprintln(w, r.Header)
}

func main() {
	http.HandleFunc("/headers/", headers)

	s := http.Server{}
	_ = s.ListenAndServeTLS("chang-cert.pem", "chang-server.key")
}
