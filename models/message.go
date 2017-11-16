package models

type Message struct {
	ID          string   `json:"id"`
	CreatedAt   uint64   `json:"created_at"`
	UserID      string   `json:"user_id"`
	UserName    string   `json:"name"`
	GroupID     string   `json:"group_id"`
	Text        string   `json:"text"`
	FavoritedBy []string `json:"favorited_by"`
}

type MessageResponse struct {
	Response Response `json:"response"`
}

type Response struct {
	Count    uint64    `json:"count"`
	Messages []Message `jons:"messages"`
}
