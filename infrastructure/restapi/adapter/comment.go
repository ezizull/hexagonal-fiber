package adapter

import (
	commentService "hexagonal-fiber/application/usecases/comment"
	commentRepository "hexagonal-fiber/infrastructure/repository/postgres/comment"
	photoRepository "hexagonal-fiber/infrastructure/repository/postgres/photo"
	commentController "hexagonal-fiber/infrastructure/restapi/controllers/comment"

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
