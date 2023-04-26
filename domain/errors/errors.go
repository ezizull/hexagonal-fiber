// Package errors defines the domain errors used in the application.
package errors

import "errors"

const (
	// NotFound error indicates a missing / not found record
	NotFound        = "NotFound"
	NotFoundMessage = "record not found"

	// ValidationError indicates an error in input validation
	ValidationError        = "ValidationError"
	ValidationErrorMessage = "validation error"

	// ResourceAlreadyExists indicates a duplicate / already existing record
	ResourceAlreadyExists     = "ResourceAlreadyExists"
	AlreadyExistsErrorMessage = "resource already exists"

	// RepositoryError indicates a repository (e.g database) error
	RepositoryError        = "RepositoryError"
	RepositoryErrorMessage = "error in repository operation"

	// NotAuthenticated indicates an authentication error
	NotAuthenticated             = "NotAuthenticated"
	NotAuthenticatedErrorMessage = "not Authenticated"

	// TokenGeneratorError indicates an token generation error
	TokenGeneratorError        = "TokenGeneratorError"
	TokenGeneratorErrorMessage = "error in token generation"

	// NotAuthorized indicates an authorization error
	NotAuthorized             = "NotAuthorized"
	NotAuthorizedErrorMessage = "not authorized"

	// UnknownError indicates an error that the app cannot find the cause for
	UnknownError        = "UnknownError"
	UnknownErrorMessage = "something went wrong"
)

// AppError defines an application (domain) error
type AppError struct {
	Err  error
	Type string
}

// NewAppError initializes a new domain error using an error and its type.
func NewAppError(err error, errType string) *AppError {
	return &AppError{
		Err:  err,
		Type: errType,
	}
}

// NewAppErrorWithType initializes a new default error for a given type.
func NewAppErrorWithType(errType string) *AppError {
	var err error

	switch errType {
	case NotFound:
		err = errors.New(NotFoundMessage)
	case ValidationError:
		err = errors.New(ValidationErrorMessage)
	case ResourceAlreadyExists:
		err = errors.New(AlreadyExistsErrorMessage)
	case RepositoryError:
		err = errors.New(RepositoryErrorMessage)
	case NotAuthenticated:
		err = errors.New(NotAuthenticatedErrorMessage)
	case NotAuthorized:
		err = errors.New(NotAuthorizedErrorMessage)
	case TokenGeneratorError:
		err = errors.New(TokenGeneratorErrorMessage)
	default:
		err = errors.New(UnknownErrorMessage)
	}

	return &AppError{
		Err:  err,
		Type: errType,
	}
}

// String converts the app error to a human-readable string.
func (appErr *AppError) Error() string {
	return appErr.Err.Error()
}
