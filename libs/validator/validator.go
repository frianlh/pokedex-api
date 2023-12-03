package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate = validator.New()

// ValidateStruct is
func ValidateStruct(modelStruct interface{}) (err error) {
	errs := validate.Struct(modelStruct)
	if errs != nil {
		var errString []string
		for i := 0; i < len(errs.(validator.ValidationErrors)); i++ {
			errString = append(errString, (errs.(validator.ValidationErrors)[i]).Error())
		}
		return errors.New(strings.Join(errString, ", "))
	}
	return nil
}
