// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	authService "hexagonal-fiber/application/usecases/auth"

	databsDomain "hexagonal-fiber/domain/database"

	roleRepository "hexagonal-fiber/infrastructure/repository/postgres/role"
	userRepository "hexagonal-fiber/infrastructure/repository/postgres/user"
	authController "hexagonal-fiber/infrastructure/restapi/controllers/auth"
)

// AuthAdapter is a function that returns a auth controller
func AuthAdapter(db databsDomain.Database) *authController.Controller {
	uRepository := userRepository.Repository{DB: db.Postgre}
	rRepository := roleRepository.Repository{DB: db.Postgre}

	service := authService.Service{
		UserRepository: uRepository,
		RoleRepository: rRepository,
	}

	return &authController.Controller{
		InfoRedis:   db.Redis,
		AuthService: service,
	}
}
