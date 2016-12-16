package domain_test

import (
	. "github.com/apavamontri/go-clean/domain"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Employee", func() {
	It("should be blank", func() {
		employee := Employee{}
		employee.ID = 1

		Expect("1").To(Equal("XXX"))
	})
})
