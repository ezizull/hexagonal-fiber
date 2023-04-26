// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	authService "hacktiv/final-project/application/usecases/auth"
	userRepository "hacktiv/final-project/infrastructure/repository/postgres/user"
	authController "hacktiv/final-project/infrastructure/restapi/controllers/auth"

	"gorm.io/gorm"
)

// AuthAdapter is a function that returns a auth controller
func AuthAdapter(db *gorm.DB) *authController.Controller {
	uRepository := userRepository.Repository{DB: db}
	service := authService.Service{UserRepository: uRepository}
	return &authController.Controller{AuthService: service}
}
