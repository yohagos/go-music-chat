package database

import (
	"database/sql"
	"go-music-chat/util"

	_ "modernc.org/sqlite"
)


func ConnectDB(file string) *sql.DB {
	db, err := sql.Open("sqlite", file)
	util.CheckErr(err)

	return db
}

func CloseDB(db *sql.DB){
	db.Close()
}

func CreateAllTables() {
	db := ConnectDB("music.db")
	tables := util.ReadInDatabaseTables()
	for _, v := range tables {
		_, err := db.Exec(v)
		util.CheckErr(err)
	}
}



