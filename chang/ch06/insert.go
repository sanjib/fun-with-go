package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func init() {
	var err error
	DB, err = sql.Open("postgres", "user=gwp password=gwp dbname=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
	fmt.Println(DB)
}

// Post corresponds to the post table
type Post struct {
	id      int
	content string
	author  string
}

// Create post record
func (p *Post) Create() error {
	query := "INSERT INTO posts (content, author) VALUES ($1, $2) RETURNING id"
	stmt, err := DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(p.content, p.author).Scan(&p.id)
	return err
}

// DB connection
var DB *sql.DB

func main() {
	post := Post{content: "hello world", author: "krebos"}
	fmt.Println(post)

	err := post.Create()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(post)
}
