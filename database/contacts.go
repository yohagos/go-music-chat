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

func GetAllContacts(user string) []ContactModel{
	db := connectDB()
	stmt, err := db.Prepare("select * from contacts where user = ?")
	util.CheckErr(err)
	defer stmt.Close()

	rows, err := stmt.Query(user, user)
	util.CheckErr(err)

	var contacts []ContactModel
	for rows.Next() {
		var res ContactModel
		if err = rows.Scan(&res.ID, &res.User, &res.Contact, &res.Since); err != nil {
			panic(err.Error())
		}
		
		contacts = append(contacts, res)
	}

	closeDB(db)
	return contacts
}

func RemoveContact(user, contact string) {
	db := connectDB()
	defer closeDB(db)
	stmt1, err := db.Prepare("delete from contacts where user = ? and contact = ?")
	util.CheckErr(err)
	defer stmt1.Close()

	_, err = stmt1.Exec(user, contact)
	util.CheckErr(err)
	_, err = stmt1.Exec(contact, user)
	util.CheckErr(err)
}