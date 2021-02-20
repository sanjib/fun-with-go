package handlers

import (
	"calhoun/webdev/views"
	"fmt"
	"net/http"
)

func HandleFAQ(w http.ResponseWriter, _ *http.Request) {
	t := views.NewView("bootstrap", "views/faq.gohtml")

	err := t.Template.ExecuteTemplate(w, t.Layout, struct {
		Title string
	}{
		Title: "FAQ",
	})
	if err != nil {
		fmt.Println(err)
	}
}
