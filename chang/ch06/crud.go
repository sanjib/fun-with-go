package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DB is the database handle
var DB *sql.DB

func init() {
	var err error
	// first parameter "postgres" is the driver
	DB, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func main() {
	// create()
	// read()
	// readAll()
	// update()
	// delete()
}

func create() {
	fmt.Println(DB)
	post := Post{content: "Hello Mars!!!", author: "Jibsan"}
	fmt.Println(post) // the id is 0 because it's default int value

	post.create()
	fmt.Println(post) // the id is now incremented based on the last insert id
}

func read() {
	p1, err := getPost(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("p1", p1)

	p5, err := getPost(5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("p5", p5)
}

func readAll() {
	posts, err := getPosts()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("posts", posts)
}

func update() {
	post := Post{2, "hello world rev2", "sanjib rev2"}
	err := post.update()
	if err != nil {
		fmt.Println(err)
	}

	postAfterUpdate, err := getPost(post.id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("post after update", postAfterUpdate)
}

func delete() {
	p := Post{id: 3}
	err := p.delete()
	if err != nil {
		fmt.Println(err)
	}
}

// Post struct provides "post" records and corresponds to the posts table.
type Post struct {
	id      int
	content string
	author  string
}

func getPosts() ([]Post, error) {
	p := []Post{}
	q := "SELECT * FROM POSTS"
	rows, err := DB.Query(q)
	if err != nil {
		return p, err
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.id, &post.content, &post.author)
		if err != nil {
			return p, err
		}
		p = append(p, post)
	}
	return p, err
}

func getPost(id int) (Post, error) {
	p := Post{}
	q := "SELECT * FROM posts WHERE id=$1"
	err := DB.QueryRow(q, id).Scan(&p.id, &p.content, &p.author)
	return p, err
}

func (p *Post) create() error {
	q := "INSERT INTO posts (content, author) VALUES ($1, $2) RETURNING id"
	stmt, err := DB.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(p.content, p.author).Scan(&p.id)
	return err
}

func (p *Post) update() error {
	q := "UPDATE posts SET content=$2, author=$3 WHERE id=$1"
	res, err := DB.Exec(q, p.id, p.content, p.author)
	fmt.Println("res", res)
	return err
}

func (p *Post) delete() error {
	q := "DELETE FROM posts WHERE id=$1"
	res, err := DB.Exec(q, p.id)
	fmt.Println("res", res)
	return err
}
