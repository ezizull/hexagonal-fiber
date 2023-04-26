package comment

import (
	"encoding/json"
	commentDomain "hacktiv/final-project/domain/comment"
	errorDomain "hacktiv/final-project/domain/errors"
	"log"

	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for comment entity
type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all comment data
func (r *Repository) GetAll(page int64, limit int64) (*commentDomain.PaginationResultComment, error) {
	var comments []commentDomain.Comment
	var total int64

	err := r.DB.Model(&commentDomain.Comment{}).Count(&total).Error
	if err != nil {
		return &commentDomain.PaginationResultComment{}, err
	}
	offset := (page - 1) * limit
	err = r.DB.Limit(int(limit)).Offset(int(offset)).Find(&comments).Error

	if err != nil {
		return &commentDomain.PaginationResultComment{}, err
	}

	numPages := (total + limit - 1) / limit
	var nextCursor, prevCursor uint
	if page < numPages {
		nextCursor = uint(page + 1)
	}
	if page > 1 {
		prevCursor = uint(page - 1)
	}

	return &commentDomain.PaginationResultComment{
		Data:       commentDomain.ArrayToDomainMapper(&comments),
		Total:      total,
		Limit:      limit,
		Current:    page,
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

// UserGetAll Fetch all comment data
func (r *Repository) UserGetAll(userId int, page int64, limit int64) (*commentDomain.PaginationResultComment, error) {
	var comments []commentDomain.Comment
	var total int64

	err := r.DB.Model(&commentDomain.Comment{}).Where("user_id = ?", userId).Count(&total).Error
	if err != nil {
		return &commentDomain.PaginationResultComment{}, err
	}
	offset := (page - 1) * limit
	err = r.DB.Limit(int(limit)).Offset(int(offset)).Find(&comments).Error

	if err != nil {
		return &commentDomain.PaginationResultComment{}, err
	}

	numPages := (total + limit - 1) / limit
	var nextCursor, prevCursor uint
	if page < numPages {
		nextCursor = uint(page + 1)
	}
	if page > 1 {
		prevCursor = uint(page - 1)
	}

	return &commentDomain.PaginationResultComment{
		Data:       commentDomain.ArrayToDomainMapper(&comments),
		Total:      total,
		Limit:      limit,
		Current:    page,
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

// GetByID ... Fetch only one comment by Id
func (r *Repository) GetByID(id int) (*commentDomain.Comment, error) {
	var comment commentDomain.Comment
	err := r.DB.Where("id = ?", id).First(&comment).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
		return &commentDomain.Comment{}, err
	}

	return &comment, nil
}

// UserGetByID ... Fetch only one comment by Id
func (r *Repository) UserGetByID(id int, userId int) (*commentDomain.Comment, error) {
	var comment commentDomain.Comment
	err := r.DB.Where("id = ?", id).Where("user_id = ?", userId).First(&comment).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		}
		return &commentDomain.Comment{}, err
	}

	return &comment, nil
}

// GetOneByMap ... Fetch only one comment by Map
func (r *Repository) GetOneByMap(commentMap map[string]interface{}) (*commentDomain.Comment, error) {
	var comment commentDomain.Comment

	err := r.DB.Where(commentMap).Limit(1).Find(&comment).Error
	if err != nil {
		err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		return nil, err
	}
	return &comment, err
}

// Create ... Insert New data
func (r *Repository) Create(newComment *commentDomain.Comment) (createdComment *commentDomain.Comment, err error) {
	tx := r.DB.Create(newComment)

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

	createdComment = newComment
	return
}

// Update ... Update comment
func (r *Repository) Update(id int, updateComment *commentDomain.Comment) (*commentDomain.Comment, error) {
	var comment commentDomain.Comment

	comment.ID = id
	err := r.DB.Model(&comment).
		Updates(updateComment).Error

	// err = config.DB.Save(comment).Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errorDomain.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &commentDomain.Comment{}, err
		}
		switch newError.Number {
		case 1062:
			err = errorDomain.NewAppErrorWithType(errorDomain.ResourceAlreadyExists)
			return &commentDomain.Comment{}, err

		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
			return &commentDomain.Comment{}, err
		}
	}

	err = r.DB.Where("id = ?", id).First(&comment).Error
	if err != nil {
		err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		return &commentDomain.Comment{}, err
	}

	return &comment, err
}

// UserUpdate ... UserUpdate comment
func (r *Repository) UserUpdate(id int, userId int, updateComment *commentDomain.Comment) (*commentDomain.Comment, error) {
	var comment commentDomain.Comment

	comment.ID = id
	comment.UserID = userId
	err := r.DB.Model(&comment).
		Updates(updateComment).Error

	// err = config.DB.Save(comment).Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errorDomain.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &commentDomain.Comment{}, err
		}
		switch newError.Number {
		case 1062:
			err = errorDomain.NewAppErrorWithType(errorDomain.ResourceAlreadyExists)
			return &commentDomain.Comment{}, err

		default:
			err = errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
			return &commentDomain.Comment{}, err
		}
	}

	err = r.DB.Where("id = ?", id).First(&comment).Error
	if err != nil {
		err = errorDomain.NewAppErrorWithType(errorDomain.NotFound)
		return &commentDomain.Comment{}, err
	}

	return &comment, err
}

// Delete ... Delete comment
func (r *Repository) Delete(id int) (err error) {
	tx := r.DB.Delete(&commentDomain.Comment{}, id)

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
