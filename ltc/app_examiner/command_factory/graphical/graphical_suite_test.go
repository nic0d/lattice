package graphical_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGraphical(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Graphical Suite")
}
