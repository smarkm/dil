package ns

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNS(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ns")
}
