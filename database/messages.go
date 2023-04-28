package database

import "go-music-chat/util"

type MessageModel struct {
	ID int32 `json:"id"`
	Receiver string `json:"receiver"`
	Text string `json:"text"`
	Send_date string `json:"send_date"`
	Sender string `json:"sender"`
	Group_ID string `json:"group_id"`
}

type Message struct {
	Receiver string `json:"receiver"`
	Text string `json:"text"`
	Send_date string `json:"send_date"`
	Sender string `json:"sender"`
	Group_ID string `json:"group_id"`
}

func (m *Message) AddMessage(){
	db := connectDB()
	defer closeDB(db)

	stmt, err := db.Prepare("insert into messages(receiver, text, send_date, sender, group_id) values(?,?,?,?,?)")
	util.CheckErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(m.Receiver, m.Text, m.Send_date, m.Sender, m.Group_ID)
	util.CheckErr(err)
}

func GetAllMessages(user, contact string) []MessageModel {
	db := connectDB()
	defer closeDB(db)

	stmt, err := db.Prepare("select * from messages where (sender = ? and receiver = ?) or (sender = ? and receiver = ?)")
	util.CheckErr(err)
	defer stmt.Close()

	var messages []MessageModel
	rows, err := stmt.Query(user, contact, contact, user)
	util.CheckErr(err)

	for rows.Next() {
		var msg MessageModel
		if err = rows.Scan(&msg.ID, &msg.Receiver, &msg.Text, &msg.Send_date, &msg.Sender, &msg.Group_ID); err != nil {
			panic(err.Error())
		}
		messages = append(messages, msg)
		println(msg.Receiver + " " + msg.Text + " " + msg.Sender)
	}
	return messages
}