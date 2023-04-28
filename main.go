package main

import (
	"go-music-chat/database"
	"go-music-chat/server"

	"net/http"
)

func main() {
	database.InitDB()
	// database.GetAllUsers()
	// database.GetAllSongs()
	// database.FindSongByArtist("bausa")
	// database.CreateRequest("yosie", "yosef")
	// database.AcceptRequest("yosie", "yosef")
	//database.GetAllContacts("admin")
	//database.RemoveContact("admin", "basir")
	//database.GetAllContacts("admin")
	/* var newMessage database.Message
	newMessage.Sender = "yosie2"
	newMessage.Receiver = "yosie"
	newMessage.Send_date = util.CreateTimestamp()
	newMessage.Text = "test5"
	newMessage.AddMessage() */
	//database.GetAllMessages("yosie2", "yosie")
	
	router := server.NewRouter()
	http.ListenAndServe(":8080",router)
}