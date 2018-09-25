package ginkgogherkin_test

import (
	. "github.com/fosrias/gingko-gherkin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Feature("Do Something",
	`So that I can accomplish some business value,
     As an actor,
	 I can perform some action`,
	func() {
		var value int

		BeforeEach(func() {

		})

		Scenario("Success", func() {
			Given("I try something", func() {
				value += 1
			})
			And("I try harder", func() {
				value += 3
			})
			Then("it works", func() {
				Expect(value).To(Equal(4))
			})
		})
	})
