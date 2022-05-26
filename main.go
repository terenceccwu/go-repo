package main

import (
	"fmt"
	"go-repo/air_table"
	"go-repo/repos"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	AirTable := &air_table.AirTable{
		BaseURL: os.Getenv("AIRTABLE_BASE_URL"),
		Token:   os.Getenv("AIRTABLE_ACCESS_TOKEN"),
		Client:  &http.Client{},
	}

	PostRepo := repos.NewPostRepo(AirTable)

	posts := PostRepo.
		Query().
		IncludeComments().
		WithComments(func(c *repos.CommentRepo) {
			c.IncludeUsers()
		}).
		Result()

	for _, post := range posts {
		fmt.Printf("-----\n")
		fmt.Printf("Post ID: %s\n", post.ID)
		fmt.Printf("Post Title: %s\n", post.Title)
		fmt.Printf("Post Body: %s\n", post.Body)
		fmt.Printf("Post Comments:\n")
		for _, comment := range post.Comments {
			fmt.Printf(" * Comment ID: %s; by %s: %s\n", comment.ID, comment.User.DisplayName, comment.Title)
		}
	}
}
