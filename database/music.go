package database

import (
	"go-music-chat/util"
)

type SongsModel struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Featuring   string `json:"featuring"`
	Genre       string `json:"genre"`
	Path        string `json:"path"`
	Uploaded_at string `json:"uploaded_at"`
	Uploaded_by string `json:"uploaded_by"`
}

func GetAllSongs() []SongsModel {
	db := connectDB()
	rows, err := db.Query("select * from music")
	util.CheckErr(err)

	var data []SongsModel

	for rows.Next() {
		var res SongsModel
		if err = rows.Scan(&res.ID, &res.Title, &res.Artist, &res.Featuring, &res.Genre, &res.Path, &res.Uploaded_at, &res.Uploaded_by); err != nil {
			panic(err.Error())
		}
		data = append(data, res)
	}

	closeDB(db)
	return data
}

func AddNewSong(song SongsModel) {
	db := connectDB()
	stmt, err := db.Prepare("INSERT INTO music(title, artist, featuring, genre, path, uploaded_at, uploaded_by) values (?, ?, ?, ?, ?, ?, ?)")
	util.CheckErr(err)

	path := util.CreateUserFolder(song.Uploaded_by + "/" + song.Title)

	_, err = stmt.Exec(song.Title, song.Artist, song.Featuring, song.Genre, path, util.CreateTimestamp(), song.Uploaded_by)
	util.CheckErr(err)

	defer stmt.Close()
	closeDB(db)
}

