package sosmed

import (
	"encoding/json"
	errorDomain "hacktiv/final-project/domain/errors"
	sosmedDomain "hacktiv/final-project/domain/sosmed"
	"log"

	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for sosmed entity
type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all sosmed data
func (r *Repository) GetAll(page int64, limit int64) (*sosmedDomain.PaginationResultSocialMedia, error) {
	var sosmeds []sosmedDomain.SocialMedia
	var total int64

	err := r.DB.Model(&sosmedDomain.SocialMedia{}).Count(&total).Error
	if err != nil {
		return &sosmedDomain.PaginationResultSocialMedia{}, err
	}
	offset := (page - 1) * limit
	err = r.DB.Limit(int(limit)).Offset(int(offset)).Find(&sosmeds).Error

	if err != nil {
		return &sosmedDomain.PaginationResultSocialMedia{}, err
	}

	numPages := (total + limit - 1) / limit
	var nextCursor, prevCursor uint
	if page < numPages {
		nextCursor = uint(page + 1)
	}
	if page > 1 {
		prevCursor = uint(page - 1)
	}

	return &sosmedDomain.PaginationResultSocialMedia{
		Data:       sosmedDomain.ArrayToDomainMapper(&sosmeds),
		Total:      total,
		Limit:      limit,
		Current:    page,
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

// UserGetAll Fetch all sosmed data
func (r *Repository) UserGetAll(userId int, page int64, limit int64) (*sosmedDomain.PaginationResultSocialMedia, error) {
	var sosmeds []sosmedDomain.SocialMedia
	var total int64

	err := r.DB.Model(&sosmedDomain.SocialMedia{}).Where("user_id = ?", userId).Count(&total).Error
	if err != nil {
		return &sosmedDomain.PaginationResultSocialMedia{}, err
	}
	offset := (page - 1) * limit
	err = r.DB.Limit(int(limit)).Offset(int(offset)).Find(&sosmeds).Error

	if err != nil {
		return &sosmedDomain.PaginationResultSocialMedia{}, err
	}

	numPages := (total + limit - 1) / limit
	var nextCursor, prevCursor uint
	if page < numPages {
		nextCursor = uint(page + 1)
	}
	if page > 1 {
		prevCursor = uint(page - 1)
	}

	return &sosmedDomain.PaginationResultSocialMedia{
		Data:       sosmedDomain.ArrayToDomainMapper(&sosmeds),
		Total:      total,
		Limit:      limit,
		Current:    page,
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

// GetByID ... Fetch only one sosmed by Id
func (r *Repository) GetByID(id int) (*sosmedDomain.SocialMedia, error) {
	var sosmed sosmedDomain.SocialMedia
	err := r.DB.Where("id = ?", id).First(&sosmed).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
		return &sosmedDomain.SocialMedia{}, err
	}

	return &sosmed, nil
}

// UserGetByID ... Fetch only one sosmed by Id
func (r *Repository) UserGetByID(id int, userId int) (*sosmedDomain.SocialMedia, error) {
	var sosmed sosmedDomain.SocialMedia
	err := r.DB.Where("id = ?", id).Where("user_id = ?", userId).First(&sosmed).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
		return &sosmedDomain.SocialMedia{}, err
	}

	return &sosmed, nil
}

// GetOneByMap ... Fetch only one sosmed by Map
func (r *Repository) GetOneByMap(sosmedMap map[string]interface{}) (*sosmedDomain.SocialMedia, error) {
	var sosmed sosmedDomain.SocialMedia

	err := r.DB.Where(sosmedMap).Limit(1).Find(&sosmed).Error
	if err != nil {
		err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		return nil, err
	}
	return &sosmed, err
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
			err = errorDomain.NewAppErrorWithType(errorDomain.ResourceAlreadyExists)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
		return
	}

	createdSocialMedia = newSocialMedia
	return
}

// Update ... Update sosmed
func (r *Repository) Update(id int, updateSocialMedia *sosmedDomain.SocialMedia) (*sosmedDomain.SocialMedia, error) {
	var sosmed sosmedDomain.SocialMedia

	sosmed.ID = id
	err := r.DB.Model(&sosmed).
		Updates(updateSocialMedia).Error

	// err = config.DB.Save(sosmed).Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errorDomain.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &sosmedDomain.SocialMedia{}, err
		}
		switch newError.Number {
		case 1062:
			err = errorDomain.NewAppErrorWithType(errorDomain.ResourceAlreadyExists)
			return &sosmedDomain.SocialMedia{}, err

		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
			return &sosmedDomain.SocialMedia{}, err
		}
	}

	err = r.DB.Where("id = ?", id).First(&sosmed).Error
	if err != nil {
		err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		return &sosmedDomain.SocialMedia{}, err
	}

	return &sosmed, err
}

// UserUpdate ... UserUpdate sosmed
func (r *Repository) UserUpdate(id int, userId int, updateSocialMedia *sosmedDomain.SocialMedia) (*sosmedDomain.SocialMedia, error) {
	var sosmed sosmedDomain.SocialMedia

	sosmed.ID = id
	sosmed.UserID = userId
	err := r.DB.Model(&sosmed).
		Updates(updateSocialMedia).Error

	// err = config.DB.Save(sosmed).Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errorDomain.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &sosmedDomain.SocialMedia{}, err
		}
		switch newError.Number {
		case 1062:
			err = errorDomain.NewAppErrorWithType(errorDomain.ResourceAlreadyExists)
			return &sosmedDomain.SocialMedia{}, err

		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
			return &sosmedDomain.SocialMedia{}, err
		}
	}

	err = r.DB.Where("id = ?", id).First(&sosmed).Error
	if err != nil {
		err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		return &sosmedDomain.SocialMedia{}, err
	}

	return &sosmed, err
}

// Delete ... Delete sosmed
func (r *Repository) Delete(id int) (err error) {
	tx := r.DB.Delete(&sosmedDomain.SocialMedia{}, id)

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
