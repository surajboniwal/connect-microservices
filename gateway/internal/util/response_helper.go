package util

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ParseError(err error) map[string]interface{} {
	errResponse := map[string]interface{}{}
	for _, e := range err.(validator.ValidationErrors) {
		errResponse[strings.ToLower(e.Field())] = errorString(e)
	}

	return errResponse
}

func errorString(e validator.FieldError) string {
	switch e.Tag() {
	case "email":
		return "Please enter a valid email address"
	case "required":
		return fmt.Sprintf("%v is required", e.Field())
	case "gte":
		return fmt.Sprintf("%s must be atleast %s characters long", e.Field(), e.Param())
	default:
		return "Invalid input"
	}
}
