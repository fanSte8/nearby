package main

import (
	"errors"
	"nearby/common/clients"
	"nearby/common/commoncontext"
	"nearby/common/jsonutils"
	"nearby/posts/internal/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *application) handlePostLike(w http.ResponseWriter, r *http.Request) {
	userId := commoncontext.ContextGetUserID(r)

	params := mux.Vars(r)
	postId, err := strconv.ParseInt(params["postId"], 10, 64)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	post, err := app.models.Posts.GetById(postId)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.httpErrors.NotFoundResponse(w, r)
		default:
			app.httpErrors.ServerErrorResponse(w, r, err)
		}
		return
	}

	exists, err := app.models.Likes.Exists(userId, postId)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	if exists {
		err = app.models.Likes.Delete(userId, postId)
	} else {
		err = app.models.Likes.Insert(userId, postId)
		go func() {
			app.notificationsClient.CreateNotification(clients.CreateNotificationInput{
				FromUserID: userId,
				ToUserID:   post.UserID,
				PostID:     postId,
				Type:       clients.LikeNotificationType,
			})
		}()
	}
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = jsonutils.WriteJSON(w, http.StatusOK, envelope{"like": !exists}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}
