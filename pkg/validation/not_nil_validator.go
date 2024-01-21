// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

const (
	notNilValidatorName = "not_nil"
)

type NotNilValidator struct {
}

func NewNotNilValidator() ConstraintValidator {
	return &NotNilValidator{}
}

func (validator *NotNilValidator) GetName() string {
	return notNilValidatorName
}

func (validator *NotNilValidator) Validate(value any, constraint Constraint) error {
	if value == nil {
		return constraint.GetViolationError()
	}

	return nil
}
