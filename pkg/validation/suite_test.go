// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package validation

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestValidation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Runs all validation tests")
}
