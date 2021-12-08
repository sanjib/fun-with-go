package main

import (
	"fmt"
	"net/url"
)

func main() {
	// extract url information
	s := "https://www.example.com:8000/user?username=joemarini"
	res, _ := url.Parse(s)
	fmt.Println("scheme\t", res.Scheme)
	fmt.Println("host\t", res.Host)
	fmt.Println("path\t", res.Path)
	fmt.Println("port\t", res.Port())
	fmt.Println("query\t", res.RawQuery)
	vals := res.Query()
	fmt.Println("vals\t", vals)

	// construct url
	newurl := &url.URL{
		Scheme:   "https",
		Host:     "www.example.com",
		Path:     "/args",
		RawQuery: "x=1&y=2",
	}
	fmt.Println("created url", newurl.String())
	newurl.Host = "joemarini.com"
	fmt.Println("created url", newurl.String())

	// newvals
	newvals := url.Values{}
	newvals.Add("x", "10")
	newvals.Add("z", "somestr")
	newurl.RawQuery = newvals.Encode()
	fmt.Println("url new vals", newurl.String())
}
