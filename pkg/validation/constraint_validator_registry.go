// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

// ConstraintValidatorRegistry is an interface which provides
// a functionality for storing all ConstraintValidator(s) in
// one place
type ConstraintValidatorRegistry interface {
	// Get gets a ConstraintValidator based on
	// passed validator's name
	//
	// Returns an error when such validator doesn't exist
	Get(name string) (ConstraintValidator, error)

	// Add adds a new ConstraintValidator into a ConstraintValidatorRegistry
	Add(validator ConstraintValidator)

	// Remove removes a ConstraintValidator from a ConstraintValidatorRegistry
	Remove(validator ConstraintValidator)
}
