// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

// Validator is an interface which provides
// a functionality to validate a value using
// passed Constraint(s)
type Validator interface {
	// Validate validates passed value using
	// passed Constraint(s)
	//
	// Returns an error when value is invalid
	Validate(value any, constraints []Constraint) error
}
