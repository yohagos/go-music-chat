package database

type SongsModel struct {
	ID int32 `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Featuring string `json:"featuring"`
	Genre string `json:"genre"`
	Path string `json:"path"`
	Uploaded_at string `json:"uploaded_at"`
	Uploaded_by string `json:"uploaded_by"`
}