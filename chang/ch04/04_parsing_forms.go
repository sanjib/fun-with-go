package main

import (
	"fmt"
	"net/http"
)

// options to try:
//   enctype="application/x-www-form-urlencoded"
//   enctype="multipart/form-data"
var frmStr = `
<form method="post" enctype="multipart/form-data">
<input type="text" name="firstname" value="" /> <br />
<input type="text" name="lastname" value="" /> <br />
<button type="submit">Send</button>
</form>
`

// test url:
//   https://chang.lab.oak.dev/hello/?foo=bar
// response:
//   map[firstname:[Sanjib] foo:[bar] lastname:[Ahmad]]
func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method == http.MethodPost {

		// r.Form supports only application/x-www-form-urlencoded
		// and also includes key-value pairs in URL
		//_ = r.ParseForm()
		//_, _ = fmt.Fprintln(w, r.Form)

		// r.ParseMultipartForm supports only multipart/form-data
		// and ignores key-value pairs in URL
		//_ = r.ParseMultipartForm(1024)
		//_, _ = fmt.Fprintln(w, r.MultipartForm)
		//_, _ = fmt.Fprintln(w, r.Form)

		// ignores key-value pairs in URL, will print form value
		// regardless of application/x-www-form-urlencoded or
		// multipart/form-data
		//_, _ = fmt.Fprintln(w, r.Form["firstname"])

		// key-value pairs from url and form
		_, _ = fmt.Fprintln(w, r.FormValue("firstname"))

		// key-value pairs from form
		_, _ = fmt.Fprintln(w, r.PostFormValue("firstname"))
	}

	_, _ = fmt.Fprint(w, frmStr)
}

func main() {
	http.HandleFunc("/hello/", handleHello)

	s := http.Server{}
	_ = s.ListenAndServeTLS("chang-cert.pem", "chang-server.key")
}
