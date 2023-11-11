package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

type Models struct {
	Users  IUserModel
	Tokens ITokenModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:  UserModel{db},
		Tokens: TokenModel{db},
	}
}

func NewMockModels() Models {
	return Models{
		Users:  MockUserModel{},
		Tokens: MockTokenModel{},
	}
}
