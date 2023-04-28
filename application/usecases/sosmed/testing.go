package sosmed

import (
	sosmedDomain "hexagonal-fiber/domain/sosmed"
	sosmedRepository "hexagonal-fiber/infrastructure/repository/postgres/sosmed"
)

type SocialMediaTesting interface {
	GetAll(page int, limit int) (*sosmedDomain.PaginationSocialMedia, error)
	UserGetAll(userId string, page int, limit int) (*sosmedDomain.PaginationSocialMedia, error)
	GetByID(id string) (*sosmedDomain.SocialMedia, error)
	UserGetByID(id string, userId string) (*sosmedDomain.SocialMedia, error)
	Create(sosmed *sosmedDomain.NewSocialMedia) (*sosmedDomain.SocialMedia, error)
	GetByMap(sosmedMap map[string]interface{}) (*sosmedDomain.SocialMedia, error)
	Delete(id string) (err error)
	Update(id string, updateSocialMedia sosmedDomain.UpdateSocialMedia) (*sosmedDomain.SocialMedia, error)
}

func NewTesting(sosmedTest sosmedRepository.SocialMediaTesting) SocialMediaTesting {
	return &Service{
		SocialMediaTesting: sosmedTest,
	}
}
