package shared

import (
	"context"
	"github.com/viant/govalidator"
)

var validator = govalidator.New()

// Validation represents validation info
type Validation struct {
	govalidator.Validation
}

func (v *Validation) Validate(any interface{}, options ...govalidator.Option) bool {
	validation, err := validator.Validate(context.Background(), any, options...)
	if err != nil {
		validation.AddViolation("", "", "error", err.Error())
		return false
	}

	if validation != nil {
		v.Violations = append(v.Violations, validation.Violations...)
		if validation.Failed {
			v.Failed = validation.Failed
		}
	}
	return v.Failed
}

func NewValidationInfo() *Validation {
	return &Validation{}
}
