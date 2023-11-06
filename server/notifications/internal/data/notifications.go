package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

const (
	NotificationLikeType    = "Like"
	NotificationCommentType = "Comment"
)

type Notification struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userId"`
	PostID    int64     `json:"postId"`
	Type      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type NotificationModel struct {
	db *sql.DB
}

func (m NotificationModel) Insert(notification *Notification) error {
	query := `
	INSERT INTO notifications (user_id, post_id, type)
	VALUES ($1, $2, $3)
	RETURNING id, created_at, updated_at`

	args := []any{notification.UserID, notification.PostID, notification.Type}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.db.QueryRowContext(ctx, query, args...).Scan(&notification.ID, &notification.CreatedAt, &notification.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m NotificationModel) Get(userId int64, pagination Pagination) ([]*Notification, error) {
	query := `
	SELECT id, user_id, post_id, type, created_at, updated_at 
	FROM notifications
	WHERE userId=$1
	LIMIT $2 OFFSET $3`

	args := []any{userId, pagination.limit(), pagination.offset()}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return []*Notification{}, nil
		default:
			return nil, err
		}
	}

	defer rows.Close()

	notifications := []*Notification{}

	for rows.Next() {
		var notification Notification

		err := rows.Scan(
			&notification.ID,
			&notification.UserID,
			&notification.PostID,
			&notification.Type,
			&notification.CreatedAt,
			&notification.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		notifications = append(notifications, &notification)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}
