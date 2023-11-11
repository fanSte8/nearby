package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

type Models struct {
	Notification INotificationModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Notification: NotificationModel{db},
	}
}

func NewMockModels() Models {
	return Models{
		Notification: MockNotificationsModel{},
	}
}
