Siberia Validation
=================

[![Author](https://img.shields.io/badge/author-@siberia_projects-green.svg)](https://github.com/siberia-projects)
[![Source Code](https://img.shields.io/badge/source-siberia/main-blue.svg)](https://github.com/siberia-projects/siberia-validation)

## What is it?
Siberia-validation is a library for validation an input

## How to download?

```console
john@doe-pc:~$ go get github.com/siberia-projects/siberia-validation
```

## How to use?
Will be described later

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
		&validation.NotNilConstraint{},
		&validation.NotBlankConstraint{},
	})

	if err != nil {
		println(err.Error())
		return
	}

	println("Everything is valid!")
}
```
