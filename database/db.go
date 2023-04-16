package database

import (
	"database/sql"
	"encoding/json"
	"time"

	_ "modernc.org/sqlite"
)

func Connect(file string) {
	println("testing sqlite3 db..", file)
	db, err := sql.Open("sqlite", file)
	if err != nil {
		panic(err)
	}

	getAllUsers(db)
	//insertNewUser(db)
	//getAllUsers(db)

	defer db.Close()
}

func getAllUsers(db *sql.DB) {
	stmt := "select * from users"
	rows, err := db.Query(stmt)

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var res UsersModel
		if err = rows.Scan(&res.ID, &res.Firstname, &res.Lastname, &res.Username, &res.Password, &res.Profile_photo, &res.Created_at); err != nil {
			panic(err.Error())
		}
		jsonBytes, _ := json.Marshal(res)
		str := string(jsonBytes)
		println(str)
	}
}

func insertNewUser(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO users(firstname, lastname, username, password, profile_photo, created_at) values (?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	currentTIme := time.Now()
	println(currentTIme.Format("2006-01-02 15:04:05"))

	_, err = stmt.Exec("büs", "geyik", "ergül", "baburch", "", currentTIme.Format("2006-01-02 15:04:05"))
	if err != nil {
		panic(err.Error())
	}
}
