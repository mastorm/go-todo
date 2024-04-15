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

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	// create tables
	if _, err := db.ExecContext(ctx, gotodo.DDL); err != nil {
		log.Fatal(err)
	}

	queries := store.New(db)
	app := gotodo.Application{
		Queries: queries,
	}
	app.Serve()
}
