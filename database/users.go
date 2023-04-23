package database

import (
	"database/sql"
	"encoding/json"
	"go-music-chat/util"

	_ "modernc.org/sqlite"
)

type UsersModel struct {
	ID            int32  `json:"id"`
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Profile_photo string `json:"profile_photo"`
	Created_at    string `json:"created_at"`
}

func GetAllUsers(db *sql.DB) {
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
		println(string(jsonBytes))
	}
}

func InsertNewUser(user UsersModel, db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO users(firstname, lastname, username, password, profile_photo, created_at) values (?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	util.CreateUserFolder(user.Username)

	_, err = stmt.Exec("büs", "geyik", "ergül", "baburch", "", util.CreateTimestamp())
	if err != nil {
		panic(err.Error())
	}
}

func FindUser(user string, db *sql.DB) {
	stmt, err := db.Prepare("SELECT * FROM users WHERE username == ?")
	if err != nil {
		panic(err.Error())
	}

	rows, err := stmt.Query(user)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var res UsersModel
		if err = rows.Scan(&res.ID, &res.Firstname, &res.Lastname, &res.Username, &res.Password, &res.Profile_photo, &res.Created_at); err != nil {
			panic(err.Error())
		}

		jsonBytes, _ := json.Marshal(res)
		println(string(jsonBytes))
	}
}

func UpdateProfilePhoto(user string, photo string, db *sql.DB) {
	stmt, err := db.Prepare("select * from users where username = ?")
	util.CheckErr(err)

	rows, err := stmt.Query(user)
	util.CheckErr(err)

	var res UsersModel
	for rows.Next() {
		if err = rows.Scan(&res.ID, &res.Firstname, &res.Lastname, &res.Username, &res.Password, &res.Profile_photo, &res.Created_at); err != nil {
			panic(err.Error())
		}
	}
	
	stmt, err = db.Prepare("update users set profile_photo = ? where username = ?")
	util.CheckErr(err)

	_, err = stmt.Exec(photo, user)
	util.CheckErr(err)
}

func DeleteUser(user string, db *sql.DB) {
	stmt, err := db.Prepare("DELETE from users where username = ?")
	util.CheckErr(err)

	_, err = stmt.Exec(user)
	util.CheckErr(err)
}
