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
	"Verifying a not nil constraint validator's functionality",
	Label("not_nil_constraint_validator"),
	func() {
		When("a getting of a validator's name is going", func() {
			It("should return not nil validator's name", func() {
				// given
				validator := NewNotNilConstraintValidator()

				// when
				validatorName := validator.GetName()

				// then
				Expect(validatorName).To(Equal(notNilConstraintValidatorName))
			})
		})

		When("a validation is going", func() {
			validator := NewNotNilConstraintValidator()

			Context("and a valid data is passed", func() {
				DescribeTable(
					"should not return any error",
					func(value any) {
						// given
						constraint := &NotNilConstraint{}

						// when
						err := validator.Validate(value, constraint)

						// then
						Expect(err).To(BeNil())
					},
					Entry("a valid value #1", ""),
					Entry("a valid value #2", "j"),
					Entry("a valid value #3", "john"),
					Entry("a valid value #4", 2),
					Entry("a valid value #5", 0),
					Entry("a valid value #6", '0'),
				)
			})

			Context("and an invalid data is passed", func() {
				DescribeTable(
					"should return an error",
					func(value any) {
						// given
						constraint := &NotNilConstraint{}

						// when
						err := validator.Validate(value, constraint)

						// then
						Expect(err).ToNot(BeNil())
					},
					Entry("an invalid value #1", nil),
				)
			})
		})
	},
)
