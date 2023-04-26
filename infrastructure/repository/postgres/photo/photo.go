package photo

import (
	"encoding/json"
	commentDomain "hacktiv/final-project/domain/comment"
	errorDomain "hacktiv/final-project/domain/errors"
	photoDomain "hacktiv/final-project/domain/photo"
	"log"

	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for photo entity
type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all photo data
func (r *Repository) GetAll(page int64, limit int64) (*photoDomain.PaginationResultPhoto, error) {
	var photos []photoDomain.Photo
	var total int64

	err := r.DB.Model(&photoDomain.Photo{}).Count(&total).Error
	if err != nil {
		return &photoDomain.PaginationResultPhoto{}, err
	}
	offset := (page - 1) * limit
	err = r.DB.Limit(int(limit)).Offset(int(offset)).Find(&photos).Error

	if err != nil {
		return &photoDomain.PaginationResultPhoto{}, err
	}

	numPages := (total + limit - 1) / limit
	var nextCursor, prevCursor uint
	if page < numPages {
		nextCursor = uint(page + 1)
	}
	if page > 1 {
		prevCursor = uint(page - 1)
	}

	return &photoDomain.PaginationResultPhoto{
		Data:       photoDomain.ArrayToDomainMapper(&photos),
		Total:      total,
		Limit:      limit,
		Current:    page,
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

// UserGetAll Fetch all photo data
func (r *Repository) UserGetAll(userId int, page int64, limit int64) (*photoDomain.PaginationResultPhoto, error) {
	var photos []photoDomain.Photo
	var total int64

	err := r.DB.Model(&photoDomain.Photo{}).Where("user_id = ?", userId).Count(&total).Error
	if err != nil {
		return &photoDomain.PaginationResultPhoto{}, err
	}

	offset := (page - 1) * limit

	err = r.DB.Limit(int(limit)).Offset(int(offset)).Find(&photos).Error
	if err != nil {
		return &photoDomain.PaginationResultPhoto{}, err
	}

	numPages := (total + limit - 1) / limit
	var nextCursor, prevCursor uint
	if page < numPages {
		nextCursor = uint(page + 1)
	}
	if page > 1 {
		prevCursor = uint(page - 1)
	}

	return &photoDomain.PaginationResultPhoto{
		Data:       photoDomain.ArrayToDomainMapper(&photos),
		Total:      total,
		Limit:      limit,
		Current:    page,
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

// GetWithComments ... Fetch a photo with comments by id
func (r *Repository) GetWithComments(id int, page int64, limit int64) (*photoDomain.ResponsePhotoComments, error) {
	var photoComments photoDomain.PhotoComment
	var total int64

	err := r.DB.Model(&commentDomain.Comment{}).Where("photo_id = ?", id).Count(&total).Error
	if err != nil {
		return &photoDomain.ResponsePhotoComments{}, err
	}

	offset := (page - 1) * limit
	photoComments.ID = id

	err = r.DB.Model(&photoDomain.Photo{}).Preload("Comment").Limit(int(limit)).Offset(int(offset)).First(&photoComments).Error
	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
		return &photoDomain.ResponsePhotoComments{}, err
	}

	numPages := (total + limit - 1) / limit
	var nextCursor, prevCursor uint
	if page < numPages {
		nextCursor = uint(page + 1)
	}
	if page > 1 {
		prevCursor = uint(page - 1)
	}

	comments := &commentDomain.PaginationResultComment{
		Data:       commentDomain.ArrayToDomainMapper(&photoComments.Comment),
		Total:      total,
		Limit:      limit,
		Current:    page,
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
			err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
		return &photoDomain.Photo{}, err
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
			err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
		return &photoDomain.Photo{}, err
	}

	return &photo, nil
}

// GetOneByMap ... Fetch only one photo by Map
func (r *Repository) GetOneByMap(photoMap map[string]interface{}) (*photoDomain.Photo, error) {
	var photo photoDomain.Photo

	err := r.DB.Where(photoMap).Limit(1).Find(&photo).Error
	if err != nil {
		err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		return nil, err
	}
	return &photo, err
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
			err = errorDomain.NewAppErrorWithType(errorDomain.ResourceAlreadyExists)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
		return
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

	// err = config.DB.Save(photo).Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errorDomain.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &photoDomain.Photo{}, err
		}
		switch newError.Number {
		case 1062:
			err = errorDomain.NewAppErrorWithType(errorDomain.ResourceAlreadyExists)
			return &photoDomain.Photo{}, err

		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
			return &photoDomain.Photo{}, err
		}
	}

	err = r.DB.Where("id = ?", id).First(&photo).Error
	if err != nil {
		err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		return &photoDomain.Photo{}, err
	}

	return &photo, err
}

// UserUpdate ... UserUpdate photo
func (r *Repository) UserUpdate(id int, userId int, updatePhoto *photoDomain.Photo) (*photoDomain.Photo, error) {
	var photo photoDomain.Photo

	photo.ID = id
	photo.UserID = userId
	err := r.DB.Model(&photo).
		Updates(updatePhoto).Error

	// err = config.DB.Save(photo).Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errorDomain.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &photoDomain.Photo{}, err
		}
		switch newError.Number {
		case 1062:
			err = errorDomain.NewAppErrorWithType(errorDomain.ResourceAlreadyExists)
			return &photoDomain.Photo{}, err

		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
			return &photoDomain.Photo{}, err
		}
	}

	err = r.DB.Where("id = ?", id).First(&photo).Error
	if err != nil {
		err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		return &photoDomain.Photo{}, err
	}

	return &photo, err
}

// Delete ... Delete photo
func (r *Repository) Delete(id int) (err error) {
	tx := r.DB.Delete(&photoDomain.Photo{}, id)

	log.Println("check ", tx)
	if tx.Error != nil {
		err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		return
	}

	if tx.RowsAffected == 0 {
		err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
	}

	return
}
