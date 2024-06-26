package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"

	"github.com/mastorm/go-todo"
	"github.com/mastorm/go-todo/store"
	_ "github.com/mattn/go-sqlite3"
)

func openDatabase(ctx context.Context, connString string) (*store.Queries, error) {
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}

	// TODO: When moving the database out of memory, this needs to be conditional
	if _, err := db.ExecContext(ctx, gotodo.DDL); err != nil {
		return nil, err
	}

	return store.New(db), nil
}

func main() {
	ctx := context.Background()
	queries, err := openDatabase(ctx, ":memory:")
	if err != nil {
		log.Fatal(err.Error())
	}

	app := gotodo.Application{
		Queries: queries,
	}
	app.Serve()
}
