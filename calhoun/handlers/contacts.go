package handlers

import (
	"calhoun/views"
	"fmt"
	"net/http"
)

func HandleContacts(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	v := views.NewView("bootstrap", "views/contacts.gohtml")

	err := v.Render(w, struct {
		Title string
	}{
		Title: "Contacts",
	})
	if err != nil {
		fmt.Println(err)
	}
}
