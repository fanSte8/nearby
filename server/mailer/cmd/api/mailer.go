package main

import (
	"time"

	"github.com/go-mail/mail/v2"
)

type mailer struct {
	dialer mail.Dialer
	sender string
}

func newMailer(cfg *config) mailer {
	dialer := mail.NewDialer(cfg.SmtpHost, cfg.SmtpPort, cfg.SmtpUsername, cfg.SmtpPassword)
	dialer.Timeout = 5 * time.Second

	return mailer{
		dialer: *dialer,
		sender: cfg.SmtpSender,
	}
}

func (m mailer) Send(recipient, subject, text string) error {
	msg := mail.NewMessage()

	msg.SetHeader("To", recipient)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", text)

	err := m.dialer.DialAndSend(msg)
	if err != nil {
		return err
	}

	return nil
}
