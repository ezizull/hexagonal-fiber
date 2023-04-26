package comment

import (
	"errors"
	commentDomain "hexagonal-fiber/domain/comment"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func updateValidation(request *commentDomain.UpdateComment) (err error) {
	var errorsValidation []string

	// Message cannot be empty
	if request.Message != nil {
		if len(*request.Message) < 1 {
			errorsValidation = append(errorsValidation, "Message cannot be empty")
		}
	}

	if errorsValidation != nil {
		err = fiber.NewError(fiber.StatusBadRequest, strings.Join(errorsValidation, ", "))
	}

	return
}

func createValidation(request commentDomain.NewComment) (err error) {
	// Message cannot be empty
	if len(request.Message) < 1 {
		return errors.New("Message cannot be empty")
	}

	// PhotoID please insert correct id
	if request.PhotoID < 1 {
		return errors.New("PhotoID please insert correct id")
	}

	return
}

func getValidation(request commentDomain.GetComment) (err error) {
	// PhotoID please insert correct id
	if request.PhotoID < 1 {
		return errors.New("PhotoID please insert correct id")
	}

	return
}
