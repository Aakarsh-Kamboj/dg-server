package infrastructure

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	Validator *validator.Validate
}

// Validate is invoked by Echo during c.Bind(...) if v := e.Validator; v != nil
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
