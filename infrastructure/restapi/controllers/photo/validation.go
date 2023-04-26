package photo

import (
	"errors"
	errorDomain "hacktiv/final-project/domain/errors"
	photoDomain "hacktiv/final-project/domain/photo"
	"strings"
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
		err = errorDomain.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), errorDomain.ValidationError)
	}
	return
}

func createValidation(request photoDomain.NewPhoto) (err error) {
	// Title cannot be empty
	if len(request.Title) < 1 {
		return errors.New("Title cannot be empty")
	}

	// PhotoUrl cannot be empty
	if len(request.PhotoUrl) < 1 {
		return errors.New("PhotoUrl cannot be empty")
	}
	return
}
