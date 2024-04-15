-- name: CreateTodo :one
INSERT INTO todos (
	task,
	done
) 
VALUES(?, ?)
RETURNING *;
