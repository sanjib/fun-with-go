package pages

import (
	"fmt"
	"net/http"
	"time"
)

func LogoutPage(w http.ResponseWriter, r *http.Request) {
	// 1. Get session from cookie
	session, err := cookieStore.Get(r, cookieName)
	if err != nil {
		fmt.Println(err)
	}

	// 2. Set session values: authenticated = false
	session.Values[sessKeyFooBar] = ""
	session.Values[sessKeyAuthenticated] = false

	// 3. Save session
	err = session.Save(r, w)
	if err != nil {
		fmt.Println(err)
	}

	_, _ = fmt.Fprintf(w, "will logout at: %v", time.Now())
}
