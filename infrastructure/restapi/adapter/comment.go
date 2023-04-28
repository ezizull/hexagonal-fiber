package adapter

import (
	commentService "hexagonal-fiber/application/usecases/comment"
	databsDomain "hexagonal-fiber/domain/database"
	commentRepository "hexagonal-fiber/infrastructure/repository/postgres/comment"
	photoRepository "hexagonal-fiber/infrastructure/repository/postgres/photo"
	commentController "hexagonal-fiber/infrastructure/restapi/controllers/comment"
)

// CommentAdapter is a function that returns a comment controller
func CommentAdapter(db databsDomain.Database) *commentController.Controller {
	cRepository := commentRepository.Repository{DB: db.Postgre}
	pRepository := photoRepository.Repository{DB: db.Postgre}

	service := commentService.Service{
		CommentRepository: cRepository,
		PhotoRepository:   pRepository,
	}
	return &commentController.Controller{CommentService: service}
}
