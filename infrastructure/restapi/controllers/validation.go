package controllers

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// initialize validation
var validate = validator.New()

// list errors
var tagErrors = map[string]string{}

// Validation funct to validate all request
func Validation(object interface{}) (err error) {
	var errors []string

	if err = customValidation(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "wrong before validation")
	}

	err = validate.Struct(object)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var message string

			if err.ActualTag() != "password" {
				message = fmt.Sprintf("%s missing validatate %v= %v", err.StructNamespace(), err.Tag(), err.Param())
			} else {
				message = fmt.Sprintf("%s missing validatate %v= %v", err.StructNamespace(), err.Tag(), tagErrors["password"])
			}

			errors = append(errors, message)
		}
	}

	if errors != nil {
		return fiber.NewError(fiber.StatusBadRequest, strings.Join(errors, ", "))
	}

	return err
}

// customValidation func for handle all custom param validation
func customValidation() (err error) {

	// Password must have minimum length of 8, at least 1 special character, 1 capital letter, 1 lowercase letter, and 1 number
	validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		tagErrors["password"] = ""

		if len(password) < 8 {
			tagErrors["password"] += "min 8 long, "
		}

		if matched, _ := regexp.MatchString(`[A-Z]+`, password); !matched {
			tagErrors["password"] += "1 capital character, "
		}

		if matched, _ := regexp.MatchString(`[a-z]+`, password); !matched {
			tagErrors["password"] += "1 lowercase character, "
		}

		if matched, _ := regexp.MatchString(`[0-9]+`, password); !matched {
			tagErrors["password"] += "1 number, "
		}

		if matched, _ := regexp.MatchString(`[\W]+`, password); !matched {
			tagErrors["password"] += "1 special character, "
		}

		return tagErrors["password"] == ""
	})

	return err
}
