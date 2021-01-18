package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logging(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("You are being logged at:", time.Now())
		fn(w, r)
	}
}
