package sosmed

import (
	sosmedDomain "hacktiv/final-project/domain/sosmed"
	sosmedRepository "hacktiv/final-project/infrastructure/repository/postgres/sosmed"
)

type SocialMediaTesting interface {
	GetAll(page int64, limit int64) (*sosmedDomain.PaginationResultSocialMedia, error)
	UserGetAll(userId int, page int64, limit int64) (*sosmedDomain.PaginationResultSocialMedia, error)
	GetByID(id int) (*sosmedDomain.SocialMedia, error)
	UserGetByID(id int, userId int) (*sosmedDomain.SocialMedia, error)
	Create(sosmed *sosmedDomain.NewSocialMedia) (*sosmedDomain.SocialMedia, error)
	GetByMap(sosmedMap map[string]interface{}) (*sosmedDomain.SocialMedia, error)
	Delete(id int) (err error)
	Update(id int, updateSocialMedia sosmedDomain.UpdateSocialMedia) (*sosmedDomain.SocialMedia, error)
}

func NewTesting(sosmedTest sosmedRepository.SocialMediaTesting) SocialMediaTesting {
	return &Service{
		SocialMediaTesting: sosmedTest,
	}
}
