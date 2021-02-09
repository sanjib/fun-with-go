package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/about/", middlewareAuth(handleAbout))
	http.HandleFunc("/login/", middlewareAuth(handleLogin))
	http.HandleFunc("/", middlewareAuth(handleHome))
	fmt.Println("starting server at port 80...")
	http.ListenAndServe(":http", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println(AuthCookie.Value)
	fmt.Fprintln(w, "hello cookie home!")
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Println(AuthCookie.Value)
	fmt.Fprintln(w, "hello about!")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	AuthCookie.set("foo", "bar", w)
	fmt.Println(AuthCookie.Value)
	fmt.Fprintln(w, "login")
}

func middlewareAuth(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		AuthCookie.init(w, r)
		f(w, r)
	}
}

// AuthCookie tracks users (guest or logged in) accessing the site
var AuthCookie cookie

type cookie http.Cookie

func (c *cookie) init(w http.ResponseWriter, r *http.Request) {
	clientCookie, err := r.Cookie("uid")
	if err == http.ErrNoCookie {
		c.set("", "", w)
		return
	}
	AuthCookie = cookie(*clientCookie)
}

func (c *cookie) set(email, token string, w http.ResponseWriter) {
	newCookie := &http.Cookie{
		Name:     "uid",
		Value:    strings.Join([]string{email, token}, ":"),
		Path:     "/",
		MaxAge:   int(time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, newCookie)
	AuthCookie = cookie(*newCookie)
}
