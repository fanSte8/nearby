package main

import (
	"nearby/common/commoncontext"
	"nearby/common/jsonutils"
	"nearby/common/validator"
	"nearby/notifications/internal/data"
	"net/http"
)

type envelope map[string]any

func (app *application) handleGetNotifications(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	pagination := app.getPaginationFromQuery(queryValues)

	userId := commoncontext.ContextGetUserID(r)

	notifications, err := app.models.Notification.GetList(userId, pagination)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	notificationsWithUserData := app.combineNotificationsWithUserData(notifications)

	err = jsonutils.WriteJSON(w, http.StatusOK, envelope{"notifications": notificationsWithUserData}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) handleCreateNotification(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FromUserID int64  `json:"fromUserId"`
		ToUserID   int64  `json:"toUserId"`
		PostID     int64  `json:"postId"`
		Type       string `json:"type"`
	}

	err := jsonutils.ReadJSON(w, r, &input)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	notification := data.Notification{
		FromUserID: input.FromUserID,
		ToUserID:   input.ToUserID,
		PostID:     input.PostID,
		Type:       input.Type,
	}

	v := validator.New()
	data.ValidateNotification(v, &notification)
	if !v.Valid() {
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Notification.Insert(&notification)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = jsonutils.WriteJSON(w, http.StatusCreated, envelope{"notification": notification}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}
