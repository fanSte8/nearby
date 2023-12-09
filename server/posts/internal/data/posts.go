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
	Liked       bool      `json:"liked"`
	Likes       int       `json:"likes"`
	Comments    int       `json:"comments"`
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

type IPostModel interface {
	Insert(post *Post) error
	Update(post *Post) error
	GetById(id int64) (*Post, error)
	GetPosts(sort string, userId int64, userLatitude, userLongitude string, radius_meters int, pagination Pagination) ([]*PostResponse, error)
	GetPost(postId, userId int64, userLatitude, userLongitude string) (*PostResponse, error)
	GetUserPost(currentUserId, targetUserId int64, userLatitude, userLongitude string, pagination Pagination) ([]*PostResponse, error)
	Delete(id int64) error
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

func (m PostModel) GetById(id int64) (*Post, error) {
	query := `
	SELECT id, user_id, description, image_url, created_at, updated_at
	FROM posts
	WHERE id = $1`

	var post Post

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.UserID,
		&post.Description,
		&post.ImageUrl,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (m PostModel) GetPosts(sort string, userId int64, userLatitude, userLongitude string, radius_meters int, pagination Pagination) ([]*PostResponse, error) {
	closestPostsQuery := `
	SELECT id, user_id, description, image_url, distance, COALESCE(liked, FALSE) as liked, likes, comments, created_at, updated_at FROM (
		SELECT 
			posts.id,
			posts.user_id,
			posts.description,
			posts.image_url,
			BOOL_OR(likes.user_id = $1) AS liked,
			COUNT(DISTINCT likes.user_id) AS likes,
			COUNT(DISTINCT comments.id) AS comments,
			ST_Distance(location::geography, ST_MakePoint($2, $3)::geography) AS distance,
			posts.created_at,
			posts.updated_at
		FROM posts
		LEFT JOIN comments ON comments.post_id = posts.id
		LEFT JOIN likes ON likes.post_id = posts.id
		GROUP BY posts.id
	) AS q
	WHERE distance < $4 
	ORDER BY distance ASC
	LIMIT $5 OFFSET $6`

	latestPostsQuery := `
	SELECT id, user_id, description, image_url, distance, COALESCE(liked, FALSE) as liked, likes, comments, created_at, updated_at FROM (
		SELECT 
			posts.id,
			posts.user_id,
			posts.description,
			posts.image_url,
			BOOL_OR(likes.user_id = $1) AS liked,
			COUNT(DISTINCT likes.id) AS likes,
			COUNT(DISTINCT comments.id) AS comments,
			ST_Distance(location::geography, ST_MakePoint($2, $3)::geography) AS distance,
			posts.created_at,
			posts.updated_at
		FROM posts
		LEFT JOIN comments ON comments.post_id = posts.id
		LEFT JOIN likes ON likes.post_id = posts.id
		GROUP BY posts.id
	) AS q
	WHERE distance < $4 
	ORDER BY created_at DESC
	LIMIT $5 OFFSET $6`

	popularPostsQuery := `
	SELECT id, user_id, description, image_url, distance, COALESCE(liked, FALSE) as liked, likes, comments, created_at, updated_at FROM (
		SELECT 
			posts.id,
			posts.user_id,
			posts.description,
			posts.image_url,
			BOOL_OR(likes.user_id = $1) AS liked,
			COUNT(DISTINCT likes.id) AS likes,
			COUNT(DISTINCT comments.id) AS comments,
			ST_Distance(location::geography, ST_MakePoint($2, $3)::geography) AS distance,
			(COUNT(DISTINCT likes.id) + COUNT(DISTINCT comments.id)) / POWER(EXTRACT(EPOCH FROM NOW() - posts.created_at), 2) AS score,
			posts.created_at,
			posts.updated_at
		FROM posts
		LEFT JOIN comments ON comments.post_id = posts.id
		LEFT JOIN likes ON likes.post_id = posts.id
		GROUP BY posts.id
	) AS q 
	WHERE distance < $4 
	ORDER BY score DESC
	LIMIT $5 OFFSET $6`

	args := []any{userId, userLatitude, userLongitude, radius_meters, pagination.limit(), pagination.offset()}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var query string
	switch {
	case sort == "latest":
		query = latestPostsQuery
	case sort == "closest":
		query = closestPostsQuery
	default:
		query = popularPostsQuery
	}

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
			&post.Liked,
			&post.Likes,
			&post.Comments,
			&post.CreatedAt,
			&post.UpdatedAt,
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

func (m PostModel) GetPost(postId, userId int64, userLatitude, userLongitude string) (*PostResponse, error) {
	query := `
	SELECT 
		posts.id,
		posts.user_id,
		posts.description,
		posts.image_url,
		ST_Distance(location::geography, ST_MakePoint($2, $3)::geography) AS distance,
		CASE WHEN likes.user_id = $1 THEN TRUE ELSE FALSE END AS liked,
		COUNT(DISTINCT likes.user_id) AS likes,
		COUNT(DISTINCT comments.id) AS comments,
		posts.created_at,
		posts.updated_at
	FROM posts
	LEFT JOIN comments ON comments.post_id = posts.id
	LEFT JOIN likes ON likes.post_id = posts.id
	WHERE posts.id=$4
	GROUP BY posts.id, likes.user_id`

	args := []any{userId, userLatitude, userLongitude, postId}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var post PostResponse

	err := m.db.QueryRowContext(ctx, query, args...).Scan(
		&post.ID,
		&post.UserID,
		&post.Description,
		&post.ImageUrl,
		&post.Distance,
		&post.Liked,
		&post.Likes,
		&post.Comments,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (m PostModel) GetUserPost(currentUserId, targetUserId int64, userLatitude, userLongitude string, pagination Pagination) ([]*PostResponse, error) {
	query := `
	SELECT id, user_id, description, image_url, distance, COALESCE(liked, FALSE) as liked, likes, comments, created_at, updated_at FROM (
		SELECT 
			posts.id,
			posts.user_id,
			posts.description,
			posts.image_url,
			BOOL_OR(likes.user_id = $1) AS liked,
			COUNT(DISTINCT likes.id) AS likes,
			COUNT(DISTINCT comments.id) AS comments,
			ST_Distance(location::geography, ST_MakePoint($2, $3)::geography) AS distance,
			(COUNT(DISTINCT likes.id) + COUNT(DISTINCT comments.id)) / POWER(EXTRACT(EPOCH FROM NOW() - posts.created_at), 2) AS score,
			posts.created_at,
			posts.updated_at
		FROM posts
		LEFT JOIN comments ON comments.post_id = posts.id
		LEFT JOIN likes ON likes.post_id = posts.id
		GROUP BY posts.id
	) AS q 
	WHERE user_id = $4 
	ORDER BY score DESC
	LIMIT $5 OFFSET $6`

	args := []any{currentUserId, userLatitude, userLongitude, targetUserId, pagination.limit(), pagination.offset()}

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
			&post.Liked,
			&post.Likes,
			&post.Comments,
			&post.CreatedAt,
			&post.UpdatedAt,
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
