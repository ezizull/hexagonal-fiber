package comment

import commentDomain "hexagonal-fiber/domain/comment"

type CommentTesting interface {
	GetAll(page int, limit int) (*commentDomain.PaginationComment, error)
	UserGetAll(page int, userId int, limit int) (*commentDomain.PaginationComment, error)
	Create(newComment *commentDomain.Comment) (createdComment *commentDomain.Comment, err error)
	GetByID(id int) (*commentDomain.Comment, error)
	UserGetByID(id int, userId int) (*commentDomain.Comment, error)
	GetOneByMap(commentMap map[string]interface{}) (*commentDomain.Comment, error)
	Update(id int, updateComment *commentDomain.Comment) (*commentDomain.Comment, error)
	Delete(id int) (err error)
}
