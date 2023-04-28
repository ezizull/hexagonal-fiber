package comment

import (
	commentDomain "hexagonal-fiber/domain/comment"
	commentRepository "hexagonal-fiber/infrastructure/repository/postgres/comment"
)

type CommentTesting interface {
	GetAll(page int, limit int) (*commentDomain.PaginationComment, error)
	UserGetAll(userId string, page int, limit int) (*commentDomain.PaginationComment, error)
	GetByID(id string) (*commentDomain.Comment, error)
	UserGetByID(id string, userId string) (*commentDomain.Comment, error)
	Create(comment *commentDomain.NewComment) (*commentDomain.Comment, error)
	GetByMap(commentMap map[string]interface{}) (*commentDomain.Comment, error)
	Delete(id string) (err error)
	Update(id string, updateComment commentDomain.UpdateComment) (*commentDomain.Comment, error)
}

func NewTesting(commentTest commentRepository.CommentTesting) CommentTesting {
	return &Service{
		CommentTesting: commentTest,
	}
}
