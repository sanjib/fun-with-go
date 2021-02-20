package handlers

import (
	"calhoun/webdev/views"
	"fmt"
	"net/http"
)

// HandleDefault handles / and 404 errors
func HandleDefault(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		v := views.NewView("bootstrap", "views/home.gohtml")

		err := v.Render(w, struct {
			Title string
		}{
			Title: "Home Page Calhoun",
		})
		if err != nil {
			fmt.Println(err)
		}

	} else {
		HandleNotFound(w, r)
	}
}
