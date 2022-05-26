package repos

import (
	"go-repo/air_table"
	"go-repo/models"
)

type UserRepo struct {
	AirTable *air_table.AirTable
	results  []*models.User
}

func NewUserRepo(AirTable *air_table.AirTable) *UserRepo {
	return &UserRepo{
		AirTable: AirTable,
	}
}

func (r *UserRepo) Query() *UserRepo {
	res := r.AirTable.ListUsers()

	for _, record := range res.Records {
		User := models.UserFromAirTable(&record.Fields)
		r.results = append(r.results, User)
	}

	return r
}

func (r *UserRepo) Result() []*models.User {
	return r.results
}
