package data

import "time"

type MockCommentModel struct {
}

func GetMockComments() []*Comment {
	return []*Comment{
		{
			ID:        1,
			UserID:    1,
			PostID:    1,
			Text:      "Comment 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			UserID:    1,
			PostID:    1,
			Text:      "Comment 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			UserID:    1,
			PostID:    1,
			Text:      "Comment 3",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func (m MockCommentModel) Insert(comment *Comment) error {
	return nil
}

func (m MockCommentModel) GetById(id int64) (*Comment, error) {
	if id > 0 && id < 4 {
		return GetMockComments()[id-1], nil
	}

	return nil, ErrRecordNotFound
}

func (m MockCommentModel) GetList(postId int64, pagination Pagination) ([]*Comment, error) {
	return GetMockComments(), nil
}

func (m MockCommentModel) Delete(id int64) error {
	return nil
}
