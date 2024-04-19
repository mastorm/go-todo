package gotodo

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mastorm/go-todo/store"
)

type Application struct {
	Queries *store.Queries
}

func (app *Application) Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /todos", http.HandlerFunc(app.ListTodos))
	mux.Handle("POST /todos", http.HandlerFunc(app.CreateTodo))
	mux.Handle("PUT /todos/{id}", http.HandlerFunc(app.UpdateTodo))

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

func (app *Application) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Task string `json:"task"`
		Done bool   `json:"done"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	changes := store.UpdateTodoParams{
		ID:   int64(id),
		Task: payload.Task,
		Done: btoi(payload.Done),
	}

	updated, err := app.Queries.UpdateTodo(r.Context(), changes)

	writeJson(w, updated, http.StatusOK)
}

func (app *Application) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := app.Queries.ListTodos(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	writeJson(w, todos, http.StatusOK)
}
