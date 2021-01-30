package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter) {
	c1 := http.Cookie{
		Name:  "first_cookie",
		Value: "chang fun with go",
		Path:  "/",
	}
	c2 := http.Cookie{
		Name:   "2nd_cookie",
		Value:  "more fun with go",
		Path:   "/",
		Secure: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func handleSaturn(w http.ResponseWriter, r *http.Request) {
	setCookie(w)
	_, _ = fmt.Fprintln(w, "hello saturn...")
}

func handleNeptune(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "hello neptune...")
	_, _ = fmt.Fprintln(w, r.Header["Cookie"])

	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Println("Can't get cookie c1")
	}
	_, _ = fmt.Fprintln(w, c1)

	_, _ = fmt.Fprintln(w, r.Cookies())
}

func main() {
	http.HandleFunc("/saturn/", handleSaturn)
	http.HandleFunc("/neptune/", handleNeptune)

	s := http.Server{}
	_ = s.ListenAndServeTLS("chang-cert.pem", "chang-server.key")
}
