package photo

import (
	photoDomain "hexagonal-fiber/domain/photo"
)

type PhotoTesting interface {
	GetAll(page int, limit int) (*photoDomain.PaginationPhoto, error)
	UserGetAll(userId string, page int, limit int) (*photoDomain.PaginationPhoto, error)
	Create(newPhoto *photoDomain.Photo) (createdPhoto *photoDomain.Photo, err error)
	GetWithComments(id string, page int, limit int) (*photoDomain.ResponsePhotoComments, error)
	GetByID(id string) (*photoDomain.Photo, error)
	UserGetByID(id string, userId string) (*photoDomain.Photo, error)
	GetOneByMap(photoMap map[string]interface{}) (*photoDomain.Photo, error)
	Update(id string, updatePhoto *photoDomain.Photo) (*photoDomain.Photo, error)
	Delete(id string) (err error)
}
