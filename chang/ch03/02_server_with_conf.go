package main

import "net/http"

func main() {
	s := http.Server{
		Addr: ":http", // :http is port 80
	}
	_ = s.ListenAndServe()
}
