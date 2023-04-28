package comment

import (
	commentDomain "hexagonal-fiber/domain/comment"

	commentRepository "hexagonal-fiber/infrastructure/repository/postgres/comment"
	photoRepository "hexagonal-fiber/infrastructure/repository/postgres/photo"
)

// Service is a struct that contains the repository implementation for comment use case
type Service struct {
	CommentTesting    commentRepository.CommentTesting
	CommentRepository commentRepository.Repository
	PhotoRepository   photoRepository.Repository
}

// GetAll is a function that returns all comments
func (s *Service) GetAll(page int, limit int) (*commentDomain.PaginationComment, error) {

	all, err := s.CommentRepository.GetAll(page, limit)
	if err != nil {
		return nil, err
	}

	return &commentDomain.PaginationComment{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

// UserGetAll is a function that returns all comments
func (s *Service) UserGetAll(userId string, page int, limit int) (*commentDomain.PaginationComment, error) {

	all, err := s.CommentRepository.UserGetAll(userId, page, limit)
	if err != nil {
		return nil, err
	}

	return &commentDomain.PaginationComment{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

// GetByID is a function that returns a comment by id
func (s *Service) GetByID(id string) (*commentDomain.Comment, error) {
	return s.CommentRepository.GetByID(id)
}

// UserGetByID is a function that returns a comment by id
func (s *Service) UserGetByID(id string, userId string) (*commentDomain.Comment, error) {
	return s.CommentRepository.UserGetByID(id, userId)
}

// Create is a function that creates a comment
func (s *Service) Create(comment *commentDomain.NewComment) (*commentDomain.Comment, error) {

	_, err := s.PhotoRepository.GetByID(comment.PhotoID)
	if err != nil {
		return nil, err
	}

	commentModel := comment.ToDomainMapper()

	return s.CommentRepository.Create(commentModel)
}

// GetByMap is a function that returns a comment by map
func (s *Service) GetByMap(commentMap map[string]interface{}) (*commentDomain.Comment, error) {
	return s.CommentRepository.GetOneByMap(commentMap)
}

// Delete is a function that deletes a comment by id
func (s *Service) Delete(id string) (err error) {
	return s.CommentRepository.Delete(id)
}

// Update is a function that updates a comment by id
func (s *Service) Update(id string, updateComment commentDomain.UpdateComment) (*commentDomain.Comment, error) {
	comment := updateComment.ToDomainMapper()
	return s.CommentRepository.Update(id, &comment)
}

// Update is a function that updates a comment by id
func (s *Service) UserUpdate(id string, userId string, updateComment commentDomain.UpdateComment) (*commentDomain.Comment, error) {
	comment := updateComment.ToDomainMapper()
	return s.CommentRepository.UserUpdate(id, userId, &comment)
}
