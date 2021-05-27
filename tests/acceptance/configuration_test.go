package acceptance

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("NRI Redis configuration", func() {

	Describe("Configuration is not valid", func() {
		Context("Missing required parameters", func() {
			It("should fail the execution", func() {
				//Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})
	})
})
