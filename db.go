package gotodo

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/mastorm/go-todo/store"
)

//go:embed schema.sql
var DDL string

func openDatabase(ctx context.Context, connString string) (*store.Queries, error) {
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}

	// TODO: When moving the database out of memory, this needs to be conditional
	if _, err := db.ExecContext(ctx, DDL); err != nil {
		return nil, err
	}

	return store.New(db), nil
}
