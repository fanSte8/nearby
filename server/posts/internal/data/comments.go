package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Comment struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userId"`
	PostID    int64     `json:"postId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CommentModel struct {
	db *sql.DB
}

func (m CommentModel) Insert(comment *Comment) error {
	query := `
	INSERT INTO comments (user_id, post_id, text)
	VALUES ($1, $2, $3)
	RETURNING id, created_at, updated_at`

	args := []any{comment.UserID, comment.PostID, comment.Text}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.db.QueryRowContext(ctx, query, args...).Scan(&comment.ID, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m CommentModel) GetById(id int64) (*Comment, error) {
	query := `
	SELECT id, user_id, post_id, text, created_at, updated_at
	FROM comments
	WHERE id = $1`

	var comment Comment

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.db.QueryRowContext(ctx, query, id).Scan(
		&comment.ID,
		&comment.UserID,
		&comment.PostID,
		&comment.Text,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &comment, nil
}

func (m CommentModel) GetList(postId int64, pagination Pagination) ([]*Comment, error) {
	query := `
	SELECT id, user_id, post_id, text 
	FROM comments
	WHERE post_id=$1
	LIMIT $2 OFFSET $3`

	args := []any{postId, pagination.limit(), pagination.offset()}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return []*Comment{}, nil
		default:
			return nil, err
		}
	}

	defer rows.Close()

	comments := []*Comment{}

	for rows.Next() {
		var comment Comment

		err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.PostID,
			&comment.Text,
		)

		if err != nil {
			return nil, err
		}

		comments = append(comments, &comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (m CommentModel) Delete(id int64) error {
	query := `DELETE FROM comments WHERE id = $1`

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
