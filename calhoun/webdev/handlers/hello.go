package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

// HandleHello handles /hello/
func HandleHello(w http.ResponseWriter, r *http.Request) {
	subPath := strings.Split(r.URL.Path, "/")
	name := ""
	if subPath[2] != "" {
		name = subPath[2]
	}
	_, _ = fmt.Fprintf(w, "<h1>Hello %s</h1>", name)
}
