package sosmed

import sosmedDomain "hexagonal-fiber/domain/sosmed"

type SocialMediaTesting interface {
	GetAll(page int, limit int) (*sosmedDomain.PaginationSocialMedia, error)
	UserGetAll(page int, userId string, limit int) (*sosmedDomain.PaginationSocialMedia, error)
	Create(newSocialMedia *sosmedDomain.SocialMedia) (createdSocialMedia *sosmedDomain.SocialMedia, err error)
	GetByID(id string) (*sosmedDomain.SocialMedia, error)
	UserGetByID(id string, userId string) (*sosmedDomain.SocialMedia, error)
	GetOneByMap(sosmedMap map[string]interface{}) (*sosmedDomain.SocialMedia, error)
	Update(id string, updateSocialMedia *sosmedDomain.SocialMedia) (*sosmedDomain.SocialMedia, error)
	Delete(id string) (err error)
}
