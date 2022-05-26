package models

import "go-repo/air_table"

type Post struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	UserID    string     `json:"user_id"`
	CreatedAt string     `json:"created_at"`
	Comments  []*Comment `json:"comments"`
}

func PostFromAirTable(res *air_table.Post) *Post {
	return &Post{
		ID:        res.ID,
		Title:     res.Title,
		Body:      res.Body,
		UserID:    res.UserID,
		CreatedAt: res.CreatedAt,
	}
}
