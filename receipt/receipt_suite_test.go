package receipt_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestReceipt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Receipt Suite")
}
