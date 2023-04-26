package photo

import (
	photoDomain "hexagonal-fiber/domain/photo"
	photoRepository "hexagonal-fiber/infrastructure/repository/postgres/photo"
)

type PhotoTesting interface {
	GetAll(page int, limit int) (*photoDomain.PaginationPhoto, error)
	UserGetAll(userId int, page int, limit int) (*photoDomain.PaginationPhoto, error)
	GetWithComments(id int, page int, limit int) (*photoDomain.ResponsePhotoComments, error)
	GetByID(id int) (*photoDomain.Photo, error)
	UserGetByID(id int, userId int) (*photoDomain.Photo, error)
	Create(photo *photoDomain.NewPhoto) (*photoDomain.Photo, error)
	GetByMap(photoMap map[string]interface{}) (*photoDomain.Photo, error)
	Delete(id int) (err error)
	Update(id int, updatePhoto photoDomain.UpdatePhoto) (*photoDomain.Photo, error)
}

func NewTesting(photoTest photoRepository.PhotoTesting) PhotoTesting {
	return &Service{
		PhotoTesting: photoTest,
	}
}
