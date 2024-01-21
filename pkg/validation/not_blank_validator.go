// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

const (
	notBlankValidatorName = "not_blank"
)

type NotBlankValidator struct {
}

func NewNotBlankValidator() ConstraintValidator {
	return &NotBlankValidator{}
}

func (validator *NotBlankValidator) GetName() string {
	return notBlankValidatorName
}

func (validator *NotBlankValidator) Validate(value any, constraint Constraint) error {
	valueString, ok := value.(string)
	if !ok {
		return constraint.GetViolationError()
	}

	if valueString == "" {
		return constraint.GetViolationError()
	}

	return nil
}
