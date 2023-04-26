package sosmed

import (
	"errors"
	errorDomain "hacktiv/final-project/domain/errors"
	sosmedDomain "hacktiv/final-project/domain/sosmed"
	"strings"
)

func updateValidation(request *sosmedDomain.UpdateSocialMedia) (err error) {
	var errorsValidation []string

	// Name cannot be empty
	if request.Name != nil {
		if len(*request.Name) < 1 {
			errorsValidation = append(errorsValidation, "Name cannot be empty")
		}
	}

	// SocialMediaUrl cannot be empty
	if request.SocialMediaUrl != nil {
		if len(*request.SocialMediaUrl) < 1 {
			errorsValidation = append(errorsValidation, "SocialMediaUrl cannot be empty")
		}
	}

	if errorsValidation != nil {
		err = errorDomain.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), errorDomain.ValidationError)
	}
	return
}

func createValidation(request sosmedDomain.NewSocialMedia) (err error) {
	// Name cannot be empty
	if len(request.Name) < 1 {
		return errors.New("Name cannot be empty")
	}

	// SocialMediaUrl cannot be empty
	if len(request.SocialMediaUrl) < 1 {
		return errors.New("SocialMediaUrl cannot be empty")
	}
	return
}
