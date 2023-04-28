package auth

import (
	"errors"
	userDomain "hexagonal-fiber/domain/user"
	"regexp"
)

func createValidation(request userDomain.NewUser) (err error) {
	// Username must have minimum length of 4
	if len(request.UserName) < 4 {
		return errors.New("Username must be at least 4 characters long")
	}

	// Email must be a valid email format
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(request.Email) {
		return errors.New("Invalid email format")
	}

	// Password must have minimum length of 8, at least 1 special character, 1 capital letter, 1 lowercase letter, and 1 number
	if len(request.Password) < 8 {
		return errors.New("Password should be at least 8 characters long")
	}
	if !regexp.MustCompile(`[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]`).MatchString(request.Password) {
		return errors.New("Password should contain at least one special character")
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(request.Password) {
		return errors.New("Password should contain at least one uppercase letter")
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(request.Password) {
		return errors.New("Password should contain at least one lowercase letter")
	}
	if !regexp.MustCompile(`\d`).MatchString(request.Password) {
		return errors.New("Password should contain at least one number")
	}

	// Age must have minimum 8 old
	if request.Age < 8 {
		return errors.New("Age must be at least 8 old")
	}

	return
}
