package data

import (
	"context"
	"database/sql"
	"errors"
	"nearby/common/validator"
	"time"
)

const (
	NotificationLikeType    = "Like"
	NotificationCommentType = "Comment"
)

type Notification struct {
	ID         int64     `json:"id"`
	FromUserID int64     `json:"fromUserId"`
	ToUserID   int64     `json:"toUserId"`
	PostID     int64     `json:"postId"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func ValidateNotification(v *validator.Validator, notification *Notification) {
	v.Check(notification.ToUserID != 0, "userId", "field is required")
	v.Check(notification.PostID != 0, "postId", "field is required")
	v.Check(notification.Type == NotificationLikeType || notification.Type == NotificationCommentType, "type", "invalid notification type")
}

type NotificationResponse struct {
	UserID    int64     `json:"userId"`
	PostID    int64     `json:"postId"`
	Type      string    `json:"type"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"createdAt"`
}

type INotificationModel interface {
	Insert(notification *Notification) error
	Get(userId, postId int64, notificationType string) (*Notification, error)
	GetList(toUserId int64, pagination Pagination) ([]*NotificationResponse, error)
	MarkViewed(userId int64) error
}

type NotificationModel struct {
	db *sql.DB
}

func (m NotificationModel) Insert(notification *Notification) error {
	query := `
	INSERT INTO notifications (from_user_id, to_user_id, post_id, type)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at`

	args := []any{notification.FromUserID, notification.ToUserID, notification.PostID, notification.Type}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.db.QueryRowContext(ctx, query, args...).Scan(&notification.ID, &notification.CreatedAt, &notification.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (m NotificationModel) Get(userId, postId int64, notificationType string) (*Notification, error) {
	query := `
	SELECT id, from_user_id, to_user_id, post_id, type, created_at, updated_at
	FROM notifications
	WHERE user_id=$1 AND post_id=$2 AND type=$3`

	args := []any{userId, postId, notificationType}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var notification Notification

	err := m.db.QueryRowContext(ctx, query, args...).Scan(
		&notification.ID,
		&notification.FromUserID,
		&notification.ToUserID,
		&notification.PostID,
		&notification.Type,
		&notification.CreatedAt,
		&notification.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, nil
		default:
			return nil, err
		}
	}

	return &notification, nil
}

func (m NotificationModel) GetList(toUserId int64, pagination Pagination) ([]*NotificationResponse, error) {
	query := `
	SELECT
	    post_id,
	    MIN(from_user_id) AS user_id,
	    type,
	    MAX(created_at) AS latest_notification,
	    COUNT(*) AS count
	FROM notifications
	WHERE to_user_id=$1
	GROUP BY post_id, type	
	LIMIT $2 OFFSET $3`

	args := []any{toUserId, pagination.limit(), pagination.offset()}

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
			&notification.CreatedAt,
			&notification.Count,
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

func (m NotificationModel) MarkViewed(userId int64) error {
	query := `UPDATE notifications SET seen = TRUE WHERE to_user_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.db.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}

	return nil
}
