package role

import (
	errorDomain "hacktiv/final-project/domain/errors"
	domainRole "hacktiv/final-project/domain/user"

	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for user entity
type Repository struct {
	DB *gorm.DB
}

// GetByID ... Fetch only one role by ID
func (r *Repository) GetByID(id string) (*domainRole.Role, error) {
	var role domainRole.Role
	err := r.DB.Where("id = ?", id).First(&role).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
	}

	return &role, err
}
