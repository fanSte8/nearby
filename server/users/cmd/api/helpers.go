package main

import (
	"nearby/users/internal/data"
	"strconv"
	"time"

	"github.com/pascaldekloe/jwt"
)

func (app *application) generateJWT(user *data.User) (string, error) {
	var claims jwt.Claims
	claims.Subject = strconv.FormatInt(user.ID, 10)
	claims.Set = map[string]any{"activated": user.Activated}
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))

	bytes, err := claims.HMACSign(jwt.HS256, []byte(app.config.JWTSecret))
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
