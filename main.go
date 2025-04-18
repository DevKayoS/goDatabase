package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}

	createTableSql := `
		CREATE TABLE foo(
			id integer not null primary key,
			name text
		);
	`
	res, err := db.Exec(createTableSql)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.RowsAffected())

	insertSql := `
		INSERT INTO foo(id, name) VALUES (1, 'teste')
	`

	res, err = db.Exec(insertSql)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.RowsAffected())

	type user struct {
		ID   int64
		Name string
	}

	querySql := `
		SELECT * FROM foo where id = ?
	`

	var u user
	if err := db.QueryRow(querySql, 1).Scan(&u.ID, &u.Name); err != nil {
		panic(err)
	}

	fmt.Println(u)

	for id := 2; id < 10; id++ {
		if _, err := db.Exec("insert into foo (id) values (?)", id); err != nil {
			panic(err)
		}
	}

	var count int64
	if err := db.QueryRow("select count(*) from foo").Scan(&count); err != nil {
		panic(err)
	}
	fmt.Println(count)

	input := "drop database"
	// deleteSql := fmt.Sprintf(`
	// 	DELETE FROM foo WHERE id = %s
	// `, input)

	deleteSql := `
		delete from foo where id = ?
	`

	if _, err := db.Exec(deleteSql, input); err != nil {
		panic(err)
	}

	if err := db.QueryRow("select count(*) from foo").Scan(&count); err != nil {
		panic(err)
	}
	fmt.Println(count)

}
