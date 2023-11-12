package data

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"nearby/common/validator"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            int64     `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	ImageUrl      string    `json:"imageUrl"`
	Email         string    `json:"email"`
	Password      password  `json:"-"`
	Activated     bool      `json:"activated"`
	PostsRadiusKm int       `json:"postsRadiusKm"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func (u User) GetProfilePictureKey() string {
	return fmt.Sprintf("%d-profile-picture", u.ID)
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

type IUserModel interface {
	Insert(user *User) error
	GetById(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	GetByToken(tokenType, tokenText string) (*User, int64, error)
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
	SELECT id, first_name, last_name, image_url, email, password, activated, posts_radius_km, created_at, updated_at
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
		&user.PostsRadiusKm,
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
	SELECT id, first_name, last_name, image_url, email, password, activated, posts_radius_km, created_at, updated_at
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
		&user.PostsRadiusKm,
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
	SET first_name = $1, last_name = $2, email = $3, image_url = $4, password = $5, activated = $6, posts_radius_km=$7, updated_at = NOW()
	WHERE id = $8`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{
		user.FirstName,
		user.LastName,
		user.Email,
		user.ImageUrl,
		user.Password.hash,
		user.Activated,
		user.PostsRadiusKm,
		user.ID,
	}

	_, err := m.db.ExecContext(ctx, query, args...)

	if err != nil {
		return err
	}

	return nil
}

func (m UserModel) GetByToken(tokenType, tokenText string) (*User, int64, error) {
	query := `
        SELECT users.id, users.first_name, users.last_name, users.email, users.image_url, users.password, users.activated, users.posts_radius_km, users.created_at, users.updated_at, tokens.id
        FROM users
        INNER JOIN tokens
        ON users.id = tokens.user_id
        WHERE tokens.hash = $1
        AND tokens.type = $2
        AND tokens.expiry > $3`

	hash := sha256.Sum256([]byte(tokenText))

	args := []any{hash[:], tokenType, time.Now()}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User
	var tokenId int64

	err := m.db.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.ImageUrl,
		&user.Password.hash,
		&user.Activated,
		&user.PostsRadiusKm,
		&user.CreatedAt,
		&user.UpdatedAt,
		&tokenId,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, 0, ErrRecordNotFound
		default:
			return nil, 0, err
		}
	}

	return &user, tokenId, nil
}
