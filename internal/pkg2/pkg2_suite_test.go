package pkg2_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPkg2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pkg2 Suite")
}
