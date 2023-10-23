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
	activationTokenMailPath    = "/v1/mailer/activation"
	passwordResetTokenMailPath = "/v1/mailer/password-reset"
)

type MailerClient struct {
	baseUrl *url.URL
	client  *http.Client
	logger  *slog.Logger
}

func NewMailerClient(baseUrl string) (*MailerClient, error) {
	url, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	return &MailerClient{
		baseUrl: url,
		client:  client,
	}, nil
}

func (m MailerClient) sendTokenMail(recipient, token, path string) error {
	fullURL := m.baseUrl.ResolveReference(&url.URL{Path: path})
	payload := fmt.Sprintf(`{"recipient": "%s", "token": "%s"}`, recipient, token)

	req, err := http.NewRequest("POST", fullURL.String(), bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return err
	}

	response, err := m.client.Do(req)

	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		defer response.Body.Close()
		m.logger.Error("Couldn't send email", "path", path, "statusCode", response.StatusCode, "responseBody", response.Body)
		return errors.New("couldn't send email")
	}

	return nil
}

func (m MailerClient) SendActivationTokenMail(recipient, token string) error {
	return m.sendTokenMail(recipient, token, activationTokenMailPath)
}

func (m MailerClient) SendPasswordResetTokenMail(recipient, token string) error {
	return m.sendTokenMail(recipient, token, passwordResetTokenMailPath)
}
