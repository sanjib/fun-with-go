package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var frmStrUpload = `
<form action="/upload/?ucolor=blue" enctype="multipart/form-data" method="post">
	<input type="text" name="pcolor" value="red" /> <br />
	<input type="file" name="testfile" /> <br />
	<input type="submit" /> <br />
</form>
`

func handleUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method == http.MethodPost {
		//err := r.ParseMultipartForm(1024)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//fileHeader := r.MultipartForm.File["testfile"][0]
		//file, err := fileHeader.Open()

		file, _, err := r.FormFile("testfile")
		if err == nil {
			bytes, err := ioutil.ReadAll(file)
			if err == nil {
				_, _ = fmt.Fprintln(w, string(bytes))
			}
		}
	}

	_, _ = fmt.Fprint(w, frmStrUpload)
}

func main() {
	http.HandleFunc("/upload/", handleUpload)

	s := http.Server{}
	_ = s.ListenAndServeTLS("chang-cert.pem", "chang-server.key")
}
