// Package user contains the business logic for the user entity
package user

import (
	"encoding/json"

	errorDomain "hacktiv/final-project/domain/errors"
	userDomain "hacktiv/final-project/domain/user"

	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for user entity
type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all user data
func (r *Repository) GetAll() (*[]userDomain.User, error) {
	var users []userDomain.User
	err := r.DB.Find(&users).Error
	if err != nil {
		err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		return nil, err
	}

	return &users, err
}

// Create ... Insert New data
func (r *Repository) Create(newUser *userDomain.User) (*userDomain.User, error) {
	txDb := r.DB.Create(newUser)
	err := txDb.Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errorDomain.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &userDomain.User{}, err
		}
		switch newError.Number {
		case 1062:
			err = errorDomain.NewAppErrorWithType(errorDomain.ResourceAlreadyExists)
			return &userDomain.User{}, err

		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
	}
	return newUser, err
}

// GetOneByMap ... Fetch only one user by Map values
func (r *Repository) GetOneByMap(userMap map[string]interface{}) (*userDomain.User, error) {
	var userRepository userDomain.User

	tx := r.DB.Where(userMap).Limit(1).Find(&userRepository)
	if tx.Error != nil {
		err := errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		return &userDomain.User{}, err
	}
	return &userRepository, nil
}

// GetWithRoleByMap ... Fetch only one user with Role by Map values
func (r *Repository) GetWithRoleByMap(userMap map[string]interface{}) (*userDomain.UserRole, error) {
	var userRole userDomain.UserRole

	err := r.DB.Preload("Role").Where(userMap).First(&userRole).Error
	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
	}

	return &userRole, nil
}

// GetWithRole ... Fetch only one user with Role by ID
func (r *Repository) GetWithRole(id int) (*userDomain.UserRole, error) {
	var userRole userDomain.UserRole
	err := r.DB.Preload("Role").Where("id = ?", id).First(&userRole).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
	}

	return &userRole, err
}

// GetByID ... Fetch only one user by ID
func (r *Repository) GetByID(id int) (*userDomain.User, error) {
	var user userDomain.User
	err := r.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
	}

	return &user, err
}

// Update ... Update user
func (r *Repository) Update(id int, updateUser *userDomain.User) (*userDomain.User, error) {
	var user userDomain.User

	user.ID = id
	err := r.DB.Model(&user).
		Updates(updateUser).Error

	// err = config.DB.Save(user).Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errorDomain.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &userDomain.User{}, err
		}
		switch newError.Number {
		case 1062:
			err = errorDomain.NewAppErrorWithType(errorDomain.ResourceAlreadyExists)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
		return &userDomain.User{}, err

	}

	err = r.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		return &userDomain.User{}, err
	}

	return &user, err
}

// Delete ... Delete user
func (r *Repository) Delete(id int) (err error) {
	tx := r.DB.Delete(&userDomain.User{}, id)
	if tx.Error != nil {
		err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		return
	}

	if tx.RowsAffected == 0 {
		err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
	}

	return
}
