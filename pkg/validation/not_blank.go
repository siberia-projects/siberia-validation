// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

import "fmt"

const (
	notBlankDefaultErrorMessage = "A value should not be blank"
)

type NotBlank struct {
	ErrorMessage string
}

func (constraint *NotBlank) GetValidatorName() string {
	return notBlankValidatorName
}

func (constraint *NotBlank) GetViolationError() error {
	errorMessage := notBlankDefaultErrorMessage
	if constraint.ErrorMessage != "" {
		errorMessage = constraint.ErrorMessage
	}

	return fmt.Errorf(errorMessage)
}
