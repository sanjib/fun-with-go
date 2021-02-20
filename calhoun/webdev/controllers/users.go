package controllers

import (
	"calhoun/webdev/views"
	"fmt"
	"net/http"
)

type Users struct {
	NewView *views.View
}

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parsed correctly and should only be used during
// initial setup.
//
// GET /signup
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

func (u *Users) New(w http.ResponseWriter, _ *http.Request) {
	err := u.NewView.Render(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Users.Create")
}
