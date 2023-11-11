package data

import "time"

func GetMockPost() *Post {
	return &Post{
		ID:          1,
		UserID:      1,
		Description: "Test post",
		ImageUrl:    "",
		Longitude:   "10",
		Latitude:    "10",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func GetMockPostResponses() []*PostResponse {
	return []*PostResponse{
		{
			ID:          1,
			UserID:      1,
			Description: "Test post response 1",
			ImageUrl:    "",
			Distance:    10,
			Liked:       true,
			Likes:       10,
			Comments:    10,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			UserID:      1,
			Description: "Test post response 2",
			ImageUrl:    "",
			Distance:    10,
			Liked:       false,
			Likes:       10,
			Comments:    10,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          3,
			UserID:      1,
			Description: "Test post response 3",
			ImageUrl:    "",
			Distance:    10,
			Liked:       true,
			Likes:       10,
			Comments:    10,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
}

type MockPostModel struct {
}

func (m MockPostModel) Insert(post *Post) error {
	return nil
}

func (m MockPostModel) Update(post *Post) error {
	return nil
}

func (m MockPostModel) GetById(id int64) (*Post, error) {
	if id == 1 {
		return GetMockPost(), nil
	}

	return nil, ErrRecordNotFound
}

func (m MockPostModel) GetPosts(sort string, userId int64, userLatitude, userLongitude string, radius_meters int, pagination Pagination) ([]*PostResponse, error) {
	return GetMockPostResponses(), nil
}

func (m MockPostModel) GetPost(postId, userId int64, userLatitude, userLongitude string) (*PostResponse, error) {
	if postId == 1 {
		return GetMockPostResponses()[0], nil
	}

	return nil, ErrRecordNotFound
}

func (m MockPostModel) Delete(id int64) error {
	return nil
}
