package common

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validate(data interface{}) error {
	valid := validator.New()
	err := valid.Struct(data)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		error_message := ""
		for _, err := range err.(validator.ValidationErrors) {
			error_message += fmt.Sprintf("Param %v %v. ", err.Field(), err.Tag())
		}
		return errors.New(error_message)
	}
	return nil
}
