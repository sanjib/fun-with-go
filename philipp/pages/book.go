package pages

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// BookViewPage returns the page requested by http.handleFunc at path "/book/view/page/"
func BookViewPage(w http.ResponseWriter, r *http.Request) {
	title, pageNr, err := bookViewPageVars(r.URL.Path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "You requested the book %q, page %d", title, pageNr)
}

func bookViewPageVars(path string) (string, int, error) {
	// /book/view/page/:title/:page
	a := make([]string, 5)
	a = strings.Split(path, "/")

	if a[4] != "" && a[5] != "" {
		title := a[4]
		if page, err := strconv.Atoi(a[5]); err == nil {
			return title, page, nil
		}
		return "", 0, errors.New("There was an error getting the page number")
	}
	return "", 0, errors.New("There was an error getting the page number and title")
}
