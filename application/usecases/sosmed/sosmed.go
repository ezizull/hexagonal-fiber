package sosmed

import (
	sosmedDomain "hacktiv/final-project/domain/sosmed"

	sosmedRepository "hacktiv/final-project/infrastructure/repository/postgres/sosmed"
)

// Service is a struct that contains the repository implementation for sosmed use case
type Service struct {
	SocialMediaTesting    sosmedRepository.SocialMediaTesting
	SocialMediaRepository sosmedRepository.Repository
}

// GetAll is a function that returns all sosmeds
func (s *Service) GetAll(page int64, limit int64) (*sosmedDomain.PaginationResultSocialMedia, error) {

	all, err := s.SocialMediaRepository.GetAll(page, limit)
	if err != nil {
		return nil, err
	}

	return &sosmedDomain.PaginationResultSocialMedia{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

// UserGetAll is a function that returns all sosmeds
func (s *Service) UserGetAll(userId int, page int64, limit int64) (*sosmedDomain.PaginationResultSocialMedia, error) {

	all, err := s.SocialMediaRepository.UserGetAll(userId, page, limit)
	if err != nil {
		return nil, err
	}

	return &sosmedDomain.PaginationResultSocialMedia{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

// GetByID is a function that returns a sosmed by id
func (s *Service) GetByID(id int) (*sosmedDomain.SocialMedia, error) {
	return s.SocialMediaRepository.GetByID(id)
}

// UserGetByID is a function that returns a sosmed by id
func (s *Service) UserGetByID(id int, userId int) (*sosmedDomain.SocialMedia, error) {
	return s.SocialMediaRepository.UserGetByID(id, userId)
}

// Create is a function that creates a sosmed
func (s *Service) Create(sosmed *sosmedDomain.NewSocialMedia) (*sosmedDomain.SocialMedia, error) {

	sosmedModel := sosmed.ToDomainMapper()

	return s.SocialMediaRepository.Create(sosmedModel)
}

// GetByMap is a function that returns a sosmed by map
func (s *Service) GetByMap(sosmedMap map[string]interface{}) (*sosmedDomain.SocialMedia, error) {
	return s.SocialMediaRepository.GetOneByMap(sosmedMap)
}

// Delete is a function that deletes a sosmed by id
func (s *Service) Delete(id int) (err error) {
	return s.SocialMediaRepository.Delete(id)
}

// Update is a function that updates a sosmed by id
func (s *Service) Update(id int, updateSocialMedia sosmedDomain.UpdateSocialMedia) (*sosmedDomain.SocialMedia, error) {
	sosmed := updateSocialMedia.ToDomainMapper()
	return s.SocialMediaRepository.Update(id, &sosmed)
}

// Update is a function that updates a sosmed by id
func (s *Service) UserUpdate(id int, userId int, updateSocialMedia sosmedDomain.UpdateSocialMedia) (*sosmedDomain.SocialMedia, error) {
	sosmed := updateSocialMedia.ToDomainMapper()
	return s.SocialMediaRepository.UserUpdate(id, userId, &sosmed)
}
