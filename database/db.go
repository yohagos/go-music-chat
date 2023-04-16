package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	id int32
}

func Connect(file string) {
	println("testing sqlite3 db..", file)
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic(err)
	}
	println(db.Stats().OpenConnections)
	defer  db.Close()
}