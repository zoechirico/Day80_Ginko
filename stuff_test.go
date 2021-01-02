package Day80_Ginko_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"example.com/m/v2"
)

var _ = Describe("Stuff", func() {

	Context("initially", func() {

		s := Day80_Ginko.S{}
		s.Name = "Zoe"

		It("Has to have name", func() {
			Expect(s.Name).Should(Equal("Zoe"))
		})

	})

})
