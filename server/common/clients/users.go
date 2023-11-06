package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	getUserByIDPath = "/internal/v1/users/"
)

type UsersClient struct {
	baseUrl *url.URL
	client  *http.Client
}

type User struct {
	ID            int64     `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	ImageUrl      string    `json:"imageUrl"`
	Email         string    `json:"email"`
	Activated     bool      `json:"activated"`
	PostsRadiusKm int       `json:"postsRadiusKm"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type UserData struct {
	User User `json:"user"`
}

func NewUsersClient(baseUrl string) (*UsersClient, error) {
	url, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	return &UsersClient{
		baseUrl: url,
		client:  client,
	}, nil
}

func (u UsersClient) GetUserByID(id int64) (*UserData, error) {
	fullURL := u.baseUrl.ResolveReference(&url.URL{Path: fmt.Sprintf("%s%d", getUserByIDPath, id)})

	req, err := http.NewRequest(http.MethodGet, fullURL.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err := u.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	userData := UserData{}

	err = json.NewDecoder(response.Body).Decode(&userData)
	if err != nil {
		return nil, err
	}

	return &userData, nil
}
