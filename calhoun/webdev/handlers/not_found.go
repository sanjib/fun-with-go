package handlers

import (
	"fmt"
	"net/http"
)

func HandleNotFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = fmt.Fprintf(w, "404 not found")
}
