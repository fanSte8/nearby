package main

import (
	"nearby/common/clients"
	"nearby/common/validator"
	"nearby/posts/internal/data"
	"net/url"
	"strconv"
	"sync"
)

type PostWithUserData struct {
	Post *data.PostResponse `json:"post"`
	User *clients.User      `json:"user"`
}

func (app *application) getPaginationFromQuery(queryValues url.Values) data.Pagination {
	pageStr := queryValues.Get("page")
	pageSizeStr := queryValues.Get("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 20
	}

	return data.Pagination{Page: page, PageSize: pageSize}
}

func (app *application) getCoordinatesFromQuery(queryValues url.Values, v *validator.Validator) (string, string) {
	latitude := queryValues.Get("latitude")
	longitude := queryValues.Get("longitude")
	data.ValidateCoordinates(v, latitude, longitude)

	return latitude, longitude
}

func (app *application) combinePostsWithUserData(posts []*data.PostResponse) []PostWithUserData {
	type channelData struct {
		post  PostWithUserData
		index int
	}

	combinedPosts := make([]PostWithUserData, len(posts))
	resultChannel := make(chan channelData, len(posts))
	var wg sync.WaitGroup

	for i, post := range posts {
		wg.Add(1)

		go func(i int, post *data.PostResponse) {
			defer wg.Done()

			url, err := app.storage.GetURL(post.ImageUrl)
			if err != nil {
				app.logger.Error("Error getting user profile picture", "error", err)
			} else {
				post.ImageUrl = url
			}

			userData, err := app.usersClient.GetUserByID(post.UserID)
			if err != nil {
				return
			}

			combinedPost := PostWithUserData{
				Post: post,
				User: &userData.User,
			}

			resultChannel <- channelData{post: combinedPost, index: i}
		}(i, post)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	for result := range resultChannel {
		combinedPosts[result.index] = result.post
	}

	return combinedPosts
}
