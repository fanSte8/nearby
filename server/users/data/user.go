package data

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	ImageUrl  string    `json:"imageUrl"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Activated bool      `json:"activated"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserModel struct {
	db *sql.DB
}
