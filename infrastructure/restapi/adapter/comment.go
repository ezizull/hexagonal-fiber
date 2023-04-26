package adapter

import (
	commentService "hacktiv/final-project/application/usecases/comment"
	commentRepository "hacktiv/final-project/infrastructure/repository/postgres/comment"
	photoRepository "hacktiv/final-project/infrastructure/repository/postgres/photo"
	commentController "hacktiv/final-project/infrastructure/restapi/controllers/comment"

	"gorm.io/gorm"
)

// CommentAdapter is a function that returns a comment controller
func CommentAdapter(db *gorm.DB) *commentController.Controller {
	cRepository := commentRepository.Repository{DB: db}
	pRepository := photoRepository.Repository{DB: db}

	service := commentService.Service{
		CommentRepository: cRepository,
		PhotoRepository:   pRepository,
	}
	return &commentController.Controller{CommentService: service}
}
