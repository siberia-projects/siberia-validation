// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

import "strings"

// ConstraintViolationError is an error which stores
// inside itself all validation errors
type ConstraintViolationError struct {
	errorMessages []string
}

// NewConstraintViolationError creates a new ConstraintViolationError
func NewConstraintViolationError() *ConstraintViolationError {
	errorMessages := make([]string, 0)
	return &ConstraintViolationError{errorMessages: errorMessages}
}

// Add adds an error into an array of errors
func (err *ConstraintViolationError) Add(errorToAdd error) {
	err.errorMessages = append(err.errorMessages, errorToAdd.Error())
}

func (err *ConstraintViolationError) Error() string {
	return strings.Join(err.errorMessages, "\n")
}
