package handlers

import (
	"fmt"
	"net/http"
)

// HandleAPIMovies handles /api/movies
func HandleAPIMovies(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/json")
	_, _ = fmt.Fprint(w, "[{ \"title\": \"Indiana Jones and the Temple of Doom\" }, { \"title\": \"The Secret of My Success\" }]")
}
