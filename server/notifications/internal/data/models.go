package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

type Models struct {
	Notification NotificationModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Notification: NotificationModel{db},
	}
}
