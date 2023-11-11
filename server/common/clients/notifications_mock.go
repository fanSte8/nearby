package clients

type MockNotificationsClient struct {
}

func (m MockNotificationsClient) CreateNotification(input CreateNotificationInput) error {
	return nil
}
