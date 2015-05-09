package password_reader_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPasswordReader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PasswordReader Suite")
}
