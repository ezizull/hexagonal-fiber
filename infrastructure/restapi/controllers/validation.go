package controllers

import (
	commentDomain "hexagonal-fiber/domain/comment"
	errorDomain "hexagonal-fiber/domain/error"
	photoDomain "hexagonal-fiber/domain/photo"
	sosmedDomain "hexagonal-fiber/domain/sosmed"
	userDomain "hexagonal-fiber/domain/user"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type generic interface {
	userDomain.LoginUser | userDomain.NewUser | sosmedDomain.NewSocialMedia | photoDomain.NewPhoto | commentDomain.NewComment
}

func Validation[T generic](object T) (T, []*errorDomain.ErrorResponse) {
	var errors []*errorDomain.ErrorResponse
	err := validate.Struct(object)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element errorDomain.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return object, errors
}
