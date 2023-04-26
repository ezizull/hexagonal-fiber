package adapter

import (
	userService "hacktiv/final-project/application/usecases/user"
	roleRepository "hacktiv/final-project/infrastructure/repository/postgres/role"
	userRepository "hacktiv/final-project/infrastructure/repository/postgres/user"
	userController "hacktiv/final-project/infrastructure/restapi/controllers/user"

	"gorm.io/gorm"
)

// UserAdapter is a function that returns a user controller
func UserAdapter(db *gorm.DB) *userController.Controller {
	uRepository := userRepository.Repository{DB: db}
	rRepository := roleRepository.Repository{DB: db}

	service := userService.Service{UserRepository: uRepository, RoleRepository: rRepository}
	return &userController.Controller{UserService: service}
}
