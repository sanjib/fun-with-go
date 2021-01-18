package pages

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DecodePage(w http.ResponseWriter, r *http.Request) {
	var game Game
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		fmt.Println(err)
	}

	if _, err := fmt.Fprintf(w,
		"The game %q was developed by %v",
		game.Name, game.Developer,
	); err != nil {
		fmt.Println(err)
	}
}
