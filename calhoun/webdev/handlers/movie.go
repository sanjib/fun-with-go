package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

// HandleAPIMovie handles /api/movie/
func HandleAPIMovie(w http.ResponseWriter, r *http.Request) {
	pathUnits := strings.Split(r.URL.Path, "/")
	movieName := "Not specified"
	if pathUnits[3] != "" {
		movieName = pathUnits[3]
	}
	w.Header().Set("Content-Type", "text/json")
	_, _ = fmt.Fprintf(w, "{ \"title\": \"%s\" }", movieName)
}
