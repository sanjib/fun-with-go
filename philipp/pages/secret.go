package pages

import (
	"fmt"
	"html/template"
	"net/http"
)

func SecretPage(w http.ResponseWriter, r *http.Request) {
	session, err := cookieStore.Get(r, cookieName)
	if err != nil {
		fmt.Println(err)
	}

	if auth, ok := session.Values[sessKeyAuthenticated].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	tmpl := template.Must(template.ParseFiles("./pages/secret.html"))
	err = tmpl.Execute(w, struct {
		User string
	}{User: session.Values[sessKeyFooBar].(string)})
	if err != nil {
		fmt.Println(err)
	}

}
