// Package user contains the user controller
package user

import (
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

	// Password must have minimum length of 8, at least 1 special character, 1 capital letter, 1 lowercase letter, and 1 number
	if request.Password != nil {
		if len(*request.Password) < 8 {
			errorsValidation = append(errorsValidation, "password must be at least 8 characters long")
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
