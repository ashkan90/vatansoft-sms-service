package validation

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

type IValidator interface {
	Validate(i interface{}) map[string]string
}

type validation struct {
	validator *validator.Validate
}

func InitValidator() IValidator {
	v := validator.New()

	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		if jsonField := field.Tag.Get("json"); jsonField != "" && jsonField != "-" {
			return jsonField
		}
		if formField := field.Tag.Get("form"); formField != "" && formField != "-" {
			return formField
		}

		return field.Tag.Get("query")
	})

	return &validation{
		validator: v,
	}
}

func (v *validation) Validate(i interface{}) map[string]string {
	messages := make(map[string]string)
	if errors := v.validator.Struct(i); errors != nil {
		for _, err := range errors.(validator.ValidationErrors) {
			messages[err.Field()] = err.Error()
		}
	}

	return messages
}
