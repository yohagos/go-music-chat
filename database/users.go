package database

type UsersModel struct {
	ID            int32  `json:"id"`
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Profile_photo string `json:"profile_photo"`
	Created_at    string `json:"created_at"`
}
