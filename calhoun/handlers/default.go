package handlers

import (
	"fmt"
	"net/http"
)

// HandleDefault handles /
func HandleDefault(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/contacts" {
		_, _ = fmt.Fprint(w, "<h1>Contacts Page</h1>")
	} else if r.URL.Path == "/" {
		_, _ = fmt.Fprint(w, "<h1>Home Page</h1><p>Welcome to calhoun</p>")
	} else {
		HandleNotFound(w, r)
		//w.WriteHeader(http.StatusNotFound)
		//_, _ = fmt.Fprint(w, "<h1>404 Page Not Found</h1>")
	}
}
