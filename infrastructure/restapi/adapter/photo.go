// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	photoService "hexagonal-fiber/application/usecases/photo"
	databsDomain "hexagonal-fiber/domain/database"
	photoRepository "hexagonal-fiber/infrastructure/repository/postgres/photo"
	photoController "hexagonal-fiber/infrastructure/restapi/controllers/photo"
)

// PhotoAdapter is a function that returns a photo controller
func PhotoAdapter(db databsDomain.Database) *photoController.Controller {
	pRepository := photoRepository.Repository{DB: db.Postgre}
	service := photoService.Service{PhotoRepository: pRepository}
	return &photoController.Controller{PhotoService: service}
}
