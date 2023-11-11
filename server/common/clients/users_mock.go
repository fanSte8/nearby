package clients

import "time"

type MockUsersClient struct {
}

func (m MockUsersClient) GetUserByID(id int64) (*UserData, error) {
	if id == 1 {
		return &UserData{
			User: User{
				ID:            1,
				FirstName:     "Test",
				LastName:      "Test",
				Email:         "test@mail.com",
				ImageUrl:      "",
				Activated:     true,
				PostsRadiusKm: 10,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			},
		}, nil
	}

	return nil, nil
}
