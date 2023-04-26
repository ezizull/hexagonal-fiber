// Package user contains the user controller
package user

import (
	"errors"
	userDomain "hexagonal-fiber/domain/user"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func updateValidation(request *userDomain.UpdateUser) (err error) {
	var errorsValidation []string

	// Username must have minimum length of 4
	if request.UserName != nil {
		if len(*request.UserName) < 4 {
			errorsValidation = append(errorsValidation, "Username must be at least 4 characters long")
		}
	}

	// Email must be a valid email format
	if request.Email != nil {
		if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(*request.Email) {
			errorsValidation = append(errorsValidation, "Invalid email format")
		}
	}

	// Password must have minimum length of 6, at least 1 special character, 1 capital letter, 1 lowercase letter, and 1 number
	if request.Password != nil {
		if len(*request.Password) < 6 {
			errorsValidation = append(errorsValidation, "password must be at least 6 characters long")
		}
		hasSpecialChar := regexp.MustCompile(`[^a-zA-Z0-9]+`).MatchString
		if !hasSpecialChar(*request.Password) {
			errorsValidation = append(errorsValidation, "password must contain at least one special character")
		}
		hasCapitalLetter := regexp.MustCompile(`[A-Z]+`).MatchString
		if !hasCapitalLetter(*request.Password) {
			errorsValidation = append(errorsValidation, "password must contain at least one capital letter")
		}
		hasLowerCase := regexp.MustCompile(`[a-z]+`).MatchString
		if !hasLowerCase(*request.Password) {
			errorsValidation = append(errorsValidation, "password must contain at least one lowercase letter")
		}
		hasNumber := regexp.MustCompile(`[0-9]+`).MatchString
		if !hasNumber(*request.Password) {
			errorsValidation = append(errorsValidation, "password must contain at least one number")
		}
	}

	// Age must have minimum 8 old
	if request.Age != nil {
		if *request.Age < 8 {
			errorsValidation = append(errorsValidation, "Age must be at least 8 old")
		}
	}

	if errorsValidation != nil {
		err = fiber.NewError(fiber.StatusBadRequest, strings.Join(errorsValidation, ", "))
	}
	return
}

func createValidation(request userDomain.NewUser) (err error) {
	// Username must have minimum length of 4
	if len(request.UserName) < 4 {
		return errors.New("Username must be at least 4 characters long")
	}

	// Email must be a valid email format
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(request.Email) {
		return errors.New("Invalid email format")
	}

	// Password must have minimum length of 6, at least 1 special character, 1 capital letter, 1 lowercase letter, and 1 number
	if len(request.Password) < 6 {
		return errors.New("Password should be at least 6 characters long")
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
