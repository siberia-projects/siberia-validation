// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

const (
	notNilConstraintValidatorName = "not_nil"
)

type NotNilConstraintValidator struct {
}

func NewNotNilConstraintValidator() ConstraintValidator {
	return &NotNilConstraintValidator{}
}

func (validator *NotNilConstraintValidator) GetName() string {
	return notNilConstraintValidatorName
}

func (validator *NotNilConstraintValidator) Validate(value any, constraint Constraint) error {
	if value == nil {
		return constraint.GetViolationError()
	}

	return nil
}
