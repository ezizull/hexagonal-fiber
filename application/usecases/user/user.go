// Package user provides the use case for user
package user

import (
	userDomain "hexagonal-fiber/domain/user"
	roleRepository "hexagonal-fiber/infrastructure/repository/postgres/role"
	userRepository "hexagonal-fiber/infrastructure/repository/postgres/user"
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
func (s *Service) GetWithRole(id string) (responUserRole *userDomain.ResponseUserRole, err error) {
	userRole, err := s.UserRepository.GetWithRole(id)
	if err != nil {
		return nil, err
	}

	responUserRole = userRole.UserToResponseMapper()
	return
}

// GetByID is a function that returns a user by id
func (s *Service) GetByID(id string) (responUser *userDomain.ResponseUser, err error) {
	user, err := s.UserRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	responUser = user.DomainToResponseMapper()
	return
}

// GetOneByMap is a function that returns a user by map
func (s *Service) GetOneByMap(userMap map[string]interface{}) (*userDomain.User, error) {
	return s.UserRepository.GetOneByMap(userMap)
}

// Delete is a function that deletes a user by id
func (s *Service) Delete(id string) error {
	return s.UserRepository.Delete(id)
}

// Update is a function that updates a user by id
func (s *Service) Update(id string, updateUser userDomain.UpdateUser) (*userDomain.User, error) {
	user := updateUser.ToDomainMapper()
	return s.UserRepository.Update(id, &user)
}
