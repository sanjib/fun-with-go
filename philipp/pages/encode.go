package pages

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func EncodePage(w http.ResponseWriter, r *http.Request) {
	games := []Game{
		{
			Name:      "Super Mario World",
			Developer: "Nintendo",
			Release:   time.Date(1990, 11, 21, 0, 0, 0, 0, time.UTC),
			Genre:     "Platform",
		},
		{
			Name:      "Undertale",
			Developer: "Toby Fox",
			Release:   time.Date(2015, 9, 15, 0, 0, 0, 0, time.UTC),
			Genre:     "Role-playing",
		},
	}

	if err := json.NewEncoder(w).Encode(games); err != nil {
		fmt.Println(err)
	}
}
