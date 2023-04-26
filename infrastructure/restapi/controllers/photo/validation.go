package photo

import (
	photoDomain "hexagonal-fiber/domain/photo"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func updateValidation(request *photoDomain.UpdatePhoto) (err error) {
	var errorsValidation []string

	// Title cannot be empty
	if request.Title != nil {
		if len(*request.Title) < 1 {
			errorsValidation = append(errorsValidation, "Title cannot be empty")
		}
	}

	// PhotoUrl cannot be empty
	if request.PhotoUrl != nil {
		if len(*request.PhotoUrl) < 1 {
			errorsValidation = append(errorsValidation, "PhotoUrl cannot be empty")
		}
	}

	if errorsValidation != nil {
		err = fiber.NewError(fiber.StatusBadRequest, strings.Join(errorsValidation, ", "))
	}
	return
}

func createValidation(request photoDomain.NewPhoto) (err error) {
	var errorsValidation []string

	// Title cannot be empty
	if len(request.Title) < 1 {
		errorsValidation = append(errorsValidation, "Title cannot be empty")
	}

	// PhotoUrl cannot be empty
	if len(request.PhotoUrl) < 1 {
		errorsValidation = append(errorsValidation, "PhotoUrl cannot be empty")
	}

	if errorsValidation != nil {
		err = fiber.NewError(fiber.StatusBadRequest, strings.Join(errorsValidation, ", "))
	}
	return
}
