package data

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	MockUserID             = 1
	MockUserEmail          = "test@mail.com"
	MockPasswordResetToken = "pwd_token"
	MockActivationToken    = "act_token"
)

func getMockUser() *User {
	user := User{
		ID:            1,
		FirstName:     "Test",
		LastName:      "Test",
		Email:         "test@mail.com",
		ImageUrl:      "",
		Activated:     false,
		PostsRadiusKm: 10,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err := user.Password.Set("password")
	if err != nil {
		panic(fmt.Sprintf("Error creating user mock data %v", err))
	}

	return &user
}

type MockUserModel struct {
	db *sql.DB
}

func (m MockUserModel) Insert(user *User) error {
	return nil
}

func (m MockUserModel) GetById(id int64) (*User, error) {
	user := getMockUser()

	if user.ID == id {
		return user, nil
	} else {
		return nil, ErrRecordNotFound
	}
}

func (m MockUserModel) GetByEmail(email string) (*User, error) {
	user := getMockUser()

	if user.Email == email {
		return user, nil
	} else {
		return nil, ErrRecordNotFound
	}
}

func (m MockUserModel) Update(user *User) error {
	return nil
}

func (m MockUserModel) GetByToken(tokenType, tokenText string) (*User, int64, error) {
	if tokenType == ActivationToken && tokenText == MockActivationToken || tokenType == PasswordResetToken && tokenText == MockPasswordResetToken {
		return getMockUser(), 1, nil
	}

	return nil, 0, ErrRecordNotFound
}
