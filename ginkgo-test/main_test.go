package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}

//var _ = Describe("Test functions",
//	func() {
//		Context("sample", func() {
//			It("one should return 1", func() {
//				Expect(1).To(Equal(one()))
//			})
//			It("name should return name", func() {
//				Expect("name").To(Equal(name()))
//			})
//		})
//	},
//)

var _ = Describe("Shopping cart", func() {
	Context("initially", func() {
		It("has 0 items", func() {})
		It("has 0 units", func() {})
		Specify("the total amount is 0.00", func() {})
	})
})