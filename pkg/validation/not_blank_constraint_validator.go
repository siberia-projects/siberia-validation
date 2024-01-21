// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

const (
	notBlankConstraintValidatorName = "not_blank"
)

type NotBlankConstraintValidator struct {
}

func NewNotBlankConstraintValidator() ConstraintValidator {
	return &NotBlankConstraintValidator{}
}

func (validator *NotBlankConstraintValidator) GetName() string {
	return notBlankConstraintValidatorName
}

func (validator *NotBlankConstraintValidator) Validate(value any, constraint Constraint) error {
	valueString, ok := value.(string)
	if !ok {
		return constraint.GetViolationError()
	}

	if valueString == "" {
		return constraint.GetViolationError()
	}

	return nil
}
