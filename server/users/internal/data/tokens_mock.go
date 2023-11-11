package data

import (
	"crypto/sha256"
	"time"
)

func getMockToken(userID int64, ttl time.Duration, tokenType string) *Token {
	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
		Type:   tokenType,
	}

	var value string
	switch {
	case tokenType == ActivationToken:
		value = MockActivationToken
	case tokenType == PasswordResetToken:
		value = MockPasswordResetToken
	default:
		value = ""
	}

	token.Text = value
	hash := sha256.Sum256([]byte(value))
	token.Hash = hash[:]

	return token
}

type MockTokenModel struct {
}

func (m MockTokenModel) New(userID int64, ttl time.Duration, tokenType string, length int) (*Token, error) {
	token := getMockToken(userID, ttl, tokenType)

	return token, nil
}

func (m MockTokenModel) Insert(token *Token) error {
	return nil
}

func (m MockTokenModel) MarkUsed(tokenId int64) error {
	return nil
}
