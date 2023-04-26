package photo

import (
	photoDomain "hexagonal-fiber/domain/photo"

	photoRepository "hexagonal-fiber/infrastructure/repository/postgres/photo"
)

// Service is a struct that contains the repository implementation for photo use case
type Service struct {
	PhotoTesting    photoRepository.PhotoTesting
	PhotoRepository photoRepository.Repository
}

// GetAll is a function that returns all photos
func (s *Service) GetAll(page int, limit int) (*photoDomain.PaginationPhoto, error) {

	all, err := s.PhotoRepository.GetAll(page, limit)
	if err != nil {
		return nil, err
	}

	return &photoDomain.PaginationPhoto{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

// UserGetAll is a function that returns all photos
func (s *Service) UserGetAll(userId int, page int, limit int) (*photoDomain.PaginationPhoto, error) {

	all, err := s.PhotoRepository.UserGetAll(userId, page, limit)
	if err != nil {
		return nil, err
	}

	return &photoDomain.PaginationPhoto{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

// GetWithComments is a function that returns a photo by id
func (s *Service) GetWithComments(id int, page int, limit int) (*photoDomain.ResponsePhotoComments, error) {
	return s.PhotoRepository.GetWithComments(id, page, limit)
}

// GetByID is a function that returns a photo with comments by id
func (s *Service) GetByID(id int) (*photoDomain.Photo, error) {
	return s.PhotoRepository.GetByID(id)
}

// UserGetByID is a function that returns a photo by id
func (s *Service) UserGetByID(id int, userId int) (*photoDomain.Photo, error) {
	return s.PhotoRepository.UserGetByID(id, userId)
}

// Create is a function that creates a photo
func (s *Service) Create(photo *photoDomain.NewPhoto) (*photoDomain.Photo, error) {

	photoModel := photo.ToDomainMapper()

	return s.PhotoRepository.Create(photoModel)
}

// GetByMap is a function that returns a photo by map
func (s *Service) GetByMap(photoMap map[string]interface{}) (*photoDomain.Photo, error) {
	return s.PhotoRepository.GetOneByMap(photoMap)
}

// Delete is a function that deletes a photo by id
func (s *Service) Delete(id int) (err error) {
	return s.PhotoRepository.Delete(id)
}

// Update is a function that updates a photo by id
func (s *Service) Update(id int, updatePhoto photoDomain.UpdatePhoto) (*photoDomain.Photo, error) {
	photo := updatePhoto.ToDomainMapper()
	return s.PhotoRepository.Update(id, &photo)
}

// Update is a function that updates a photo by id
func (s *Service) UserUpdate(id int, userId int, updatePhoto photoDomain.UpdatePhoto) (*photoDomain.Photo, error) {
	photo := updatePhoto.ToDomainMapper()
	return s.PhotoRepository.UserUpdate(id, userId, &photo)
}
