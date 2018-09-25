/*
Ginkgo-gherkin Gingko-gherkin adds some gherkin sugar to the Ginkgo BDD Testing Framework.

Ginkgo-gherkin on Github: http://github.com/fosrias/ginkgo-gherkin

Ginkgo-gherkin is MIT-Licensed
*/

package ginkgogherkin

import (
	"bufio"
	"strings"

	"github.com/fosrias/gingko-gherkin/reporters/stenographer"
	ginkgo "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
)

const (
	feature  = "Feature: "
	scenario = "Scenario: "
	given    = "Given "
	when     = "When "
	then     = "Then "
	and      = "And "
	but      = "But "
	padding  = "  "
)

//RunGherkinSpecs is the entry point for the Ginkgo-gherkin test runner.
//You must call this within a Golang testing TestX(t *testing.T) function.
//
//To bootstrap a test suite you can use the GinkgoGherkin CLI:
//
//	gingkgo-gherkin bootstrap
func RunGherkinSpecs(t ginkgo.GinkgoTestingT, description string) bool {
	remoteReportingServer := config.GinkgoConfig.StreamHost
	if remoteReportingServer == "" {
		return ginkgo.RunSpecsWithCustomReporters(t, description, []ginkgo.Reporter{buildGherkinReporter()})
	}
	return ginkgo.RunSpecs(t, description)
}

func buildGherkinReporter() ginkgo.Reporter {
	stenographer := stenographer.New(!config.DefaultReporterConfig.NoColor, config.GinkgoConfig.FlakeAttempts > 1)
	return reporters.NewDefaultReporter(config.DefaultReporterConfig, stenographer)
}

//Feature blocks allow you to organize your functional specs.  A Feature block can contain any number of
//BeforeEach, AfterEach, JustBeforeEach, Scenario, Given, When, Then, And, But, It, and Measurement blocks.
//
//In addition you can nest Scenario, Describe and Context blocks.  Scenario, Describe and Context blocks are
//functionally equivalent.  The difference is purely semantic -- you typically Describe the behavior of an object
//or method and, within that Describe, outline a number of Contexts and use Feature blocks expressing
//functional scenarios.
func Feature(text string, description string, body func()) bool {
	return ginkgo.Describe(feature+padMultilineDescription(text, description), body)
}

func padMultilineDescription(text string, description string) string {
	if description == "" {
		return text
	}

	scanner := bufio.NewScanner(strings.NewReader(description))
	paddedDescription := "\n"

	for scanner.Scan() {
		paddedDescription += padding + strings.TrimSpace(scanner.Text()) + "\n"
	}
	return text + paddedDescription + "\n"
}

//FFeature focuses tests within a feature block
func FFeature(text string, description string, body func()) bool {
	return ginkgo.FDescribe(feature+padMultilineDescription(text, description), body)
}

//PFeature marks tests within a feature block as pending
func PFeature(text string, description string, body func()) bool {
	return ginkgo.PDescribe(feature+padMultilineDescription(text, description), body)
}

//XFeature marks tests within a feature block as pending
func XFeature(text string, description string, body func()) bool {
	return ginkgo.XDescribe(feature+padMultilineDescription(text, description), body)
}

//Scenario blocks allow you to organize your functional specs.  A Feature block can contain any number of
//BeforeEach, AfterEach, JustBeforeEach, Scenario, Given, When, Then, And, But, It, and Measurement blocks.
//
//In addition you can nest Scenario, Describe and Context blocks.  Scenario, Describe and Context blocks are
//functionally equivalent.  The difference is purely semantic -- you typically Describe the behavior of an object
//or method and, within that Describe, outline a number of Contexts and use Feature blocks expressing
//functional scenarios.
func Scenario(text string, body func()) bool {
	return ginkgo.Context(scenario+text, body)
}

//FScenario focuses tests within a scenario block
func FScenario(text string, description string, body func()) bool {
	return ginkgo.FContext(scenario+text, body)
}

//PScenario marks tests within a scenario block as pending
func PScenario(text string, description string, body func()) bool {
	return ginkgo.PContext(scenario+text, body)
}

