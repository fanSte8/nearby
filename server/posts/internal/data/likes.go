package data

import (
	"context"
	"database/sql"
	"time"
)

type ILikeModel interface {
	Exists(userId, postId int64) (bool, error)
	Insert(userId, postId int64) error
	Delete(userId, postId int64) error
}

type LikeModel struct {
	db *sql.DB
}

func (m LikeModel) Exists(userId, postId int64) (bool, error) {
	query := `SELECT COUNT(*) FROM likes WHERE user_id=$1 AND post_id=$2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var likes int
	err := m.db.QueryRowContext(ctx, query, userId, postId).Scan(&likes)

	if err != nil {
		return false, err
	}

	return likes == 1, nil
}

func (m LikeModel) Insert(userId, postId int64) error {
	query := `INSERT INTO likes (user_id, post_id) VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.db.QueryRowContext(ctx, query, userId, postId).Err()

	if err != nil {
		return err
	}

	return nil
}

func (m LikeModel) Delete(userId, postId int64) error {
	query := `DELETE FROM likes WHERE user_id = $1 AND post_id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.db.ExecContext(ctx, query, userId, postId)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
