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
	"Verifying a simple constraint validator registry's functionality",
	Label("simple_constraint_validator_registry"),
	func() {
		When("an initialization is going", func() {
			Context("and an empty registry was requested", func() {
				It("should return a new instance of the registry without any validator", func() {
					// when
					registry := NewEmptySimpleConstraintValidatorRegistry()

					// then
					registryCasted, ok := registry.(*SimpleConstraintValidatorRegistry)

					Expect(ok).To(BeTrue())
					Expect(registryCasted.validators).To(HaveLen(0))
				})
			})

			Context("and a registry with passed validators was requested", func() {
				It("should return a new instance of the registry the validators", func() {
					// given
					notNilValidator := NewNotNilValidator()
					notNilValidatorName := notNilValidator.GetName()

					notBlankValidator := NewNotBlankValidator()
					notBlankValidatorName := notBlankValidator.GetName()

					// when
					registry := NewSimpleConstraintValidatorRegistry(
						notNilValidator,
						notBlankValidator,
					)

					// then
					registryCasted, ok := registry.(*SimpleConstraintValidatorRegistry)
					validators := registryCasted.validators

					Expect(ok).To(BeTrue())
					Expect(validators).To(HaveLen(2))

					notNilValidatorFound, ok := validators[notNilValidatorName]

					Expect(ok).To(BeTrue())
					Expect(notNilValidatorFound).ToNot(BeNil())

					notBlankValidatorFound, ok := validators[notBlankValidatorName]

					Expect(ok).To(BeTrue())
					Expect(notBlankValidatorFound).ToNot(BeNil())
				})
			})

			Context("and a registry with default validators was requested", func() {
				It("should return a new instance of the registry with default validators", func() {
					// given
					notNilValidator := NewNotNilValidator()
					notNilValidatorName := notNilValidator.GetName()

					notBlankValidator := NewNotBlankValidator()
					notBlankValidatorName := notBlankValidator.GetName()

					// when
					registry := NewSimpleConstraintValidatorRegistryWithDefaultValidators()

					// then
					registryCasted, ok := registry.(*SimpleConstraintValidatorRegistry)
					validators := registryCasted.validators

					Expect(ok).To(BeTrue())
					Expect(validators).To(HaveLen(2))

					notNilValidatorFound, ok := validators[notNilValidatorName]

					Expect(ok).To(BeTrue())
					Expect(notNilValidatorFound).ToNot(BeNil())

					notBlankValidatorFound, ok := validators[notBlankValidatorName]

					Expect(ok).To(BeTrue())
					Expect(notBlankValidatorFound).ToNot(BeNil())
				})
			})
		})

		When("an getting is going", func() {
			var registry ConstraintValidatorRegistry

			BeforeEach(func() {
				registry = NewEmptySimpleConstraintValidatorRegistry()
			})

			Context("and a validator doesn't exist in a registry yet", func() {
				It("should return an error", func() {
					// when
					validator, err := registry.Get("unknown_validator")

					// then
					Expect(err).ToNot(BeNil())
					Expect(validator).To(BeNil())
				})
			})

			Context("and a validator exists in a registry", func() {
				It("should successfully return it without any error", func() {
					// given
					notNilValidator := NewNotNilValidator()
					notNilValidatorName := notNilValidator.GetName()

					registry.Add(notNilValidator)

					// when
					validator, err := registry.Get(notNilValidatorName)

					// then
					Expect(err).To(BeNil())
					Expect(validator).ToNot(BeNil())
				})
			})
		})

		When("an adding is going", func() {
			var registry ConstraintValidatorRegistry

			BeforeEach(func() {
				registry = NewEmptySimpleConstraintValidatorRegistry()
			})

			Context("and a validator doesn't exist in a registry yet", func() {
				It("should successfully add it", func() {
					// given
					notNilValidator := NewNotNilValidator()
					notNilValidatorName := notNilValidator.GetName()

					// when / then
					registryCasted, ok := registry.(*SimpleConstraintValidatorRegistry)

					Expect(ok).To(BeTrue())
					Expect(registryCasted.validators).To(HaveLen(0))

					registry.Add(notNilValidator)
					Expect(registryCasted.validators).To(HaveLen(1))

					notNilValidatorFound, ok := registryCasted.validators[notNilValidatorName]

					Expect(ok).To(BeTrue())
					Expect(notNilValidatorFound).ToNot(BeNil())
				})
			})

			Context("and a validator already exists in a registry", func() {
				It("should just skip it", func() {
					// given
					notNilValidator := NewNotNilValidator()
					notNilValidatorName := notNilValidator.GetName()

					newValidatorWithSameName := &testValidator{name: notNilValidatorName}

					// when / then
					registryCasted, ok := registry.(*SimpleConstraintValidatorRegistry)

					registry.Add(notNilValidator)
					registry.Add(newValidatorWithSameName)

					Expect(registryCasted.validators).To(HaveLen(1))

					notNilValidatorFound, ok := registryCasted.validators[notNilValidatorName]

					Expect(ok).To(BeTrue())
					Expect(notNilValidatorFound).To(Equal(notNilValidator))
				})
			})
		})

		When("a removing is going", func() {
			var registry ConstraintValidatorRegistry

			BeforeEach(func() {
				registry = NewEmptySimpleConstraintValidatorRegistry()
			})

			Context("and a validator exists in a registry", func() {
				It("should successfully remove it", func() {
					// given
					notNilValidator := NewNotNilValidator()
					registry.Add(notNilValidator)

					// when / then
					registryCasted, ok := registry.(*SimpleConstraintValidatorRegistry)

					Expect(ok).To(BeTrue())
					Expect(registryCasted.validators).To(HaveLen(1))

					registry.Remove(notNilValidator)
					Expect(registryCasted.validators).To(HaveLen(0))
				})
			})

			Context("and a validator doesn't exist in a registry", func() {
				It("should just skip it", func() {
					// given
					notNilValidator := NewNotNilValidator()

					// when / then
					registryCasted, ok := registry.(*SimpleConstraintValidatorRegistry)

					Expect(ok).To(BeTrue())
					Expect(registryCasted.validators).To(HaveLen(0))

					registry.Remove(notNilValidator)
					Expect(registryCasted.validators).To(HaveLen(0))
				})
			})
		})
	},
)

type testValidator struct {
	name string
}

func (validator *testValidator) GetName() string {
	return validator.name
}

func (validator *testValidator) Validate(value any, constraint Constraint) error {
	return nil
}
