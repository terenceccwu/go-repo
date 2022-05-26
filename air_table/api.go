package air_table

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AirTable struct {
	BaseURL string
	Client  *http.Client
	Token   string
}

func (a *AirTable) Request(request *Request, response interface{}) error {
	fmt.Printf("Requesting: [%s] %s\n", request.Method, request.Url)

	req, _ := http.NewRequest(request.Method, a.BaseURL+request.Url, nil)
	req.Header.Set("Authorization", "Bearer "+a.Token)

	res, err := a.Client.Do(req)
	if err != nil {
		return err
	}

	json.NewDecoder(res.Body).Decode(&response)

	return nil
}

func (a *AirTable) ListPosts() *ListPostsResponse {
	res := &ListPostsResponse{}
	a.Request(&Request{
		Method: "GET",
		Url:    "/Posts",
	}, res)

	return res
}

func (a *AirTable) ListComments() *ListCommentsResponse {
	res := &ListCommentsResponse{}
	a.Request(&Request{
		Method: "GET",
		Url:    "/Comments",
	}, res)

	return res
}

func (a *AirTable) ListUsers() *ListUsersResponse {
	res := &ListUsersResponse{}
	a.Request(&Request{
		Method: "GET",
		Url:    "/Users",
	}, res)

	return res
}
