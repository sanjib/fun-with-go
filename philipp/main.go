package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerFuncHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<img src=\"/static/assets/smile.gif\"/><h3>Hello, you've requested %s</h3>", r.URL.Path)
}

func main() {
	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlerFuncHome)

	server := http.Server{
		Addr:    ":3000",
		Handler: nil,
	}

	// Create 2 funcs in main pkg certkey.go that returns the path to the
	// cert and key files:
	// - localCertKey() (string string)
	// - remoteCertKey() (string string)
	//
	// For example:
	// func localCertKey() (string, string) {
	//     cert := filepath.Join("[dir containing cert]", "cert")
	// 	   key := filepath.Join("[dir containing key]", "key")
	// return cert, key
	// }
	cert, key := localCertKey()
	// cert, key := remoteCertKey()

	fmt.Println("--> Go server starting at port 3000...")
	log.Fatal(server.ListenAndServeTLS(cert, key))
}
