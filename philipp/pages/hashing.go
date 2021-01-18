package pages

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

func HashPassword(password string) (string, error) {
	//bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type HashingPageDetails struct {
	PasswordValue string
	HashValue     string
	MatchValue    bool
}

func HashingPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./pages/hashing.html"))

	password := "JohnnyWasASailor2021#?!!"
	hash, _ := HashPassword(password)
	match := CheckPasswordHash(password, hash)

	if err := tmpl.Execute(w, HashingPageDetails{
		PasswordValue: password,
		HashValue:     hash,
		MatchValue:    match,
	}); err != nil {
		fmt.Println(err)
	}
}
