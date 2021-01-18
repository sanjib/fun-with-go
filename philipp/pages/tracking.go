package pages

import (
	"fmt"
	"html/template"
	"net/http"
)

func TrackingPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./pages/tracking.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
