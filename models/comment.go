package models

type Comment struct {
	ID         int    `json:"id"`
	UserID     string `json:"user_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
	MovieID    int    `json:"movieID"`
}
