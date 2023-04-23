package database

import "go-music-chat/util"

type ContactModel struct {
	ID      int32  `json:"id"`
	User    string `json:"user"`
	Contact string `json:"contact"`
	Since   string `json:"since"`
}

type Contact struct {
	User    string `json:"user"`
	Contact string `json:"contact"`
	Since   string `json:"since"`
}

func CreateNewContact(user, contact string) {
	db := connectDB()
	stmt, err := db.Prepare("Insert Into contacts(user, contact, since) values(?, ?, ?)")
	util.CheckErr(err)
	defer stmt.Close()

	_, err = stmt.Query(user, contact, util.CreateTimestamp())
	util.CheckErr(err)
	closeDB(db)
}

func GetAllContacts() {
	//db := connectDB()
}