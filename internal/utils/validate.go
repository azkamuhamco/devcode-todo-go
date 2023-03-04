package utils

import (
	"devcode-todo-go/internal/models"

	"github.com/go-playground/validator"
)

var validate = validator.New()

func ValidateStruct(data interface{}) []*models.ErrorResponse {
	var errors []*models.ErrorResponse
	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element models.ErrorResponse
			var errorField string

			if err.Field() == "Title" {
				errorField = "title"
			} else if err.Field() == "ActivityGroupID" {
				errorField = "activity_group_id"
			}

			element.FailedField = errorField
			errors = append(errors, &element)
		}
	}
	return errors
}
