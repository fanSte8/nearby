package storage

import "mime/multipart"

type MockStorage struct {
}

func (s MockStorage) Save(key string, file multipart.File) error {
	return nil
}

func (s MockStorage) GetURL(key string) (string, error) {
	return "", nil
}
