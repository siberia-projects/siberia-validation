// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe(
	"Verifying a simple validator's functionality",
	Label("simple_validator"),
	func() {
		When("a validation is going", func() {
			registry := NewSimpleConstraintValidatorRegistryWithDefaultValidators()
			validator := NewSimpleValidator(registry)

			Context("and no constraint's validator is found for passed constraint", func() {
				It("should return an error", func() {
					// when
					err := validator.Validate("valid_value", []Constraint{
						&testConstraint{},
					})

					// then
					Expect(err).ToNot(BeNil())
				})
			})

			Context("and one validator returned an error", func() {
				It("should return a violation error contains 1 validator's error", func() {
					// when
					err := validator.Validate("", []Constraint{
						&NotNil{},
						&NotBlank{},
					})

					// then
					Expect(err).ToNot(BeNil())

					errCasted, ok := err.(*ConstraintViolationError)

					Expect(ok).To(BeTrue())
					Expect(errCasted.errorMessages).To(HaveLen(1))
				})
			})

			Context("and all validators returned errors", func() {
				It("should return a violation error contains all validator' errors", func() {
					// when
					err := validator.Validate(nil, []Constraint{
						&NotNil{},
					})

					// then
					Expect(err).ToNot(BeNil())

					errCasted, ok := err.(*ConstraintViolationError)

					Expect(ok).To(BeTrue())
					Expect(errCasted.errorMessages).To(HaveLen(1))
				})
			})

			Context("and value is valid", func() {
				It("should not return any error", func() {
					// given
					value := "valid_value"

					// when
					err := validator.Validate(value, []Constraint{
						&NotNil{},
						&NotBlank{},
					})

					// then
					Expect(err).To(BeNil())
				})
			})
		})
	},
)

type testConstraint struct {
}

func (constraint *testConstraint) GetValidatorName() string {
	return "test"
}

func (constraint *testConstraint) GetViolationError() error {
	return fmt.Errorf("invalid")
}
