// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

// Constraint is an interface which provides
// a functionality of storing important information for
// a further validation
//
// e.g. An "EmailConstraint" can store inside itself a regexp of email,
// a violation message etc
type Constraint interface {
	// GetValidatorName gets a name of Validator
	// which is responsible for validation of this Constraint
	GetValidatorName() string

	// GetViolationError gets a violation error
	GetViolationError() error
}
