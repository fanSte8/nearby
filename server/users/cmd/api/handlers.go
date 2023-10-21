package main

import (
	"fmt"
	"nearby/common/jsonutils"
	"nearby/common/validator"
	"nearby/users/internal/data"
	"net/http"
	"strconv"
	"time"

	"github.com/pascaldekloe/jwt"
)

type envelope = map[string]any

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	err := jsonutils.ReadJSON(w, r, &input)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	checkUser, err := app.models.Users.GetByEmail(input.Email)

	if err != nil && err != data.ErrRecordNotFound {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	if checkUser != nil {
		v.AddError("email", "Email already exists")
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	user := &data.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}

	user.Password.Set(input.Password)

	data.ValidateUser(v, user)
	if !v.Valid() {
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Users.Insert(user)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	fmt.Println(user)

	err = jsonutils.WriteJSON(w, http.StatusCreated, envelope{"user": user}, http.Header{})
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := jsonutils.ReadJSON(w, r, &input)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	v.Check(input.Email != "", "email", "Field is required")
	v.Check(input.Password != "", "password", "Field is required")
	if !v.Valid() {
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	user, err := app.models.Users.GetByEmail(input.Email)
	if err != nil {
		if err == data.ErrRecordNotFound {
			app.httpErrors.InvalidCredentialsResponse(w, r)
			return
		} else {
			app.httpErrors.ServerErrorResponse(w, r, err)
			return
		}
	}

	match, err := user.Password.Matches(input.Password)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	if !match {
		app.httpErrors.InvalidCredentialsResponse(w, r)
		return
	}

	var claims jwt.Claims
	claims.Subject = strconv.FormatInt(user.ID, 10)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(app.config.JWTSecret))
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = jsonutils.WriteJSON(w, http.StatusCreated, envelope{"token": string(jwtBytes)}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}
