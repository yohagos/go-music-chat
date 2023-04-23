package database

import "go-music-chat/util"

type ContactRequestModel struct {
	ID        int32  `json:"id"`
	User      string `json:"user"`
	Requested string `json:"requested"`
}

func CreateRequest(user, request string) {
	db := connectDB()
	stmt, err := db.Prepare("Insert Into contactRequests(user, requested) values(?, ?)")
	util.CheckErr(err)
	defer stmt.Close()

	_, err = stmt.Query(user, request)
	util.CheckErr(err)
	closeDB(db)
}

func DeclineRequest(user, request string) {
	db := connectDB()
	stmt, err := db.Prepare("Delete from contactRequests where user = ? AND requested = ?")
	util.CheckErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(user, request)
	util.CheckErr(err)
	_, err = stmt.Exec(request, user)
	util.CheckErr(err)
	println("called declineRequest")
}

func AcceptRequest(user, request string) {
	db := connectDB()
	stmt, err := db.Prepare("Select * from contactRequests where user = ? And requested = ?")
	util.CheckErr(err)
	defer stmt.Close()

	rows, err := stmt.Query(user, request)
	util.CheckErr(err)
	
	var req ContactRequestModel
	for rows.Next() {
		if err = rows.Scan(&req.ID, &req.User, &req.Requested); err != nil {
			panic(err.Error())
		}
	}

	stmt, err = db.Prepare("Select * from contactRequests where user = ? And requested = ?")
	util.CheckErr(err)
	rows, err = stmt.Query(request, user)
	util.CheckErr(err)
	
	for rows.Next() {
		if err = rows.Scan(&req.ID, &req.User, &req.Requested); err != nil {
			panic(err.Error())
		}
	}
	
	DeclineRequest(user, request)
	closeDB(db)
}
