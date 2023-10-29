package storage

import "mime/multipart"

type Storage interface {
	Save(key string, file multipart.File) error
	GetURL(key string) (string, error)
}
