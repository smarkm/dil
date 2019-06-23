package tsync_test

import (
	"log"
	"sync"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tsync", func() {

	It("test sync", func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go (func() {
			defer wg.Done()
			time.Sleep(5 * time.Second)

		})()
		wg.Wait()
		log.Println("down")
	})
})

func TestNS(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ns")
}
