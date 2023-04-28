package sosmed

import (
	"encoding/json"
	errorDomain "hexagonal-fiber/domain/error"
	sosmedDomain "hexagonal-fiber/domain/sosmed"

	mssgConst "hexagonal-fiber/utils/constant/message"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for sosmed entity
type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all sosmed data
func (r *Repository) GetAll(page int, limit int) (*sosmedDomain.PaginationSocialMedia, error) {
	var sosmeds []sosmedDomain.SocialMedia
	var total int64

	err := r.DB.Model(&sosmedDomain.SocialMedia{}).Count(&total).Error
	if err != nil {
		return &sosmedDomain.PaginationSocialMedia{}, err
	}

	offset := (page - 1) * limit
	if err = r.DB.Limit(limit).Offset(offset).Find(&sosmeds).Error; err != nil {
		return nil, err
	}

	numPages := (total + int64(limit) - 1) / int64(limit)
	var nextCursor, prevCursor uint
	if page < int(numPages) {
		nextCursor = uint(page + 1)
	}
	if page > 1 {
		prevCursor = uint(page - 1)
	}

	return &sosmedDomain.PaginationSocialMedia{
		Data:       sosmedDomain.ArrayToDomainMapper(&sosmeds),
		Total:      total,
		Limit:      int64(limit),
		Current:    int64(page),
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

// UserGetAll Fetch all sosmed data
func (r *Repository) UserGetAll(userId string, page int, limit int) (*sosmedDomain.PaginationSocialMedia, error) {
	var sosmeds []sosmedDomain.SocialMedia
	var total int64

	err := r.DB.Model(&sosmedDomain.SocialMedia{}).Where("user_id = ?", userId).Count(&total).Error
	if err != nil {
		return &sosmedDomain.PaginationSocialMedia{}, err
	}

	offset := (page - 1) * limit
	if err = r.DB.Limit(limit).Offset(offset).Find(&sosmeds).Error; err != nil {
		return nil, err
	}

	numPages := (total + int64(limit) - 1) / int64(limit)
	var nextCursor, prevCursor uint
	if page < int(numPages) {
		nextCursor = uint(page + 1)
	}
	if page > 1 {
		prevCursor = uint(page - 1)
	}

	return &sosmedDomain.PaginationSocialMedia{
		Data:       sosmedDomain.ArrayToDomainMapper(&sosmeds),
		Total:      total,
		Limit:      int64(limit),
		Current:    int64(page),
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

// GetByID ... Fetch only one sosmed by Id
func (r *Repository) GetByID(id string) (*sosmedDomain.SocialMedia, error) {
	var sosmed sosmedDomain.SocialMedia
	err := r.DB.Where("id = ?", id).First(&sosmed).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			return nil, fiber.NewError(fiber.StatusNotFound, "social media not found")
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
	}

	return &sosmed, nil
}

// UserGetByID ... Fetch only one sosmed by Id
func (r *Repository) UserGetByID(id string, userId string) (*sosmedDomain.SocialMedia, error) {
	var sosmed sosmedDomain.SocialMedia
	err := r.DB.Where("id = ?", id).Where("user_id = ?", userId).First(&sosmed).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			return nil, fiber.NewError(fiber.StatusNotFound, "social media not found")
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
	}

	return &sosmed, nil
}

// GetOneByMap ... Fetch only one sosmed by Map
func (r *Repository) GetOneByMap(sosmedMap map[string]interface{}) (*sosmedDomain.SocialMedia, error) {
	var sosmed sosmedDomain.SocialMedia

	err := r.DB.Where(sosmedMap).Limit(1).Find(&sosmed).Error
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
	}

	return &sosmed, nil
}

// Create ... Insert New data
func (r *Repository) Create(newSocialMedia *sosmedDomain.SocialMedia) (createdSocialMedia *sosmedDomain.SocialMedia, err error) {
	tx := r.DB.Create(newSocialMedia)

	if tx.Error != nil {
		byteErr, _ := json.Marshal(tx.Error)
		var newError errorDomain.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return
		}

		switch newError.Number {
		case 1062:
			return nil, fiber.NewError(fiber.StatusConflict, mssgConst.ResourceAlreadyExists)
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
	}

	createdSocialMedia = newSocialMedia
	return
}

// Update ... Update sosmed
func (r *Repository) Update(id string, updateSocialMedia *sosmedDomain.SocialMedia) (*sosmedDomain.SocialMedia, error) {
	var sosmed sosmedDomain.SocialMedia

	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "incorrect sosmed id")
	}

	sosmed.ID = uuid
	err = r.DB.Model(&sosmed).
		Updates(updateSocialMedia).Error

	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errorDomain.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {

		}

		switch newError.Number {
		case 1062:
			return nil, fiber.NewError(fiber.StatusConflict, mssgConst.ResourceAlreadyExists)
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
	}

	err = r.DB.Where("id = ?", id).First(&sosmed).Error
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "social media not found")
	}

	return &sosmed, nil
}

// UserUpdate ... UserUpdate sosmed
func (r *Repository) UserUpdate(id string, userId string, updateSocialMedia *sosmedDomain.SocialMedia) (*sosmedDomain.SocialMedia, error) {
	var sosmed sosmedDomain.SocialMedia

	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "incorrect sosmed id")
	}

	sosmed.ID = uuid
	sosmed.UserID = userId
	err = r.DB.Model(&sosmed).
		Updates(updateSocialMedia).Error

	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errorDomain.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {

		}

		switch newError.Number {
		case 1062:
			return nil, fiber.NewError(fiber.StatusConflict, mssgConst.ResourceAlreadyExists)
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
	}

	err = r.DB.Where("id = ?", id).First(&sosmed).Error
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "social media not found")
	}

	return &sosmed, nil
}

// Delete ... Delete sosmed
func (r *Repository) Delete(id string) (err error) {
	tx := r.DB.Delete(&sosmedDomain.SocialMedia{}, id)

	log.Println("check ", tx)
	if tx.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
	}

	if tx.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "social media not found")
	}

	return
}
