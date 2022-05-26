package air_table

type ListUsersResponse struct {
	Records []ListUsersResponseRecord `json:"records"`
}

type ListUsersResponseRecord struct {
	Fields User `json:"fields"`
}

type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
