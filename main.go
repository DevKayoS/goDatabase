package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	urlExample := "postgres://pg:password@localhost:8541/test"
	db, err := pgxpool.New(context.Background(), urlExample)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}

	// query := "create table foo (id bigserial primary key, bar varchar(255))"
	// if _, err := db.Exec(context.Background(), query); err != nil {
	// 	panic(err)
	// }

	query := "insert into foo (bar) values ($1);"

	if _, err := db.Exec(context.Background(), query, "abcdsa"); err != nil {
		panic(err)
	}

	query = "select * from foo;"
	type foobar struct {
		id  int64
		bar string
	}

	var foo foobar
	if err := db.QueryRow(context.Background(), query).Scan(&foo.id, &foo.bar); err != nil {
		panic(err)
	}

	fmt.Printf("%#+v\n", foo)
}
