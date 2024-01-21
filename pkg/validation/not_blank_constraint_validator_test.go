// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe(
	"Verifying a not blank constraint validator's functionality",
	Label("not_blank_constraint_validator"),
	func() {
		When("a getting of a validator's name is going", func() {
			It("should return not blank validator's name", func() {
				// given
				validator := NewNotBlankConstraintValidator()

				// when
				validatorName := validator.GetName()

				// then
				Expect(validatorName).To(Equal(notBlankConstraintValidatorName))
			})
		})

		When("a validation is going", func() {
			validator := NewNotBlankConstraintValidator()

			Context("and a valid data is passed", func() {
				DescribeTable(
					"should not return any error",
					func(value any) {
						// given
						constraint := &NotBlankConstraint{}

						// when
						err := validator.Validate(value, constraint)

						// then
						Expect(err).To(BeNil())
					},
					Entry("a valid value #1", "j"),
					Entry("a valid value #2", "john"),
				)
			})

			Context("and an invalid data is passed", func() {
				DescribeTable(
					"should return an error",
					func(value any) {
						// given
						constraint := &NotBlankConstraint{}

						// when
						err := validator.Validate(value, constraint)

						// then
						Expect(err).ToNot(BeNil())
					},
					Entry("an invalid value #1", ""),
					Entry("an invalid value #2", nil),
					Entry("an invalid value #3", 'a'),
				)
			})
		})
	},
)
