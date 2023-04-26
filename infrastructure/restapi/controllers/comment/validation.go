package comment

import (
	"errors"
	commentDomain "hacktiv/final-project/domain/comment"
	errorDomain "hacktiv/final-project/domain/errors"
	"strings"
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
		err = errorDomain.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), errorDomain.ValidationError)
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
