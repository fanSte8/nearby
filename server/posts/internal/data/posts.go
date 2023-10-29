package data

import (
	"context"
	"database/sql"
	"time"
)

type Post struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageUrl"`
	Longitude   string    `json:'-'`
	Latitude    string    `json:'-'`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type PostModel struct {
	db *sql.DB
}

func (m PostModel) Insert(post *Post) error {
	query := `
	INSERT INTO posts (user_id, description, image_url, location)
	VALUES ($1, $2, $3, ST_MakePoint($4, $5))
	RETURNING id, created_at, updated_at`

	args := []any{post.UserID, post.Description, post.ImageUrl, post.Latitude, post.Longitude}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.db.QueryRowContext(ctx, query, args...).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m PostModel) GetLatest(userLatitude, userLongitude string, userID int64) ([]Post, error) {
	return []Post{}, nil
}

func (m PostModel) GetPopular(userLatitude, userLongitude string, userID int64) ([]Post, error) {
	return []Post{}, nil
}

func (m PostModel) Delete(id int64) error {
	query := `DELETE FROM posts WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrRecordNotFound
	}

	return nil
}
