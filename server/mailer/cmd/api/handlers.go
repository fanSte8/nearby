package main

import (
	"nearby/common/jsonutils"
	"nearby/common/validator"
	"net/http"
)

func (app *application) handleActivationTokenMail(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Token     string `json:"token"`
		Recipient string `json:"recipient"`
	}

	err := jsonutils.ReadJSON(w, r, &input)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	v.Check(input.Token != "", "token", "Password token must be provided")
	v.Check(input.Recipient != "", "recipient", "Message recipient must not be empty")
	v.Check(validator.Matches(input.Recipient, validator.EmailRX), "recipient", "Recipient must be a valid email address")

	if !v.Valid() {
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.mailer.Send(input.Recipient, "Nearby account activation token", "Your nearby account activation token is "+input.Token)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = jsonutils.WriteJSON(w, http.StatusOK, map[string]any{"message": "Email sent"}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) handlePasswordResetTokenMail(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Token     string `json:"token"`
		Recipient string `json:"recipient"`
	}

	err := jsonutils.ReadJSON(w, r, &input)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	v.Check(input.Token != "", "token", "Password reset token must be provided")
	v.Check(input.Recipient != "", "recipient", "Recipient must not be empty")
	v.Check(validator.Matches(input.Recipient, validator.EmailRX), "recipient", "Recipient must be valid email address")

	if !v.Valid() {
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.mailer.Send(input.Recipient, "Nearby password reset token", "Your nearby password reset token is "+input.Token)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = jsonutils.WriteJSON(w, http.StatusOK, map[string]any{"message": "Email sent"}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}
