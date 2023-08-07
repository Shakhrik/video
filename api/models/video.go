package models

type Video struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description"`
	Duration    int64  `json:"duration"`
	CreatedAt   string `json:"created_at"`
}
