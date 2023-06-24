package utils

import (
	"errors"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ValidationUtil struct {
	validator *validator.Validate
}

func NewValidationUtil() echo.Validator {
	return &ValidationUtil{validator: validator.New()}
}

func (v *ValidationUtil) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func BindValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	if err := c.Validate(i); err != nil {
		return err
	}
	return nil
}


func ValidateImageSizeUnder1MB(c echo.Context, fieldName string) error {
	// Get the uploaded file from the request
	file, err := c.FormFile(fieldName)
	if err != nil {
		return err
	}
	// Get the size of the file in bytes
	size := file.Size

	// Check if the file is less than 1 MB in size
	if size > 1000000 {
		return errors.New("File size exceeds 1 MB limit")
	}

	return nil
}