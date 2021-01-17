package pages

import (
	"fmt"
	"html/template"
	"net/http"
)

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

// Todo struct
type Todo struct {
	ID    int
	Title string
	Done  bool
}

// TodosPage serves /todos
func TodosPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./pages/todos.html"))

	// createTodosTable()
	// insertTodos()
	todos := selectTodos()

	data := TodoPageData{
		PageTitle: "List of All Todos",
		Todos:     todos,
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func selectTodos() []Todo {
	rows, err := DB.Query(`SELECT * FROM todos`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		todo := Todo{}
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Done)
		if err != nil {
			fmt.Println(err)
		}
		todos = append(todos, todo)
	}
	return todos
}

func insertTodos() {
	todos := []Todo{
		Todo{Title: "Task 1", Done: false},
		Todo{Title: "Task 2", Done: true},
		Todo{Title: "Task 3", Done: true},
	}
	for _, todo := range todos {
		res, err := DB.Exec(`INSERT INTO todos (title, done) values (?, ?)`,
			todo.Title, todo.Done,
		)
		if err != nil {
			fmt.Println(err)
		}

		taskID, err := res.LastInsertId()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("inserted todo with ID:", taskID)
	}
}

func createTodosTable() {
	query := `CREATE TABLE todos(
		id INT AUTO_INCREMENT,
		title text NOT NULL,
		done tinyint(1) NOT NULL,
		PRIMARY KEY (id)
	)`
	res, err := DB.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
