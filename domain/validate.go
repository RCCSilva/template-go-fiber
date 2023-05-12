package domain

import "github.com/go-playground/validator"

var v = validator.New()

type InvalidFields struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func validate(data any) error {
	err := v.Struct(data)
	if err == nil {
		return nil
	}

	var errors []*InvalidFields

	for _, err := range err.(validator.ValidationErrors) {
		vr := InvalidFields{
			Field: err.Field(),
			Tag:   err.Tag(),
			Value: err.Param(),
		}
		errors = append(errors, &vr)
	}

	if len(errors) == 0 {
		return nil
	}

	return &ResultError{
		StatusCode: 400,
		Message:    "bad request",
		Fields:     errors,
	}
}
