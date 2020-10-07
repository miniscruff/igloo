package igloo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIgloo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Igloo Suite")
}
