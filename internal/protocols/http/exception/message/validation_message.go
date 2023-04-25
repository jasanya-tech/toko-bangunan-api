package message

import (
	"github.com/go-playground/validator/v10"
)

type ValidationRespon struct {
	Field     string
	Tag       string
	ActualTag string
	Type      string
	Value     any
	Param     string
}

func ValidationMessage(errs error) []ValidationRespon {
	var validations []ValidationRespon

	for _, err := range errs.(validator.ValidationErrors) {
		validation := &ValidationRespon{
			Field:     err.Field(),
			Tag:       err.Tag(),
			ActualTag: err.ActualTag(),
			Type:      err.Type().Name(),
			Value:     err.Value(),
			Param:     err.Param(),
		}
		validations = append(validations, *validation)
	}
	return validations
}
