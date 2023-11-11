package main

import (
	"errors"
	"nearby/common/jsonutils"
	"nearby/common/validator"
	"nearby/users/internal/data"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pascaldekloe/jwt"
)

type envelope = map[string]any

func (app *application) handleCurrentUserData(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	if user.ImageUrl != "" {
		url, err := app.storage.GetURL(user.ImageUrl)
		if err != nil {
			app.logger.Error("Error getting user profile picture", "error", err)
		} else {
			user.ImageUrl = url
		}
	}

	err := jsonutils.WriteJSON(w, http.StatusCreated, envelope{"user": user}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	user, err := app.models.Users.GetById(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.httpErrors.NotFoundResponse(w, r)
		default:
			app.httpErrors.ServerErrorResponse(w, r, err)
		}
		return
	}

	if user.ImageUrl != "" {
		url, err := app.storage.GetURL(user.ImageUrl)
		if err != nil {
			app.logger.Error("Error getting user profile picture", "error", err)
		} else {
			user.ImageUrl = url
		}
	}

	err = jsonutils.WriteJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) handleRegisterUser(w http.ResponseWriter, r *http.Request) {
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

	if err != nil && !errors.Is(err, data.ErrRecordNotFound) {
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

	token, err := app.models.Tokens.New(user.ID, 10*time.Minute, data.ActivationToken, 8)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	go func() {
		err = app.mailerClient.SendActivationTokenMail(user.Email, token.Text)
		if err != nil {
			app.logger.Error("Error sending email", "error", err)
		}
	}()

	err = jsonutils.WriteJSON(w, http.StatusCreated, envelope{"user": user}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) handleLogin(w http.ResponseWriter, r *http.Request) {
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
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.httpErrors.InvalidCredentialsResponse(w, r)
		default:
			app.httpErrors.ServerErrorResponse(w, r, err)
		}
		return
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

	err = jsonutils.WriteJSON(w, http.StatusOK, envelope{"token": string(jwtBytes)}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) handleChangePassword(w http.ResponseWriter, r *http.Request) {
	var input struct {
		OldPassword string `json:"oldPassword"`
		Password    string `json:"password"`
	}

	err := jsonutils.ReadJSON(w, r, &input)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	user := app.contextGetUser(r)

	v := validator.New()

	match, err := user.Password.Matches(input.OldPassword)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	v.Check(match, "oldPassword", "Incorrect password")
	data.ValidatePassword(v, input.Password)

	if !v.Valid() {
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	user.Password.Set(input.Password)

	err = app.models.Users.Update(user)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = jsonutils.WriteJSON(w, http.StatusOK, envelope{"message": "Password updated"}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) handleNewActivationToken(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	token, err := app.models.Tokens.New(user.ID, 10*time.Minute, data.ActivationToken, 8)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
	go func() {
		err = app.mailerClient.SendActivationTokenMail(user.Email, token.Text)
		if err != nil {
			app.logger.Error("Error sending email", "error", err)
		}
	}()

	err = jsonutils.WriteJSON(w, http.StatusOK, envelope{"message": "New activation token sent"}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) handleActivateAccount(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Token string `json:"token"`
	}

	err := jsonutils.ReadJSON(w, r, &input)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	v.Check(input.Token != "", "token", "Must be provided")
	if !v.Valid() {
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	user := app.contextGetUser(r)

	tokenUser, tokenId, err := app.models.Users.GetByToken(data.ActivationToken, input.Token)
	if err != nil || user.ID != tokenUser.ID {
		switch {
		case errors.Is(err, data.ErrRecordNotFound) || user.ID != tokenUser.ID:
			v.AddError("token", "Incorrect or expired token")
			app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		default:
			app.httpErrors.ServerErrorResponse(w, r, err)
		}
		return
	}

	user.Activated = true

	err = app.models.Users.Update(user)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = app.models.Tokens.MarkUsed(tokenId)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = jsonutils.WriteJSON(w, http.StatusOK, envelope{"message": "Acount activated"}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) handleForgottenPassword(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
	}

	err := jsonutils.ReadJSON(w, r, &input)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	v := validator.New()
	data.ValidateEmail(v, input.Email)

	if !v.Valid() {
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	user, err := app.models.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.httpErrors.NotFoundResponse(w, r)
		default:
			app.httpErrors.ServerErrorResponse(w, r, err)
		}
		return
	}

	token, err := app.models.Tokens.New(user.ID, 10*time.Minute, data.PasswordResetToken, 8)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	go func() {
		err = app.mailerClient.SendPasswordResetTokenMail(user.Email, token.Text)
		if err != nil {
			app.logger.Error("Error sending email", "error", err)
		}
	}()

	err = jsonutils.WriteJSON(w, http.StatusOK, envelope{"message": "Password reset token sent"}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) handleResetPassword(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Token    string `json:"token"`
		Password string `json:"password"`
	}

	err := jsonutils.ReadJSON(w, r, &input)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	v := validator.New()
	data.ValidatePassword(v, input.Password)

	if !v.Valid() {
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	user, tokenId, err := app.models.Users.GetByToken(data.PasswordResetToken, input.Token)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			v.AddError("token", "Incorrect or expired token")
			app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		default:
			app.httpErrors.ServerErrorResponse(w, r, err)
		}
		return
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = app.models.Users.Update(user)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = app.models.Tokens.MarkUsed(tokenId)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = jsonutils.WriteJSON(w, http.StatusOK, envelope{"message": "Password upadted"}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) handleProfilePictureUpload(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	profilePictureKey := user.GetProfilePictureKey()

	f, _, err := r.FormFile("profile-picture")
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = app.storage.Save(profilePictureKey, f)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	user.ImageUrl = profilePictureKey
	err = app.models.Users.Update(user)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = jsonutils.WriteJSON(w, http.StatusOK, envelope{"message": "Profile picture set"}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) handlePostsRadiusUpdate(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Radius int `json:"radius"`
	}

	err := jsonutils.ReadJSON(w, r, &input)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	v := validator.New()
	v.Check(input.Radius > 0, "radius", "Value must be greater than 0")

	if !v.Valid() {
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	user := app.contextGetUser(r)

	user.PostsRadiusKm = input.Radius

	err = app.models.Users.Update(user)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = jsonutils.WriteJSON(w, http.StatusOK, envelope{"message": "Posts radius updated"}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}
