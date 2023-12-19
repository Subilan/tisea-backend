package structs

import "time"

type DatabaseDynamic struct {
	ID         uint64    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Author     string    `json:"author"`
	Hidden     bool      `json:"hidden"`
	Categories string    `json:"categories"`
	Tags       string    `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateDynamicRequest struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Author     string `json:"author"`
	Categories string `json:"categories"`
	Tags       string `json:"tags"`
}
