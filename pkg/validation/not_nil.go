// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

import "fmt"

const (
	notNilDefaultErrorMessage = "A value should not be nil"
)

type NotNil struct {
	ErrorMessage string
}

func (constraint *NotNil) GetValidatorName() string {
	return notNilValidatorName
}

func (constraint *NotNil) GetViolationError() error {
	errorMessage := notNilDefaultErrorMessage
	if constraint.ErrorMessage != "" {
		errorMessage = constraint.ErrorMessage
	}

	return fmt.Errorf(errorMessage)
}
