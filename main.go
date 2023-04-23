package main

import (
	"go-music-chat/database"
	"go-music-chat/server"
	"net/http"
	"os"
)

func main() {
	database.ConnectDB("music.db")
	database.CreateAllTables()

	os.Exit(12)
	
	router := server.NewRouter()
	
	http.ListenAndServe(":8080",router)
}