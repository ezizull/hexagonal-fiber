package adapter

import (
	userService "hexagonal-fiber/application/usecases/user"
	roleRepository "hexagonal-fiber/infrastructure/repository/postgres/role"
	userRepository "hexagonal-fiber/infrastructure/repository/postgres/user"
	userController "hexagonal-fiber/infrastructure/restapi/controllers/user"

	"gorm.io/gorm"
)

// UserAdapter is a function that returns a user controller
func UserAdapter(db *gorm.DB) *userController.Controller {
	uRepository := userRepository.Repository{DB: db}
	rRepository := roleRepository.Repository{DB: db}

	service := userService.Service{UserRepository: uRepository, RoleRepository: rRepository}
	return &userController.Controller{UserService: service}
}
