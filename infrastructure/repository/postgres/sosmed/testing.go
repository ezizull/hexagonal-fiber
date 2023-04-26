package sosmed

import sosmedDomain "hacktiv/final-project/domain/sosmed"

type SocialMediaTesting interface {
	GetAll(page int64, limit int64) (*sosmedDomain.PaginationResultSocialMedia, error)
	UserGetAll(page int64, userId int, limit int64) (*sosmedDomain.PaginationResultSocialMedia, error)
	Create(newSocialMedia *sosmedDomain.SocialMedia) (createdSocialMedia *sosmedDomain.SocialMedia, err error)
	GetByID(id int) (*sosmedDomain.SocialMedia, error)
	UserGetByID(id int, userId int) (*sosmedDomain.SocialMedia, error)
	GetOneByMap(sosmedMap map[string]interface{}) (*sosmedDomain.SocialMedia, error)
	Update(id int, updateSocialMedia *sosmedDomain.SocialMedia) (*sosmedDomain.SocialMedia, error)
	Delete(id int) (err error)
}
