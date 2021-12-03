// get_response gets a URL and returns the value of the
// Content-Type response HTTP header. Examples:
//
// get_response.go -url="https://google.com"
// Content-Type: text/html; charset=ISO-8859-1
//
// get_response.go -url="https://golang.org"
// Content-Type: text/html; charset=utf-8
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	url := flag.String("url", "", "URL to get and return Content-Type, example: https://google.com")
	flag.Parse()

	ct, err := contentType(*url)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Content-Type:", ct)
}

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	return ct, nil
}
