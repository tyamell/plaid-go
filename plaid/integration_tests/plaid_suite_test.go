package plaid_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPlaid(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Plaid Suite")
}
