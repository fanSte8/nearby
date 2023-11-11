package data

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"math/rand"
	"time"
)

const (
	ActivationToken    = "ActivationToken"
	PasswordResetToken = "PasswordResetToken"
)

type Token struct {
	Text   string    `json:"token"`
	Hash   []byte    `json:"-"`
	UserID int64     `json:"-"`
	Expiry time.Time `json:"expiry"`
	Type   string    `json:"-"`
}

func generateToken(userId int64, ttl time.Duration, tokenType string, length int) *Token {
	token := &Token{
		UserID: userId,
		Expiry: time.Now().Add(ttl),
		Type:   tokenType,
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomBytes := make([]byte, length)

	for i := range randomBytes {
		randomBytes[i] = charset[seededRand.Intn(len(charset))]
	}

	token.Text = string(randomBytes)
	hash := sha256.Sum256([]byte(token.Text))
	token.Hash = hash[:]

	return token
}

type ITokenModel interface {
	New(userID int64, ttl time.Duration, tokenType string, length int) (*Token, error)
	Insert(token *Token) error
	MarkUsed(tokenId int64) error
}

type TokenModel struct {
	db *sql.DB
}

func (m TokenModel) New(userID int64, ttl time.Duration, tokenType string, length int) (*Token, error) {
	token := generateToken(userID, ttl, tokenType, length)

	err := m.Insert(token)
	if err != nil {
		return nil, err
	}

	return token, err
}

func (m TokenModel) Insert(token *Token) error {
	query := `
        INSERT INTO tokens (hash, user_id, expiry, type) 
        VALUES ($1, $2, $3, $4)`

	args := []any{token.Hash, token.UserID, token.Expiry, token.Type}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.db.ExecContext(ctx, query, args...)
	return err
}

func (m TokenModel) MarkUsed(tokenId int64) error {
	query := `UPDATE tokens SET used=true WHERE id=$1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.db.ExecContext(ctx, query, tokenId)
	if err != nil {
		return err
	}

	return nil
}
