package main

import (
	"nearby/common/commoncontext"
	"nearby/common/jsonutils"
	"nearby/common/validator"
	"nearby/posts/internal/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *application) handleCreateComment(w http.ResponseWriter, r *http.Request) {
	userId := commoncontext.ContextGetUserID(r)

	params := mux.Vars(r)
	postId, err := strconv.ParseInt(params["postId"], 10, 64)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	var input struct {
		Text string `json:"text"`
	}

	err = jsonutils.ReadJSON(w, r, &input)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	v.Check(input.Text != "", "text", "field is required")
	if !v.Valid() {
		app.httpErrors.FailedValidationResponse(w, r, v.Errors)
		return
	}

	comment := &data.Comment{
		UserID: userId,
		PostID: postId,
		Text:   input.Text,
	}

	err = app.models.Comments.Insert(comment)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	err = jsonutils.WriteJSON(w, http.StatusOK, envelope{"comment": comment}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) handleDeleteComment(w http.ResponseWriter, r *http.Request) {
	userId := commoncontext.ContextGetUserID(r)

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	comment, err := app.models.Comments.GetById(id)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	if comment.UserID != userId {
		app.httpErrors.ForbiddenActionResponse(w, r)
		return
	}

	err = app.models.Comments.Delete(id)
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

func (app *application) handleGetComments(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	pagination := app.getPaginationFromQuery(queryValues)

	params := mux.Vars(r)
	postId, err := strconv.ParseInt(params["postId"], 10, 64)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	comments, err := app.models.Comments.GetList(postId, pagination)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}

	commentsWithUsers := app.combineCommentsWithUserData(comments)
	err = jsonutils.WriteJSON(w, http.StatusCreated, envelope{"comments": commentsWithUsers}, nil)
	if err != nil {
		app.httpErrors.ServerErrorResponse(w, r, err)
		return
	}
}
