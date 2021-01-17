package pages

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func ContactUsPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./pages/contact_us.html"))

	if r.Method != http.MethodPost {
		if err := tmpl.Execute(w, nil); err != nil {
			fmt.Println(err)
		}
		return
	}

	contactDetails := ContactDetails{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}
	fmt.Println(contactDetails)

	if err := tmpl.Execute(w, struct {
		Success bool
	}{Success: true}); err != nil {
		fmt.Println(err)
	}
}
