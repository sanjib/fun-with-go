package main

import "net/http"

func main() {
	_ = http.ListenAndServe("", nil)
}
