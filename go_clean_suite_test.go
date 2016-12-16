package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoClean(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoClean Suite")
}
