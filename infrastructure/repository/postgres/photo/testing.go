package photo

import (
	photoDomain "hexagonal-fiber/domain/photo"
)

type PhotoTesting interface {
	GetAll(page int, limit int) (*photoDomain.PaginationPhoto, error)
	UserGetAll(userId int, page int, limit int) (*photoDomain.PaginationPhoto, error)
	Create(newPhoto *photoDomain.Photo) (createdPhoto *photoDomain.Photo, err error)
	GetWithComments(id int, page int, limit int) (*photoDomain.ResponsePhotoComments, error)
	GetByID(id int) (*photoDomain.Photo, error)
	UserGetByID(id int, userId int) (*photoDomain.Photo, error)
	GetOneByMap(photoMap map[string]interface{}) (*photoDomain.Photo, error)
	Update(id int, updatePhoto *photoDomain.Photo) (*photoDomain.Photo, error)
	Delete(id int) (err error)
}
