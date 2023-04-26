package comment

import (
	commentDomain "hexagonal-fiber/domain/comment"
	commentRepository "hexagonal-fiber/infrastructure/repository/postgres/comment"
)

type CommentTesting interface {
	GetAll(page int, limit int) (*commentDomain.PaginationComment, error)
	UserGetAll(userId int, page int, limit int) (*commentDomain.PaginationComment, error)
	GetByID(id int) (*commentDomain.Comment, error)
	UserGetByID(id int, userId int) (*commentDomain.Comment, error)
	Create(comment *commentDomain.NewComment) (*commentDomain.Comment, error)
	GetByMap(commentMap map[string]interface{}) (*commentDomain.Comment, error)
	Delete(id int) (err error)
	Update(id int, updateComment commentDomain.UpdateComment) (*commentDomain.Comment, error)
}

func NewTesting(commentTest commentRepository.CommentTesting) CommentTesting {
	return &Service{
		CommentTesting: commentTest,
	}
}
