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
	"Verifying a not nil constraint's functionality",
	Label("not_nil_constraint"),
	func() {
		When("a getting of a validator's name is going", func() {
			It("should return not nil validator's name", func() {
				// given
				constraint := &NotNilConstraint{}

				// when
				validatorName := constraint.GetValidatorName()

				// then
				Expect(validatorName).To(Equal(notNilConstraintValidatorName))
			})
		})

		When("a getting of a violation error is going", func() {
			Context("and an error message is not specified", func() {
				It("should return a default error", func() {
					// given
					constraint := &NotNilConstraint{}

					// when
					err := constraint.GetViolationError()

					// then
					Expect(err).ToNot(BeNil())
					Expect(err.Error()).To(Equal(notNilConstraintDefaultErrorMessage))
				})
			})

			Context("and an error message is specified", func() {
				It("should return the error message", func() {
					// given
					errorMessage := "invalid"
					constraint := &NotNilConstraint{ErrorMessage: errorMessage}

					// when
					err := constraint.GetViolationError()

					// then
					Expect(err).ToNot(BeNil())
					Expect(err.Error()).To(Equal(errorMessage))
				})
			})
		})
	},
)
