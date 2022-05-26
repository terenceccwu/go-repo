package models

import "go-repo/air_table"

type Comment struct {
	ID        string `json:"id"`
	PostID    string `json:"post_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserID    string `json:"user_id"`
	User      *User  `json:"user"`
	CreatedAt string `json:"created_at"`
}

func CommentFromAirTable(res *air_table.Comment) *Comment {
	return &Comment{
		ID:        res.ID,
		PostID:    res.PostID,
		Title:     res.Title,
		Body:      res.Body,
		UserID:    res.UserID,
		CreatedAt: res.CreatedAt,
	}
}
