package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

const (
	Beef   = "Beef"
	Lamb   = "Lamb"
	Fish   = "Fish"
	Shrimp = "Shrimp"
	Crab   = "Crab"
)

type User struct {
	Name          string
	Age           int
	FavoriteFoods []string
}

func HandleTest(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("./hello.gohtml"))

	user := User{
		Name:          "Sanjib",
		Age:           47,
		FavoriteFoods: []string{Beef, Lamb, Crab},
	}
	colors := []string{"blue", "red", "white", "green", "orange", "purple", "yellow"}

	_ = tmpl.Execute(w, struct {
		User            User
		Colors          []string
		ConsoleLogMsgHi string
	}{
		User:            user,
		Colors:          colors,
		ConsoleLogMsgHi: template.JSEscapeString("Hello gohtml ") + fmt.Sprintf("%v", time.Now()),
	})
}

func main() {
	http.HandleFunc("/", HandleTest)
	fmt.Println("--> starting server at port 3000...")
	_ = http.ListenAndServe(":3000", nil)
}
