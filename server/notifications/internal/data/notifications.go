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
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type NotificationResponse struct {
	UserID    int64     `json:"userId"`
	PostID    int64     `json:"postId"`
	Type      string    `json:"type"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"createdAt"`
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

func (m NotificationModel) Get(userId int64, pagination Pagination) ([]*NotificationResponse, error) {
	query := `
	SELECT
	    post_id,
	    MIN(user_id) AS user_id,
	    type,
	    MAX(created_at) AS latest_notification,
	    COUNT(*) AS count
	FROM notifications
	GROUP BY post_id, type`

	args := []any{userId, pagination.limit(), pagination.offset()}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return []*NotificationResponse{}, nil
		default:
			return nil, err
		}
	}

	defer rows.Close()

	notifications := []*NotificationResponse{}

	for rows.Next() {
		var notification NotificationResponse

		err := rows.Scan(
			&notification.PostID,
			&notification.UserID,
			&notification.Type,
			&notification.Count,
			&notification.CreatedAt,
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
