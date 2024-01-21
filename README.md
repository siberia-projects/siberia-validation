Siberia Validation
=================

[![Author](https://img.shields.io/badge/author-@siberia_projects-green.svg)](https://github.com/siberia-projects)
[![Source Code](https://img.shields.io/badge/source-siberia/main-blue.svg)](https://github.com/siberia-projects/siberia-validation)

## What is it?
Siberia-validation is a library for a simple and convenient validation a passed value

It's pretty easy to use and extend if you need a custom validation logic

## How to download?

```console
john@doe-pc:~$ go get github.com/siberia-projects/siberia-validation
```

## How to use?
To use it you have to do the following: 
 - Create all necessary **ConstraintValidator(s)**
 - Create a **SimpleConstraintValidatorRegistry**
 - Put them into the **SimpleConstraintValidatorRegistry**
 - Create a **SimpleValidator** with passing the registry
 - Execute a method **Validate()** with passing a value you want to validate
and all needed **Constraint(s)**

For your convenience the **NewSimpleConstraintValidatorRegistryWithDefaultValidators** exists

If a value is invalid you will have a **ConstraintViolationError** which holds every violation error inside itself

## Examples
```go
package main

import (
	"github.com/siberia-projects/siberia-validation/pkg/validation"
)

func main() {
	registry := validation.NewSimpleConstraintValidatorRegistryWithDefaultValidators()
	validator := validation.NewSimpleValidator(registry)

	err := validator.Validate("valid_value", []validation.Constraint{
		&validation.NotNil{},
		&validation.NotBlank{},
	})

	if err != nil {
		println(err.Error())
		return
	}

	println("Everything is valid!")
}
```
