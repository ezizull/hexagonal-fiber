package adapter

import (
	sosmedService "hexagonal-fiber/application/usecases/sosmed"
	databsDomain "hexagonal-fiber/domain/database"
	sosmedRepository "hexagonal-fiber/infrastructure/repository/postgres/sosmed"
	sosmedController "hexagonal-fiber/infrastructure/restapi/controllers/sosmed"
)

// SocialMediaAdapter is a function that returns a sosmed controller
func SocialMediaAdapter(db databsDomain.Database) *sosmedController.Controller {
	sRepository := sosmedRepository.Repository{DB: db.Postgre}
	service := sosmedService.Service{SocialMediaRepository: sRepository}
	return &sosmedController.Controller{SocialMediaService: service}
}
