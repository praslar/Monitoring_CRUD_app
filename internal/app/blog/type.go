package blog

import "time"

type (
	Blog struct {
		ID        string    `json:"blog_id"`
		Title     string    `json:"title,omitempty"`
		Content   string    `json:"content,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty"`
		UpdateAt  time.Time `json:"updated_at,omitempty" `
	}
	CreateBlog struct {
		Title   string `json:"title,omitempty"`
		Content string `json:"content,omitempty"`
	}
)
