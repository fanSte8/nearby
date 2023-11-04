package data

import (
	"context"
	"database/sql"
	"errors"
	"nearby/common/validator"
	"strconv"
	"time"
)

type Post struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageUrl"`
	Longitude   string    `json:"longitude"`
	Latitude    string    `json:"latitude"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type PostResponse struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageUrl"`
	Distance    float32   `json:"distance"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func ValidateCoordinates(v *validator.Validator, latitude, longitude string) {
	v.Check(latitude != "", "latitude", "must be provided")

	_, err := strconv.ParseFloat(latitude, 64)
	v.Check(err == nil, "latitude", "must be valid float")

	v.Check(longitude != "", "longitude", "must be provided")

	_, err = strconv.ParseFloat(longitude, 64)
	v.Check(err == nil, "longitude", "must be valid float")
}

func ValidatePost(v *validator.Validator, p *Post) {
	v.Check(p.Description != "", "descriptions", "must be provided")

	ValidateCoordinates(v, p.Latitude, p.Longitude)
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

func (m PostModel) Update(post *Post) error {
	query := `
	UPDATE posts
	SET description=$1, image_url=$2
	WHERE id=$3`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{post.Description, post.ImageUrl, post.ID}

	_, err := m.db.ExecContext(ctx, query, args...)

	if err != nil {
		return err
	}

	return nil
}

func (m PostModel) GetLatest(userLatitude, userLongitude string, radius_meters int, pagination Pagination) ([]*PostResponse, error) {
	query := `
	SELECT id, user_id, description, image_url, distance FROM (
		SELECT 
			id,
			user_id,
			description,
			image_url,
			ST_Distance(location::geography, ST_MakePoint($1, $2)::geography) AS distance 
		FROM posts
	) 
	WHERE distance < $3 
	ORDER BY distance ASC
	LIMIT $4 OFFSET $5`

	args := []any{userLatitude, userLongitude, radius_meters, pagination.limit(), pagination.offset()}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return []*PostResponse{}, nil
		default:
			return nil, err
		}
	}

	defer rows.Close()

	posts := []*PostResponse{}

	for rows.Next() {
		var post PostResponse

		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Description,
			&post.ImageUrl,
			&post.Distance,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (m PostModel) GetPopular(userLatitude, userLongitude string, radius_meters int, userID int64) ([]Post, error) {
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