//XScenario marks tests within a scenario block as pending
func XScenario(text string, description string, body func()) bool {
	return ginkgo.XContext(scenario+text, body)
}

//Given blocks allow you to organize your functional specs.  A Feature block can contain any number of
//BeforeEach, AfterEach, JustBeforeEach, Scenario, Given, When, Then, And, But, It, and Measurement blocks.
//
//In addition you can nest Scenario, Describe and Context blocks.  Scenario, Describe and Context blocks are
//functionally equivalent.  The difference is purely semantic -- you typically Describe the behavior of an object
//or method and, within that Describe, outline a number of Contexts and use Feature blocks expressing
//functional scenarios.
func Given(text string, body interface{}, timeout ...float64) bool {
	return ginkgo.It(given+text, body, timeout...)
}

//FGiven focuses individual Givens
func FGiven(text string, body interface{}, timeout ...float64) bool {
	return ginkgo.It(given+text, body, timeout...)
}

//PGiven marks individual Givens as pending
func PGiven(text string, body interface{}, timeout ...float64) bool {
	return ginkgo.It(given+text, body, timeout...)
}

//XGiven marks individual Givens as pending
func XGiven(text string, body interface{}, timeout ...float64) bool {
	return ginkgo.It(given+text, body, timeout...)
}

//When blocks allow you to organize your functional specs.  A Feature block can contain any number of
//BeforeEach, AfterEach, JustBeforeEach, Scenario, Given, When, Then, And, But, It, and Measurement blocks.
//
//In addition you can nest Scenario, Describe and Context blocks.  Scenario, Describe and Context blocks are
//functionally equivalent.  The difference is purely semantic -- you typically Describe the behavior of an object
//or method and, within that Describe, outline a number of Contexts and use Feature blocks expressing
//functional scenarios.
func When(text string, body interface{}, timeout ...float64) bool {
	return ginkgo.It(when+text, body, timeout...)
}

//Then blocks allow you to organize your functional specs.  A Feature block can contain any number of
//BeforeEach, AfterEach, JustBeforeEach, Scenario, Given, When, Then, And, But, It, and Measurement blocks.
//
//In addition you can nest Scenario, Describe and Context blocks.  Scenario, Describe and Context blocks are
//functionally equivalent.  The difference is purely semantic -- you typically Describe the behavior of an object
//or method and, within that Describe, outline a number of Contexts and use Feature blocks expressing
//functional scenarios.
func Then(text string, body interface{}, timeout ...float64) bool {
	return ginkgo.It(then+text, body, timeout...)
}

func FThen(text string, body interface{}, timeout ...float64) bool {
	return ginkgo.FIt(then+text, body, timeout...)
}

func PThen(text string, body interface{}) bool {
	return ginkgo.PIt(then+text, body, 0)
}

func XThen(text string, body interface{}) bool {
	return ginkgo.XIt(then+text, body, 0)
}

//AnD blocks allow you to organize your functional specs.  A Feature block can contain any number of
//BeforeEach, AfterEach, JustBeforeEach, Scenario, Given, When, Then, And, But, It, and Measurement blocks.
//
//In addition you can nest Scenario, Describe and Context blocks.  Scenario, Describe and Context blocks are
//functionally equivalent.  The difference is purely semantic -- you typically Describe the behavior of an object
//or method and, within that Describe, outline a number of Contexts and use Feature blocks expressing
//functional scenarios.
func And(text string, body interface{}, timeout ...float64) bool {
	return ginkgo.It(and+text, body, timeout...)
}

//But blocks allow you to organize your functional specs.  A Feature block can contain any number of
//BeforeEach, AfterEach, JustBeforeEach, Scenario, Given, When, Then, And, But, It, and Measurement blocks.
//
//In addition you can nest Scenario, Describe and Context blocks.  Scenario, Describe and Context blocks are
//functionally equivalent.  The difference is purely semantic -- you typically Describe the behavior of an object
//or method and, within that Describe, outline a number of Contexts and use Feature blocks expressing
//functional scenarios.
func But(text string, body interface{}, timeout ...float64) bool {
	return ginkgo.It(but+text, body, timeout...)
}
