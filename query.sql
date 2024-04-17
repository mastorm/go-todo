-- name: CreateTodo :one
INSERT INTO todos (
	task,
	done
) 
VALUES(?, ?)
RETURNING *;

-- name: ListTodos :many
SELECT *
FROM todos
