package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"./pages"
	_ "github.com/go-sql-driver/mysql"
)

func initDB() {
	db, err := sql.Open("mysql", "root@(127.0.0.1:3306)/va_go_test?parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	// Manually inject db connection to pages package
	pages.DB = db
}

func main() {
	initDB()
	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/book/view/page/", pages.BookViewPage)
	http.HandleFunc("/contact-us/", pages.ContactUsPage)
	http.HandleFunc("/db/", pages.DBPage)
	http.HandleFunc("/todos/", pages.TodosPage)
	http.HandleFunc("/", pages.HomePage)

	server := http.Server{
		Addr:    ":3000",
		Handler: nil,
	}

	// Create 2 funcs in certkey.go that returns the path to the
	// cert and key files:
	// - localCertKey() (string string)
	// - remoteCertKey() (string string)
	//
	// Toggle for local or remote: localCertKey() or remoteCertKey()
	cert, key := localCertKey()

	fmt.Println("--> Go TLS server starting at port 3000...")
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal(err)
	}
}
