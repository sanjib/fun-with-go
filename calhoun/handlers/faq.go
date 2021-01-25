package handlers

import (
	"fmt"
	"net/http"
)

func HandleFAQ(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "<h1>FAQ</h1><p>1...2...3...</p>")
}
