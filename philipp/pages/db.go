package pages

import (
	"fmt"
	"net/http"
	"time"
)

// DBPage is used by handleFunc for /db
func DBPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey db")

	// Steps to create a table, with user
	// createTableUser()
	// insertUsers()
	// selectUser(2)
	selectUsers()
	// deleteUser(1)
}

// User corresponds to the user table in the database
type User struct {
	id        int
	username  string
	password  string
	createdAt time.Time
}

func deleteUser(userID int) {
	result, err := DB.Exec(`DELETE FROM users WHERE id=?`, userID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.RowsAffected())
}

func selectUsers() {
	rows, err := DB.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("found users:", users)
}

func selectUser(userID int) {
	user := User{}
	query := `SELECT id, username, password, created_at FROM users WHERE id=?`
	row := DB.QueryRow(query, userID)
	err := row.Scan(&user.id, &user.username, &user.password, &user.createdAt)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("found user:", user)
}

func insertUsers() {
	username := "johndoe"
	password := "secret"
	createdAt := time.Now()
	result, err := DB.Exec(`
		INSERT INTO users (username, password, created_at) values (?, ?, ?)`,
		username, password, createdAt,
	)
	if err != nil {
		fmt.Println(err)
	}
	userID, err := result.LastInsertId()
	fmt.Println("user inserted with ID", userID)

	username = "alice"
	password = "bob"
	createdAt = time.Now()
	result, err = DB.Exec(`
		INSERT INTO users (username, password, created_at) values (?, ?, ?)`,
		username, password, createdAt,
	)
	if err != nil {
		fmt.Println(err)
	}
	userID, err = result.LastInsertId()
	fmt.Println("user inserted with ID", userID)
}

func createTableUser() {
	query := `
		CREATE TABLE users(
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);`
	_, err := DB.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
}
