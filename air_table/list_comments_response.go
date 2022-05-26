package air_table

type ListCommentsResponse struct {
	Records []ListCommentsResponseRecord `json:"records"`
}

type ListCommentsResponseRecord struct {
	Fields Comment `json:"fields"`
}

type Comment struct {
	ID        string `json:"id"`
	PostID    string `json:"post_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}
