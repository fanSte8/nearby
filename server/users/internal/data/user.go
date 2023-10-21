package data

import (
	"context"
	"database/sql"
	"errors"
	"nearby/common/validator"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	ImageUrl  string    `json:"imageUrl"`
	Email     string    `json:"email"`
	Password  password  `json:"-"`
	Activated bool      `json:"activated"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type password struct {
	text *string
	hash []byte
}

func (p *password) Set(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	p.text = &password
	p.hash = hash

	return nil
}

func (p *password) Matches(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

type UserModel struct {
	db *sql.DB
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePassword(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 characters long")
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.FirstName != "", "firstName", "must be provided")
	v.Check(len(user.FirstName) <= 50, "firstName", "must not be more than 50 characters long")

	v.Check(user.LastName != "", "LastName", "must be provided")
	v.Check(len(user.LastName) <= 50, "lastName", "must not be more than 50 characters long")

	ValidateEmail(v, user.Email)

	if user.Password.text != nil {
		ValidatePassword(v, *user.Password.text)
	}

	if user.Password.hash == nil {
		panic("missing password hash for user")
	}
}

func (m UserModel) Insert(user *User) error {
	query := `
	INSERT INTO users (first_name, last_name, email, image_url, password, activated)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, created_at, updated_at`

	args := []any{user.FirstName, user.LastName, user.Email, "", user.Password.hash, false}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.db.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m UserModel) GetById(id int64) (*User, error) {
	query := `
	SELECT id, first_name, last_name, image_url, email, password, activated, created_at, updated_at
	FROM users
	WHERE id = $1`

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.ImageUrl,
		&user.Email,
		&user.Password.hash,
		&user.Activated,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (m UserModel) GetByEmail(email string) (*User, error) {
	query := `
	SELECT id, first_name, last_name, image_url, email, password, activated, created_at, updated_at
	FROM users
	WHERE email = $1`

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.ImageUrl,
		&user.Email,
		&user.Password.hash,
		&user.Activated,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}
func (m UserModel) Update(user *User) error {
	query := `
	UPDATE users
	SET first_name = $1, last_name = $2, image_url = $3, password = $4, activated = $5
	WHERE id = $6`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{
		user.FirstName,
		user.LastName,
		user.ImageUrl,
		user.Password.hash,
		user.Activated,
		user.ID,
	}

	err := m.db.QueryRowContext(ctx, query, args).Err()

	if err != nil {
		return err
	}

	return nil
}
