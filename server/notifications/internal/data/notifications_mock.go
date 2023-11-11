package data

import "time"

type MockNotificationsModel struct {
}

func (m MockNotificationsModel) Insert(notification *Notification) error {
	return nil
}

func (m MockNotificationsModel) Get(userId, postId int64, notificationType string) (*Notification, error) {
	return &Notification{
		ID:         1,
		FromUserID: 1,
		ToUserID:   2,
		PostID:     1,
		Type:       NotificationLikeType,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}

func (m MockNotificationsModel) GetList(toUserId int64, pagination Pagination) ([]*NotificationResponse, error) {
	return []*NotificationResponse{
		{
			UserID:    1,
			PostID:    1,
			Type:      NotificationLikeType,
			Count:     1,
			CreatedAt: time.Now(),
		},
	}, nil
}
