package photo

import (
	"encoding/json"
	commentDomain "hexagonal-fiber/domain/comment"
	errorDomain "hexagonal-fiber/domain/error"
	photoDomain "hexagonal-fiber/domain/photo"

	mssgConst "hexagonal-fiber/utils/constant/message"

	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for photo entity
type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all photo data
func (r *Repository) GetAll(page int, limit int) (*photoDomain.PaginationPhoto, error) {
	var photos []photoDomain.Photo
	var total int64

	err := r.DB.Model(&photoDomain.Photo{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * limit
	if err = r.DB.Limit(limit).Offset(offset).Find(&photos).Error; err != nil {
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

	return &photoDomain.PaginationPhoto{
		Data:       photoDomain.ArrayToDomainMapper(&photos),
		Total:      total,
		Limit:      int64(limit),
		Current:    int64(page),
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

// UserGetAll Fetch all photo data
func (r *Repository) UserGetAll(userId int, page int, limit int) (*photoDomain.PaginationPhoto, error) {
	var photos []photoDomain.Photo
	var total int64

	err := r.DB.Model(&photoDomain.Photo{}).Where("user_id = ?", userId).Count(&total).Error
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * limit
	if err = r.DB.Limit(limit).Offset(offset).Find(&photos).Error; err != nil {
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

	return &photoDomain.PaginationPhoto{
		Data:       photoDomain.ArrayToDomainMapper(&photos),
		Total:      total,
		Limit:      int64(limit),
		Current:    int64(page),
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

// GetWithComments ... Fetch a photo with comments by id
func (r *Repository) GetWithComments(id int, page int, limit int) (*photoDomain.ResponsePhotoComments, error) {
	var photoComments photoDomain.PhotoComment
	var total int64

	err := r.DB.Model(&commentDomain.Comment{}).Where("photo_id = ?", id).Count(&total).Error
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * limit
	photoComments.ID = id

	err = r.DB.Model(&photoDomain.Photo{}).Preload("Comment").Limit(limit).Offset(offset).First(&photoComments).Error
	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
		case gorm.ErrRecordNotFound.Error():
			return nil, fiber.NewError(fiber.StatusNotFound, "photo with comment not found")
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
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

	comments := &commentDomain.PaginationComment{
		Data:       commentDomain.ArrayToDomainMapper(&photoComments.Comment),
		Total:      total,
		Limit:      int64(limit),
		Current:    int64(page),
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}

	return &photoDomain.ResponsePhotoComments{
		Photo:    photoComments.Photo,
		Comments: *comments,
	}, nil
}

// GetByID ... Fetch only one photo by Id
func (r *Repository) GetByID(id int) (*photoDomain.Photo, error) {
	var photo photoDomain.Photo
	err := r.DB.Where("id = ?", id).First(&photo).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			return nil, fiber.NewError(fiber.StatusNotFound, "photo not found")
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
	}

	return &photo, nil
}

// UserGetByID ... Fetch only one photo by Id
func (r *Repository) UserGetByID(id int, userId int) (*photoDomain.Photo, error) {
	var photo photoDomain.Photo
	err := r.DB.Where("id = ?", id).Where("user_id = ?", userId).First(&photo).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			return nil, fiber.NewError(fiber.StatusNotFound, "photo not found")
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
	}

	return &photo, nil
}

// GetOneByMap ... Fetch only one photo by Map
func (r *Repository) GetOneByMap(photoMap map[string]interface{}) (*photoDomain.Photo, error) {
	var photo photoDomain.Photo

	err := r.DB.Where(photoMap).Limit(1).Find(&photo).Error
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
	}
	return &photo, nil
}

// Create ... Insert New data
func (r *Repository) Create(newPhoto *photoDomain.Photo) (createdPhoto *photoDomain.Photo, err error) {
	tx := r.DB.Create(newPhoto)

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

	createdPhoto = newPhoto
	return
}

// Update ... Update photo
func (r *Repository) Update(id int, updatePhoto *photoDomain.Photo) (*photoDomain.Photo, error) {
	var photo photoDomain.Photo

	photo.ID = id
	err := r.DB.Model(&photo).
		Updates(updatePhoto).Error

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

	err = r.DB.Where("id = ?", id).First(&photo).Error
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "photo not found")
	}

	return &photo, nil
}

// UserUpdate ... UserUpdate photo
func (r *Repository) UserUpdate(id int, userId int, updatePhoto *photoDomain.Photo) (*photoDomain.Photo, error) {
	var photo photoDomain.Photo

	photo.ID = id
	photo.UserID = userId
	err := r.DB.Model(&photo).
		Updates(updatePhoto).Error

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

	err = r.DB.Where("id = ?", id).First(&photo).Error
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "photo not found")
	}

	return &photo, nil
}

// Delete ... Delete photo
func (r *Repository) Delete(id int) (err error) {
	tx := r.DB.Delete(&photoDomain.Photo{}, id)

	log.Println("check ", tx)
	if tx.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
	}

	if tx.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "photo not found")
	}

	return
}
