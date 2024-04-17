package gotodo

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/mastorm/go-todo/store"
	"net/http"
)

//go:embed schema.sql
var DDL string

type Application struct {
	Queries *store.Queries
}

func (app *Application) Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /todos", http.HandlerFunc(app.ListTodos))
	mux.Handle("POST /todos", http.HandlerFunc(app.CreateTodo))

	http.ListenAndServe(":9001", mux)
}

func (app *Application) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Task string `json:"task"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	createTodoParams := store.CreateTodoParams{Task: payload.Task, Done: 0}
	todo, err := app.Queries.CreateTodo(r.Context(), createTodoParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(todo)

	w.Write([]byte("Hello World"))
}

func (app *Application) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := app.Queries.ListTodos(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	writeJson(w, todos, http.StatusOK)
}
