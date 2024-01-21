// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

// SimpleValidator is a simple implementation of a Validator
type SimpleValidator struct {
	registry ConstraintValidatorRegistry
}

// NewSimpleValidator creates a new SimpleValidator
func NewSimpleValidator(registry ConstraintValidatorRegistry) Validator {
	return &SimpleValidator{registry: registry}
}

func (validator *SimpleValidator) Validate(value any, constraints []Constraint) error {
	violationError := NewConstraintViolationError()

	for _, constraint := range constraints {
		validatorName := constraint.GetValidatorName()

		constraintValidator, err := validator.registry.Get(validatorName)
		if err != nil {
			return err
		}

		err = constraintValidator.Validate(value, constraint)
		if err != nil {
			violationError.Add(err)
		}
	}

	if len(violationError.errorMessages) == 0 {
		return nil
	}

	return violationError
}
