package models

type Comment struct {
	ID         *int    `json:"id"`
	UserID     *int    `json:"user_id"`
	Content    *string `json:"content"`
	CreateTime *string `json:"create_time"`
}
