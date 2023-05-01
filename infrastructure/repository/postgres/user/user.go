// Package user contains the business logic for the user entity
package user

import (
	"encoding/json"

	errorDomain "hexagonal-fiber/domain/error"
	userDomain "hexagonal-fiber/domain/user"

	mssgConst "hexagonal-fiber/utils/constant/message"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

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
		return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
	}

	return &users, nil
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
			return nil, err
		}

		switch newError.Number {
		case 1062:
			return nil, fiber.NewError(fiber.StatusConflict, mssgConst.ResourceAlreadyExists)
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
	}
	return newUser, err
}

// GetOneByMap ... Fetch only one user by Map values
func (r *Repository) GetOneByMap(userMap map[string]interface{}) (*userDomain.User, error) {
	var userRepository userDomain.User

	tx := r.DB.Where(userMap).Limit(1).Find(&userRepository)
	if tx.Error != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
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
			return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
	}

	return &userRole, nil
}

// GetWithRole ... Fetch only one user with Role by ID
func (r *Repository) GetWithRole(id string) (*userDomain.UserRole, error) {
	var userRole userDomain.UserRole
	err := r.DB.Preload("Role").Where("id = ?", id).First(&userRole).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
	}

	return &userRole, err
}

// GetByID ... Fetch only one user by ID
func (r *Repository) GetByID(id string) (*userDomain.User, error) {
	var user userDomain.User
	err := r.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
	}

	return &user, err
}

// Update ... Update user
func (r *Repository) Update(id string, updateUser *userDomain.User) (*userDomain.User, error) {
	var user userDomain.User

	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "incorrect user id")
	}

	user.ID = uuid
	err = r.DB.Model(&user).
		Updates(updateUser).Error

	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errorDomain.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return nil, err
		}

		switch newError.Number {
		case 1062:
			return nil, fiber.NewError(fiber.StatusConflict, mssgConst.ResourceAlreadyExists)
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
	}

	err = r.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	return &user, err
}

// Delete ... Delete user
func (r *Repository) Delete(id string) (err error) {
	tx := r.DB.Where("id = ?", id).Delete(&userDomain.User{})
	if tx.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
	}

	if tx.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	return
}
