package database

import (
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

func GetAllUsers() []UsersModel {
	db := connectDB()
	stmt := "select * from users"
	rows, err := db.Query(stmt)
	util.CheckErr(err)

	var data []UsersModel

	for rows.Next() {
		var res UsersModel
		if err = rows.Scan(&res.ID, &res.Firstname, &res.Lastname, &res.Username, &res.Password, &res.Profile_photo, &res.Created_at); err != nil {
			panic(err.Error())
		}

		data = append(data, res)
	}

	closeDB(db)
	return data
}

func InsertNewUser(user UsersModel) {
	db := connectDB()
	stmt, err := db.Prepare("INSERT INTO users(firstname, lastname, username, password, profile_photo, created_at) values (?,?,?,?,?,?)")
	util.CheckErr(err)
	defer stmt.Close()

	util.CreateUserFolder(user.Username)

	_, err = stmt.Exec("büs", "geyik", "ergül", "baburch", "", util.CreateTimestamp())
	util.CheckErr(err)
	closeDB(db)
}

func FindUser(user string) UsersModel {
	db := connectDB()
	stmt, err := db.Prepare("SELECT * FROM users WHERE username == ?")
	util.CheckErr(err)

	rows, err := stmt.Query(user)
	util.CheckErr(err)
	defer stmt.Close()

	var res UsersModel
	for rows.Next() {
		if err = rows.Scan(&res.ID, &res.Firstname, &res.Lastname, &res.Username, &res.Password, &res.Profile_photo, &res.Created_at); err != nil {
			panic(err.Error())
		}
	}

	closeDB(db)
	return res
}

func UpdateProfilePhoto(user string, photo string) {
	db := connectDB()
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
	defer stmt.Close()
	closeDB(db)
}

func DeleteUser(user string) {
	db := connectDB()
	stmt, err := db.Prepare("DELETE from users where username = ?")
	util.CheckErr(err)

	_, err = stmt.Exec(user)
	util.CheckErr(err)
	defer stmt.Close()
	closeDB(db)
}
