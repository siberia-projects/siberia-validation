// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

import "fmt"

const (
	notBlankConstraintDefaultErrorMessage = "A value should not be blank"
)

type NotBlankConstraint struct {
	ErrorMessage string
}

func (constraint *NotBlankConstraint) GetValidatorName() string {
	return notBlankConstraintValidatorName
}

func (constraint *NotBlankConstraint) GetViolationError() error {
	errorMessage := notBlankConstraintDefaultErrorMessage
	if constraint.ErrorMessage != "" {
		errorMessage = constraint.ErrorMessage
	}

	return fmt.Errorf(errorMessage)
}
