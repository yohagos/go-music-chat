package database

import (
	"database/sql"
	"go-music-chat/util"

	"github.com/spf13/viper"
	_ "modernc.org/sqlite"
)

var (
	dbFile string
)

func InitDB() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./database")
	err := viper.ReadInConfig()
	util.CheckErr(err)
	dbFile = viper.GetString("dbPath")

	createAllTables()
}

func connectDB() *sql.DB {
	db, err := sql.Open("sqlite", dbFile)
	util.CheckErr(err)

	return db
}

func closeDB(db *sql.DB){
	db.Close()
}

func createAllTables() {
	db := connectDB()
	tables := util.ReadInDatabaseTables()
	for _, v := range tables {
		_, err := db.Exec(v)
		util.CheckErr(err)
	}
}



