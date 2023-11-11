package main

import (
	"errors"
	"nearby/common/commoncontext"
	"nearby/common/jsonutils"
	"nearby/common/validator"
	"nearby/posts/internal/data"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (app *application) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	image, _, err := r.FormFile("image")
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	description := r.FormValue("description")
	longitude := r.FormValue("longitude")
	latitude := r.FormValue("latitude")

	imageKey := uuid.New().String()
	err = app.storage.Save(imageKey, image)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	userId := commoncontext.ContextGetUserID(r)

	post := &data.Post{
		UserID:      userId,
		Description: description,
		ImageUrl:    imageKey,
		Longitude:   longitude,
		Latitude:    latitude,
	}

	v := validator.New()
	data.ValidatePost(v, post)
	if !v.Valid() {
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Posts.Insert(post)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = jsonutils.WriteJSON(w, http.StatusCreated, envelope{"post": post}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) handleGetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")
	data.ValidateCoordinates(v, latitude, longitude)
	if !v.Valid() {
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	userId := commoncontext.ContextGetUserID(r)
	userData, err := app.usersClient.GetUserByID(userId)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	post, err := app.models.Posts.GetPost(id, userId, latitude, longitude)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.httpErrors.NotFoundResponse(w, r)
		default:
			app.httpErrors.ServerErrorResponse(w, r, err)
		}
		return
	}

	err = jsonutils.WriteJSON(w, http.StatusOK, envelope{"post": post, "user": userData.User}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) handleGetPosts(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	pagination := app.getPaginationFromQuery(queryValues)

	v := validator.New()

	latitude := queryValues.Get("latitude")
	longitude := queryValues.Get("longitude")
	data.ValidateCoordinates(v, latitude, longitude)
	if !v.Valid() {
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	sort := queryValues.Get("sort")

	userId := commoncontext.ContextGetUserID(r)
	userData, err := app.usersClient.GetUserByID(userId)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	posts, err := app.models.Posts.GetPosts(sort, userId, latitude, longitude, userData.User.PostsRadiusKm*1000, pagination)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	postsWithUsers := app.combinePostsWithUserData(posts)

	err = jsonutils.WriteJSON(w, http.StatusCreated, envelope{"posts": postsWithUsers}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) handleDeletePost(w http.ResponseWriter, r *http.Request) {
	userId := commoncontext.ContextGetUserID(r)

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	post, err := app.models.Posts.GetById(id)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	if post.UserID != userId {
		app.httpErrors.ForbiddenActionResponse(w, r)
		return
	}

	err = app.models.Posts.Delete(id)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = jsonutils.WriteJSON(w, http.StatusOK, envelope{"message": "Comment deleted"}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}
