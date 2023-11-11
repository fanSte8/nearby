package clients

type MockMailerClient struct {
}

func (m MockMailerClient) SendActivationTokenMail(recipient, token string) error {
	return nil
}

func (m MockMailerClient) SendPasswordResetTokenMail(recipient, token string) error {
	return nil
}
