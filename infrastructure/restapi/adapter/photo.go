// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	photoService "hexagonal-fiber/application/usecases/photo"
	photoRepository "hexagonal-fiber/infrastructure/repository/postgres/photo"
	photoController "hexagonal-fiber/infrastructure/restapi/controllers/photo"

	"gorm.io/gorm"
)

// PhotoAdapter is a function that returns a photo controller
func PhotoAdapter(db *gorm.DB) *photoController.Controller {
	mRepository := photoRepository.Repository{DB: db}
	service := photoService.Service{PhotoRepository: mRepository}
	return &photoController.Controller{PhotoService: service}
}
