package models

type Content struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}
