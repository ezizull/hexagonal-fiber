package message

const (
	// ValidationError indicates an error in input validation
	ValidationError = "validation error"

	// ResourceAlreadyExists indicates a duplicate / already existing record
	AlreadyExistsError = "resource already exists"

	// ResourceAlreadyExists indicates a duplicate / already existing record
	ResourceAlreadyExists = "resource already exists"

	// RepositoryError indicates a repository (e.g database) error
	RepositoryError = "error in repository operation"

	// NotAuthenticated indicates an authentication error
	NotAuthenticatedError = "not Authenticated"

	// TokenGeneratorError indicates an token generation error
	TokenGeneratorError = "error in token generation"

	// NotAuthorized indicates an authorization error
	NotAuthorizedError = "not authorized"

	// UnknownError indicates an error that the app cannot find the cause for
	UnknownError = "something went wrong"

	// StatusBadRequest indicates a bad request error
	StatusBadRequest = "bad request"

	// StatusUnauthorized indicates an unauthorized error
	StatusUnauthorized = "unauthorized"

	// StatusForbidden indicates a forbidden error
	StatusForbidden = "forbidden"

	// StatusNotFound indicates a not found error
	StatusNotFound = "not found"

	// StatusMethodNotAllowed indicates a method not allowed error
	StatusMethodNotAllowed = "method not allowed"

	// StatusRequestTimeout indicates a request timeout error
	StatusRequestTimeout = "request timeout"

	// StatusConflict indicates a conflict error
	StatusConflict = "conflict"

	// StatusUnsupportedMediaType indicates an unsupported media type error
	StatusUnsupportedMediaType = "unsupported media type"

	// StatusTooManyRequests indicates a too many requests error
	StatusTooManyRequests = "too many requests"

	// StatusInternalServerError indicates an internal server error
	StatusInternalServerError = "internal server error"

	// StatusNotImplemented indicates a not implemented error
	StatusNotImplemented = "not implemented"

	// StatusServiceUnavailable indicates a service unavailable error
	StatusServiceUnavailable = "service unavailable"
)
