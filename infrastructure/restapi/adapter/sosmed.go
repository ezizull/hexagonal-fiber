package adapter

import (
	sosmedService "hacktiv/final-project/application/usecases/sosmed"
	sosmedRepository "hacktiv/final-project/infrastructure/repository/postgres/sosmed"
	sosmedController "hacktiv/final-project/infrastructure/restapi/controllers/sosmed"

	"gorm.io/gorm"
)

// SocialMediaAdapter is a function that returns a sosmed controller
func SocialMediaAdapter(db *gorm.DB) *sosmedController.Controller {
	mRepository := sosmedRepository.Repository{DB: db}
	service := sosmedService.Service{SocialMediaRepository: mRepository}
	return &sosmedController.Controller{SocialMediaService: service}
}
