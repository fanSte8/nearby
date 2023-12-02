package main

import (
	"nearby/common/clients"
	"nearby/notifications/internal/data"
	"net/url"
	"strconv"
	"sync"
)

type NotificationsWithUserData struct {
	Notification *data.NotificationResponse `json:"notifications"`
	User         *clients.User              `json:"user"`
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

func (app *application) combineNotificationsWithUserData(notifications []*data.NotificationResponse) []NotificationsWithUserData {
	type channelData struct {
		comment NotificationsWithUserData
		index   int
	}

	combinedNotifications := make([]NotificationsWithUserData, len(notifications))
	resultChannel := make(chan channelData, len(notifications))
	var wg sync.WaitGroup

	for i, comment := range notifications {
		wg.Add(1)

		go func(i int, notification *data.NotificationResponse) {
			defer wg.Done()

			var combinedNotification NotificationsWithUserData

			userData, err := app.usersClient.GetUserByID(notification.UserID)

			if err != nil || userData == nil {
				combinedNotification = NotificationsWithUserData{
					Notification: notification,
					User:         nil,
				}
			} else {
				combinedNotification = NotificationsWithUserData{
					Notification: notification,
					User:         &userData.User,
				}
			}

			resultChannel <- channelData{comment: combinedNotification, index: i}
		}(i, comment)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	for result := range resultChannel {
		combinedNotifications[result.index] = result.comment
	}

	return combinedNotifications
}
