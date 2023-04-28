package comment

import commentDomain "hexagonal-fiber/domain/comment"

type CommentTesting interface {
	GetAll(page int, limit int) (*commentDomain.PaginationComment, error)
	UserGetAll(page int, userId string, limit int) (*commentDomain.PaginationComment, error)
	Create(newComment *commentDomain.Comment) (createdComment *commentDomain.Comment, err error)
	GetByID(id string) (*commentDomain.Comment, error)
	UserGetByID(id string, userId string) (*commentDomain.Comment, error)
	GetOneByMap(commentMap map[string]interface{}) (*commentDomain.Comment, error)
	Update(id string, updateComment *commentDomain.Comment) (*commentDomain.Comment, error)
	Delete(id string) (err error)
}
