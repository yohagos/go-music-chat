package database

type ContactRequestModel struct {
	ID int32 `json:"id"`
	User string  `json:"user"`
	Requested string  `json:"requested"`
}