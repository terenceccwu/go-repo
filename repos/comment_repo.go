package repos

import (
	"go-repo/air_table"
	"go-repo/models"
)

type CommentRepo struct {
	AirTable *air_table.AirTable
	UserRepo *UserRepo
	results  []*models.Comment
}

func NewCommentRepo(AirTable *air_table.AirTable) *CommentRepo {
	return &CommentRepo{
		UserRepo: NewUserRepo(AirTable),
		AirTable: AirTable,
	}
}

func (r *CommentRepo) Query() *CommentRepo {
	res := r.AirTable.ListComments()

	for _, record := range res.Records {
		comment := models.CommentFromAirTable(&record.Fields)
		r.results = append(r.results, comment)
	}

	return r
}

func (r *CommentRepo) Result() []*models.Comment {
	return r.results
}

func (r *CommentRepo) IncludeUsers() *CommentRepo {
	users := r.UserRepo.Query().Result()

	mapping := map[string]*models.User{}
	for _, user := range users {
		mapping[user.ID] = user
	}

	for _, comment := range r.results {
		comment.User = mapping[comment.UserID]
	}

	return r
}
