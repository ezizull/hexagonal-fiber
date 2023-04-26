package adapter

import (
	sosmedService "hexagonal-fiber/application/usecases/sosmed"
	sosmedRepository "hexagonal-fiber/infrastructure/repository/postgres/sosmed"
	sosmedController "hexagonal-fiber/infrastructure/restapi/controllers/sosmed"

	"gorm.io/gorm"
)

// SocialMediaAdapter is a function that returns a sosmed controller
func SocialMediaAdapter(db *gorm.DB) *sosmedController.Controller {
	mRepository := sosmedRepository.Repository{DB: db}
	service := sosmedService.Service{SocialMediaRepository: mRepository}
	return &sosmedController.Controller{SocialMediaService: service}
}
