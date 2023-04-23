package main

import (
	"go-music-chat/database"
	"go-music-chat/server"
	"net/http"
	"os"
)

func main() {
	database.InitDB()
	// database.GetAllUsers()
	// database.GetAllSongs()
	// database.FindSongByArtist("bausa")
	// database.CreateRequest("yosie", "yosef")
	database.AcceptRequest("yosie", "yosef")

	os.Exit(12)
	
	router := server.NewRouter()
	http.ListenAndServe(":8080",router)
}