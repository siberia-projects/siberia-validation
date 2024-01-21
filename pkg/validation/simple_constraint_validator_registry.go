// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

import "fmt"

// SimpleConstraintValidatorRegistry is an implementation of
// a ConstraintValidatorRegistry
type SimpleConstraintValidatorRegistry struct {
	validators map[string]ConstraintValidator
}

// NewEmptySimpleConstraintValidatorRegistry creates
// a new empty SimpleConstraintValidatorRegistry
func NewEmptySimpleConstraintValidatorRegistry() ConstraintValidatorRegistry {
	validators := make(map[string]ConstraintValidator)
	return &SimpleConstraintValidatorRegistry{validators: validators}
}

// NewSimpleConstraintValidatorRegistryWithDefaultValidators creates
// a new SimpleConstraintValidatorRegistry with default ConstraintValidator(s)
func NewSimpleConstraintValidatorRegistryWithDefaultValidators() ConstraintValidatorRegistry {
	notNilValidator := NewNotNilValidator()
	notBlankValidator := NewNotBlankValidator()

	return NewSimpleConstraintValidatorRegistry(notBlankValidator, notNilValidator)
}

// NewSimpleConstraintValidatorRegistry creates
// a new SimpleConstraintValidatorRegistry with passed validators
func NewSimpleConstraintValidatorRegistry(validators ...ConstraintValidator) ConstraintValidatorRegistry {
	validatorsMap := make(map[string]ConstraintValidator)

	for _, validator := range validators {
		validatorName := validator.GetName()
		validatorsMap[validatorName] = validator
	}

	return &SimpleConstraintValidatorRegistry{validators: validatorsMap}
}

func (registry *SimpleConstraintValidatorRegistry) Get(name string) (ConstraintValidator, error) {
	validator, exists := registry.validators[name]
	if !exists {
		return nil, fmt.Errorf("a validator with \"%s\" name is not presented in a registry", name)
	}

	return validator, nil
}

func (registry *SimpleConstraintValidatorRegistry) Add(validator ConstraintValidator) {
	name := validator.GetName()

	_, exists := registry.validators[name]
	if exists {
		return
	}

	registry.validators[name] = validator
}

func (registry *SimpleConstraintValidatorRegistry) Remove(validator ConstraintValidator) {
	validatorName := validator.GetName()
	delete(registry.validators, validatorName)
}
