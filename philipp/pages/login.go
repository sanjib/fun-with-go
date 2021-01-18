package pages

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type LoginPageDetails struct {
	Success bool
	Message string
}

func loginFormInitial(tmpl *template.Template, w http.ResponseWriter) {
	if err := tmpl.Execute(w, LoginPageDetails{
		Success: false,
		Message: "",
	}); err != nil {
		fmt.Println(err)
	}
}

func loginFormInvalidAttempt(tmpl *template.Template, w http.ResponseWriter) {
	if err := tmpl.Execute(w, LoginPageDetails{
		Success: false,
		Message: fmt.Sprintf("inavlid user/pass: %v",
			time.Now()),
	}); err != nil {
		fmt.Println(err)
	}
}

func loginPageSuccessfullyLoggedIn(tmpl *template.Template, w http.ResponseWriter) {
	if err := tmpl.Execute(w, LoginPageDetails{
		Success: true,
		Message: fmt.Sprintf("you are now logged in: %v",
			time.Now()),
	}); err != nil {
		fmt.Println(err)
	}
}

func loginPageAlreadyLoggedIn(tmpl *template.Template, w http.ResponseWriter) {
	if err := tmpl.Execute(w, LoginPageDetails{
		Success: true,
		Message: fmt.Sprintf("you are already logged in: %v",
			time.Now()),
	}); err != nil {
		fmt.Println(err)
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	// 1. Get session from cookie
	session, err := cookieStore.Get(r, cookieName)
	if err != nil {
		fmt.Println(err)
	}

	tmpl := template.Must(template.ParseFiles("./pages/login.html"))

	if auth, ok := session.Values[sessKeyAuthenticated].(bool); !ok || !auth {
		// 2. Show login form

		// Did not post yet - Initial login form
		if r.Method != http.MethodPost {
			loginFormInitial(tmpl, w)
			return
		}

		// Posted
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "sanjib" && password == "foobar" {
			// Success
			// Save session
			session.Values[sessKeyFooBar] = username
			session.Values[sessKeyAuthenticated] = true
			err = session.Save(r, w)
			if err != nil {
				fmt.Println(err)
			}

			loginPageSuccessfullyLoggedIn(tmpl, w)
			return
		}

		// Invalid user/pass
		loginFormInvalidAttempt(tmpl, w)
		return

	}

	loginPageAlreadyLoggedIn(tmpl, w)
}
