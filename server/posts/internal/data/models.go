package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

type Models struct {
	Posts    IPostModel
	Likes    ILikeModel
	Comments ICommentModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Posts:    PostModel{db},
		Likes:    LikeModel{db},
		Comments: CommentModel{db},
	}
}

func NewMockModels() Models {
	return Models{
		Posts:    MockPostModel{},
		Likes:    MockLikeModel{},
		Comments: MockCommentModel{},
	}
}
