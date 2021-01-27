package handlers

import (
	"calhoun/views"
	"fmt"
	"net/http"
)

func HandleSignup(w http.ResponseWriter, r *http.Request) {
	t := views.NewView("bootstrap", "views/signup.gohtml")
	err := t.Render(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
