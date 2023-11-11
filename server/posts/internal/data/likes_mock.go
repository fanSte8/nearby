package data

type MockLikeModel struct {
}

func (m MockLikeModel) Exists(userId, postId int64) (bool, error) {
	return userId == 1 && postId == 1, nil
}

func (m MockLikeModel) Insert(userId, postId int64) error {
	return nil
}

func (m MockLikeModel) Delete(userId, postId int64) error {
	return nil
}
