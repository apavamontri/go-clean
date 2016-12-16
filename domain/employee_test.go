package domain_test

import (
	. "github.com/apavamontri/go-clean/domain"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Employee", func() {
	Describe("Employee in string", func() {
		Context("Empty Employee", func() {
			It("should have empty string display", func() {
				emp := Employee{}

				Expect(emp.String()).To(Equal(""))
			})
		})

		Context("Normal employee", func() {
			It("should have correct string display", func() {
				emp := Employee{
					ID:        1,
					FirstName: "John",
					LastName:  "Doe",
				}

				Expect(emp.String()).To(Equal("1: John Doe"))
			})
		})
	})
})
