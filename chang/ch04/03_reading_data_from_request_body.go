package main

import (
	"fmt"
	"net/http"
)

func handleBody(w http.ResponseWriter, r *http.Request) {
	conLen := r.ContentLength
	body := make([]byte, conLen)
	_, _ = r.Body.Read(body)
	_, _ = fmt.Fprintln(w, string(body))
}

func main() {
	// To test type:
	//   curl -id "foo=bar" localhost/body
	http.HandleFunc("/body/", handleBody)

	//s := http.Server{}
	//_ = s.ListenAndServeTLS("chang-cert.pem", "chang-server.key")
	_ = http.ListenAndServe(":http", nil)
}
