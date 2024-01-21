// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

// ConstraintValidator is an interface which provides
// a validation's functionality for a specific Constraint
//
// It validates a passed value with Constraint's data
type ConstraintValidator interface {
	// GetName gets a name of a ConstraintValidator
	GetName() string

	// Validate validates a value using passed Constraint
	//
	// Returns an error when value is not valid
	Validate(value any, constraint Constraint) error
}
