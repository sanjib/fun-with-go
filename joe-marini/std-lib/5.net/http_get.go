package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	const hbinGet = "https://httpbin.org/get"

	// get
	res, err := http.Get(hbinGet)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	fmt.Println("res status:\t", res.Status)
	fmt.Println("res status c:\t", res.StatusCode)
	fmt.Println("res protocol:\t", res.Proto)
	fmt.Println("res con len:\t", res.ContentLength)

	var sb strings.Builder
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	n, err := sb.Write(bytes)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("res body:\t", n, sb.String())
}
