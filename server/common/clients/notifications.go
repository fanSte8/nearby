package clients

import (
	"bytes"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
)

const (
	createNotification = "/internal/v1/notifications/"
)

const (
	LikeNotificationType    = "Like"
	CommentNotificationType = "Comment"
)

type NotificationsClient struct {
	baseUrl *url.URL
	client  *http.Client
	logger  *slog.Logger
}

type CreateNotificationInput struct {
	FromUserID int64
	ToUserID   int64
	PostID     int64
	Type       string
}

func NewNotificationsClient(baseUrl string) (*NotificationsClient, error) {
	url, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	return &NotificationsClient{
		baseUrl: url,
		client:  client,
	}, nil
}

func (n NotificationsClient) CreateNotification(input CreateNotificationInput) error {
	fullURL := n.baseUrl.ResolveReference(&url.URL{Path: createNotification})

	payload := fmt.Sprintf(`{"fromUserId": "%d", "toUserId": "%d", "postId": "%d", "type": "%s"}`, input.FromUserID, input.ToUserID, input.PostID, input.Type)

	req, err := http.NewRequest(http.MethodPost, fullURL.String(), bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return err
	}

	response, err := n.client.Do(req)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		defer response.Body.Close()
		n.logger.Error("Couldn't create notification", "statusCode", response.StatusCode, "responseBody", response.Body)
		return errors.New("couldn't send email")
	}

	return nil
}
