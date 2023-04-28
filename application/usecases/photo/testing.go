package photo

import (
	photoDomain "hexagonal-fiber/domain/photo"
	photoRepository "hexagonal-fiber/infrastructure/repository/postgres/photo"
)

type PhotoTesting interface {
	GetAll(page int, limit int) (*photoDomain.PaginationPhoto, error)
	UserGetAll(userId string, page int, limit int) (*photoDomain.PaginationPhoto, error)
	GetWithComments(id string, page int, limit int) (*photoDomain.ResponsePhotoComments, error)
	GetByID(id string) (*photoDomain.Photo, error)
	UserGetByID(id string, userId string) (*photoDomain.Photo, error)
	Create(photo *photoDomain.NewPhoto) (*photoDomain.Photo, error)
	GetByMap(photoMap map[string]interface{}) (*photoDomain.Photo, error)
	Delete(id string) (err error)
	Update(id string, updatePhoto photoDomain.UpdatePhoto) (*photoDomain.Photo, error)
}

func NewTesting(photoTest photoRepository.PhotoTesting) PhotoTesting {
	return &Service{
		PhotoTesting: photoTest,
	}
}
