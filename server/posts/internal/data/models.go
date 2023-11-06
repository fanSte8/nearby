package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

type Models struct {
	Posts    PostModel
	Likes    LikeModel
	Comments CommentModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Posts:    PostModel{db},
		Likes:    LikeModel{db},
		Comments: CommentModel{db},
	}
}
