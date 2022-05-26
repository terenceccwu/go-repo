package air_table

type ListPostsResponse struct {
	Records []ListPostsResponseRecord `json:"records"`
}

type ListPostsResponseRecord struct {
	Fields Post `json:"fields"`
}

type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}
