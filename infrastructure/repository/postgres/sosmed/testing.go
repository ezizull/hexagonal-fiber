package sosmed

import sosmedDomain "hexagonal-fiber/domain/sosmed"

type SocialMediaTesting interface {
	GetAll(page int, limit int) (*sosmedDomain.PaginationSocialMedia, error)
	UserGetAll(page int, userId int, limit int) (*sosmedDomain.PaginationSocialMedia, error)
	Create(newSocialMedia *sosmedDomain.SocialMedia) (createdSocialMedia *sosmedDomain.SocialMedia, err error)
	GetByID(id int) (*sosmedDomain.SocialMedia, error)
	UserGetByID(id int, userId int) (*sosmedDomain.SocialMedia, error)
	GetOneByMap(sosmedMap map[string]interface{}) (*sosmedDomain.SocialMedia, error)
	Update(id int, updateSocialMedia *sosmedDomain.SocialMedia) (*sosmedDomain.SocialMedia, error)
	Delete(id int) (err error)
}
