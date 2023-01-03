package pkg1_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/captain-examples/go-ginkgo/internal/pkg1"
)

var _ = Describe("Bar", func() {
	It("passes", func() {
		Expect(pkg1.Bar("arg")).To(Equal("Bar: arg"))
	})
})
