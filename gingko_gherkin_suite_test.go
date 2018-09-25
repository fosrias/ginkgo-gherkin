package ginkgogherkin_test

import (
	"testing"

	. "github.com/fosrias/gingko-gherkin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGingkoGherkin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunGherkinSpecs(t, "GingkoGherkin Suite")
}
