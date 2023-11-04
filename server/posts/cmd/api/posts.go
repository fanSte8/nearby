package main

import (
	"nearby/common/commoncontext"
	"nearby/common/jsonutils"
	"nearby/common/validator"
	"nearby/posts/internal/data"
	"net/http"

	"github.com/google/uuid"
)

type envelope = map[string]any

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

func (app *application) handleGetLatestPosts(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	pagination := app.getPaginationFromQuery(queryValues)

	v := validator.New()

	radius := app.getRadiusFromQuery(queryValues, v)
	if !v.Valid() {
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	latitude, longitude := app.getCoordinatesFromQuery(queryValues, v)
	if !v.Valid() {
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	posts, err := app.models.Posts.GetLatest(latitude, longitude, radius, pagination)
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
