package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "user=gwp password=gwp dbname=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func main() {
	// var err error

	// Create
	// p1 := post{content: "hi shadows", author: "yohji yamamoto"}
	// err = p1.create()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(p1)
	// p2 := post{content: "aaa", author: "bbb"}
	// err = p2.create()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(p2)

	// Read, id 6 is "hi shadows"
	// p := getPost(6)
	// fmt.Println(p)

	// Read all
	ps := getPosts()
	fmt.Printf("%#v", ps)

	// Update
	// pu := post{7, "aaa2", "bbb2"}
	// err = pu.update()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Delete
	// pd := post{id: 8}
	// if err = pd.delete(); err != nil {
	// 	fmt.Println(err)
	// }

	// Read all, again
	// ps = getPosts()
	// fmt.Println(ps)
}

type post struct {
	id      int
	content string
	author  string
}

func (p *post) delete() error {
	q := "DELETE FROM posts WHERE id=$1"
	_, err := DB.Exec(q, p.id)
	return err
}

func (p *post) update() error {
	q := "UPDATE posts SET content=$2, author=$3 WHERE id=$1"
	_, err := DB.Exec(q, p.id, p.content, p.author)
	return err
}

func getPosts() []post {
	ps := []post{}
	q := "SELECT * FROM posts"
	rows, err := DB.Query(q)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		p := post{}
		err := rows.Scan(&p.id, &p.content, &p.author)
		if err != nil {
			fmt.Println(err)
		}
		ps = append(ps, p)
	}
	return ps
}

func getPost(id int) post {
	p := post{}
	q := "SELECT * FROM posts WHERE id=$1"
	row := DB.QueryRow(q, id)
	err := row.Scan(&p.id, &p.content, &p.author)
	if err != nil {
		fmt.Println(err)
	}
	return p
}

func (p *post) create() error {
	q := "INSERT INTO posts (content, author) VALUES ($1, $2) RETURNING id"
	stmt, err := DB.Prepare(q)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(p.content, p.author).Scan(&p.id)
	return err
}
