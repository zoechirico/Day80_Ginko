package Day80_Ginko_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"example.com/m/v2"
)

var _ = Describe("Stuff", func() {

	s := Day80_Ginko.S{}
	s.A = 3

	Context("With more than 300 pages", func() {
		It("should be a novel", func() {
			Expect(s.A).To(Equal(3))
		})
	})

})
