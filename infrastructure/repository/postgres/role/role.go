package role

import (
	domainRole "hexagonal-fiber/domain/user"

	mssgConst "hexagonal-fiber/utils/constant/message"

	"github.com/gofiber/fiber/v2"
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
			return nil, fiber.NewError(fiber.StatusNotFound, "role not found")
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
	}

	return &role, err
}
