package pages

import (
	"fmt"
	"net/http"
)

// HomePage returns the page requested by http.handleFunc at path "/"
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<img src=\"/static/assets/smile.gif\"/><h3>Hello, you've requested %s</h3>", r.URL.Path)
}
