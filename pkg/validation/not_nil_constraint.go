// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

import "fmt"

const (
	notNilConstraintDefaultErrorMessage = "A value should not be nil"
)

type NotNilConstraint struct {
	ErrorMessage string
}

func (constraint *NotNilConstraint) GetValidatorName() string {
	return notNilConstraintValidatorName
}

func (constraint *NotNilConstraint) GetViolationError() error {
	errorMessage := notNilConstraintDefaultErrorMessage
	if constraint.ErrorMessage != "" {
		errorMessage = constraint.ErrorMessage
	}

	return fmt.Errorf(errorMessage)
}
