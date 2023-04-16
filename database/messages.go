package database

type MessageModel struct {
	ID int32 `json:"id"`
	Receiver string `json:"receiver"`
	Text string `json:"text"`
	Send_date string `json:"send_date"`
	Sender string `json:"sender"`
	Group_ID string `json:"group_id"`
}