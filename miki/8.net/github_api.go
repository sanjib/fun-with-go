package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

func main() {
	user, err := userInfo("sanjib")
	if err != nil {
		log.Println(err)
	}

	log.Println(user)
}

func userInfo(login string) (*User, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s", login))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//var jsonRaw string
	//scanner := bufio.NewScanner(resp.Body)
	//for scanner.Scan() {
	//	jsonRaw = scanner.Text()
	//}

	user := &User{}
	//err = json.Unmarshal([]byte(jsonRaw), user)
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}
