package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	newurl := url.URL{
		Scheme: "https",
		Host:   "httpbin.org",
		Path:   "/post",
	}
	hbinPostURL := newurl.String()

	// post raw json
	json := strings.NewReader(`
	{
		"field1": "this is field 1",
		"field2" : 250
	}
	`)
	res, err := http.Post(hbinPostURL, "application/json", json)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("status", res.Status)
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()
	fmt.Println("res", string(bytes))

	// post form
	data := url.Values{}
	data.Add("field1", "field added via url.Values")
	data.Add("field2", "300")
	resPost, err := http.PostForm(hbinPostURL, data)
	if err != nil {
		log.Println(err)
	}
	resPostBytes, err := ioutil.ReadAll(resPost.Body)
	if err != nil {
		log.Println(err)
	}
	defer resPost.Body.Close()
	fmt.Println("res post form", string(resPostBytes))
}
