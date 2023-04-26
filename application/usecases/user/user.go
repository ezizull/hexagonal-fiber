// Package user provides the use case for user
package user

import (
	"errors"
	errorDomain "hacktiv/final-project/domain/errors"
	userDomain "hacktiv/final-project/domain/user"
	roleRepository "hacktiv/final-project/infrastructure/repository/postgres/role"
	userRepository "hacktiv/final-project/infrastructure/repository/postgres/user"

	"golang.org/x/crypto/bcrypt"
)

// Service is a struct that contains the repository implementation for user use case
type Service struct {
	UserRepository userRepository.Repository
	RoleRepository roleRepository.Repository
}

// GetAll is a function that returns all users
func (s *Service) GetAll() (*[]userDomain.User, error) {
	return s.UserRepository.GetAll()
}

// GetWithRole is a function that returns a user with role by id
func (s *Service) GetWithRole(id int) (responUserRole *userDomain.ResponseUserRole, err error) {
	userRole, err := s.UserRepository.GetWithRole(id)
	responUserRole = userRole.UserToResponseMapper()
	return
}

// GetByID is a function that returns a user by id
func (s *Service) GetByID(id int) (responUser *userDomain.ResponseUser, err error) {
	user, err := s.UserRepository.GetByID(id)
	responUser = user.DomainToResponseMapper()
	return
}

// Create is a function that creates a new user
func (s *Service) Create(newUser userDomain.NewUser) (*userDomain.User, error) {

	_, err := s.RoleRepository.GetByID(newUser.RoleID)
	if err != nil {
		return &userDomain.User{}, errorDomain.NewAppError(errors.New("role not found"), errorDomain.NotFound)
	}

	user := newUser.ToDomainMapper()

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return &userDomain.User{}, err
	}
	user.HashPassword = string(hash)

	return s.UserRepository.Create(user)
}

// GetOneByMap is a function that returns a user by map
func (s *Service) GetOneByMap(userMap map[string]interface{}) (*userDomain.User, error) {
	return s.UserRepository.GetOneByMap(userMap)
}

// Delete is a function that deletes a user by id
func (s *Service) Delete(id int) error {
	return s.UserRepository.Delete(id)
}

// Update is a function that updates a user by id
func (s *Service) Update(id int, updateUser userDomain.UpdateUser) (*userDomain.User, error) {
	user := updateUser.ToDomainMapper()
	return s.UserRepository.Update(id, &user)
}
