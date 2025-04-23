package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func sqlcTern() {
	urlExample := "postgres://pg:password@localhost:8541/test"
	db, err := pgxpool.New(context.Background(), urlExample)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}

	queries := New(db)
	ctx := context.Background()

	authors, err := queries.ListAuthors(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println(authors)

	authorCreated, err := queries.CreateAuthor(ctx, CreateAuthorParams{
		Name: "Kayo Vinicius",
		Bio:  pgtype.Text{String: "Desenvolvedor de Software", Valid: true},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(authorCreated)

	authors, err = queries.ListAuthors(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println(authors)

}
