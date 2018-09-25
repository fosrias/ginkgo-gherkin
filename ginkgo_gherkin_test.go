package ginkgogherkin_test

import (
	gg "github.com/fosrias/gingko-gherkin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = gg.Feature("Do Something",
	`So that I can accomplish some business value,
     As an actor,
	 I can perform some action`,
	func() {
		var value int

		BeforeEach(func() {

		})

		gg.Scenario("Success", func() {
			gg.Given("I try something", func() {
				value += 1
			})
			gg.And("I try harder", func() {
				value += 3
			})
			gg.Then("it works", func() {
				Expect(value).To(Equal(4))
			})
		})
	})
