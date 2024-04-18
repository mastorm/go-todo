-- name: CreateTodo :one
INSERT INTO todos (
	task,
	done
) 
VALUES(?, ?)
RETURNING *;

-- name: ListTodos :many
SELECT *
FROM todos;

-- name: UpdateTodo :one
UPDATE todos
SET task = ?, done = ?
WHERE id = ?
RETURNING *;
