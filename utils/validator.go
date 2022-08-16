package utils

import "github.com/go-playground/validator/v10"

func GetValidationResult(err error) map[string]string {
	errResult := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		errResult[err.Field()] = err.Tag()
	}
	return errResult
}
