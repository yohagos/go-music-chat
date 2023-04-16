package models

type Contact struct {
	ID int32 `json:"id"`
	User string `json:"user"`
	Contact string `json:"contact"`
	Since string `json:"since"`
}