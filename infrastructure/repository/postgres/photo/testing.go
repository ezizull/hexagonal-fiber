package photo

import (
	photoDomain "hacktiv/final-project/domain/photo"
)

type PhotoTesting interface {
	GetAll(page int64, limit int64) (*photoDomain.PaginationResultPhoto, error)
	UserGetAll(userId int, page int64, limit int64) (*photoDomain.PaginationResultPhoto, error)
	Create(newPhoto *photoDomain.Photo) (createdPhoto *photoDomain.Photo, err error)
	GetWithComments(id int, page int64, limit int64) (*photoDomain.ResponsePhotoComments, error)
	GetByID(id int) (*photoDomain.Photo, error)
	UserGetByID(id int, userId int) (*photoDomain.Photo, error)
	GetOneByMap(photoMap map[string]interface{}) (*photoDomain.Photo, error)
	Update(id int, updatePhoto *photoDomain.Photo) (*photoDomain.Photo, error)
	Delete(id int) (err error)
}
