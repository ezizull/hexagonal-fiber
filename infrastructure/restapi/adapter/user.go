package adapter

import (
	databsDomain "hexagonal-fiber/domain/database"

	userService "hexagonal-fiber/application/usecases/user"
	roleRepository "hexagonal-fiber/infrastructure/repository/postgres/role"
	userRepository "hexagonal-fiber/infrastructure/repository/postgres/user"
	userController "hexagonal-fiber/infrastructure/restapi/controllers/user"
)

// UserAdapter is a function that returns a user controller
func UserAdapter(db databsDomain.Database) *userController.Controller {
	uRepository := userRepository.Repository{DB: db.Postgre}
	rRepository := roleRepository.Repository{DB: db.Postgre}

	service := userService.Service{UserRepository: uRepository, RoleRepository: rRepository}
	return &userController.Controller{UserService: service}
}
