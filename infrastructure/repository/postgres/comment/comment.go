package comment

import (
	"encoding/json"
	commentDomain "hexagonal-fiber/domain/comment"
	errorDomain "hexagonal-fiber/domain/error"
	"log"

	mssgConst "hexagonal-fiber/utils/constant/message"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for comment entity
type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all comment data
func (r *Repository) GetAll(page int, limit int) (*commentDomain.PaginationComment, error) {
	var comments []commentDomain.Comment
	var total int64

	err := r.DB.Model(&commentDomain.Comment{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * limit
	if err = r.DB.Limit(limit).Offset(offset).Find(&comments).Error; err != nil {
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

	return &commentDomain.PaginationComment{
		Data:       commentDomain.ArrayToDomainMapper(&comments),
		Total:      total,
		Limit:      int64(limit),
		Current:    int64(page),
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

// UserGetAll Fetch all comment data
func (r *Repository) UserGetAll(userId int, page int, limit int) (*commentDomain.PaginationComment, error) {
	var comments []commentDomain.Comment
	var total int64

	err := r.DB.Model(&commentDomain.Comment{}).Where("user_id = ?", userId).Count(&total).Error
	if err != nil {
		return nil, err
	}
	offset := (page - 1) * limit
	err = r.DB.Limit(limit).Offset(offset).Find(&comments).Error

	if err != nil {
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

	return &commentDomain.PaginationComment{
		Data:       commentDomain.ArrayToDomainMapper(&comments),
		Total:      total,
		Limit:      int64(limit),
		Current:    int64(page),
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
			return nil, fiber.NewError(fiber.StatusNotFound, "comment not found")
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
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
			return nil, fiber.NewError(fiber.StatusNotFound, "comment not found")
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
	}

	return &comment, nil
}

// GetOneByMap ... Fetch only one comment by Map
func (r *Repository) GetOneByMap(commentMap map[string]interface{}) (*commentDomain.Comment, error) {
	var comment commentDomain.Comment

	err := r.DB.Where(commentMap).Limit(1).Find(&comment).Error
	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		return nil, err
	}
	return &comment, nil
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
			return nil, fiber.NewError(fiber.StatusConflict, mssgConst.ResourceAlreadyExists)
		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
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

	err = r.DB.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "comment not found")
	}

	return &comment, nil
}

// UserUpdate ... UserUpdate comment
func (r *Repository) UserUpdate(id int, userId int, updateComment *commentDomain.Comment) (*commentDomain.Comment, error) {
	var comment commentDomain.Comment

	comment.ID = id
	comment.UserID = userId
	err := r.DB.Model(&comment).
		Updates(updateComment).Error

	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errorDomain.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &commentDomain.Comment{}, err
		}

		switch newError.Number {
		case 1062:
			return nil, fiber.NewError(fiber.StatusConflict, mssgConst.ResourceAlreadyExists)

		default:
			return nil, fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
		}
	}

	err = r.DB.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "comment not found")
	}

	return &comment, nil
}

// Delete ... Delete comment
func (r *Repository) Delete(id int) (err error) {
	tx := r.DB.Delete(&commentDomain.Comment{}, id)

	log.Println("check ", tx)
	if tx.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, mssgConst.UnknownError)
	}

	if tx.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "comment not found")
	}

	return
}
