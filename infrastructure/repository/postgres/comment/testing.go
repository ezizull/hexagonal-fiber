package comment

import commentDomain "hacktiv/final-project/domain/comment"

type CommentTesting interface {
	GetAll(page int64, limit int64) (*commentDomain.PaginationResultComment, error)
	UserGetAll(page int64, userId int, limit int64) (*commentDomain.PaginationResultComment, error)
	Create(newComment *commentDomain.Comment) (createdComment *commentDomain.Comment, err error)
	GetByID(id int) (*commentDomain.Comment, error)
	UserGetByID(id int, userId int) (*commentDomain.Comment, error)
	GetOneByMap(commentMap map[string]interface{}) (*commentDomain.Comment, error)
	Update(id int, updateComment *commentDomain.Comment) (*commentDomain.Comment, error)
	Delete(id int) (err error)
}
