package pkg2_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/captain-examples/go-ginkgo/internal/pkg2"
)

var _ = Describe("Foo", func() {
	It("passes", func() {
		Expect(pkg2.Foo("arg")).To(Equal("Foo: arg"))
	})
})
