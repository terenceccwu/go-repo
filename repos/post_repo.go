package repos

import (
	"go-repo/air_table"
	"go-repo/models"
)

type PostRepo struct {
	AirTable    *air_table.AirTable
	CommentRepo *CommentRepo
	results     []*models.Post
}

func NewPostRepo(AirTable *air_table.AirTable) *PostRepo {
	return &PostRepo{
		CommentRepo: NewCommentRepo(AirTable),
		AirTable:    AirTable,
	}
}

func (r *PostRepo) Query() *PostRepo {
	res := r.AirTable.ListPosts()

	for _, record := range res.Records {
		post := models.PostFromAirTable(&record.Fields)
		r.results = append(r.results, post)
	}

	return r
}

func (r *PostRepo) Result() []*models.Post {
	return r.results
}

func (r *PostRepo) IncludeComments(opts ...interface{}) *PostRepo {
	comments := r.CommentRepo.Query().Result()

	mapping := map[string]*models.Post{}
	for _, post := range r.results {
		mapping[post.ID] = post
	}

	for _, comment := range comments {
		mapping[comment.PostID].Comments = append(mapping[comment.PostID].Comments, comment)
	}

	return r
}

func (r *PostRepo) WithComments(cb func(*CommentRepo)) *PostRepo {
	cb(r.CommentRepo)
	return r
}
