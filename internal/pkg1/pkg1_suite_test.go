package pkg1_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPkg1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pkg1 Suite")
}
