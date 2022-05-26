package models

import "go-repo/air_table"

type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func UserFromAirTable(res *air_table.User) *User {
	return &User{
		ID:          res.ID,
		DisplayName: res.DisplayName,
		Email:       res.Email,
		Password:    res.Password,
	}
}
