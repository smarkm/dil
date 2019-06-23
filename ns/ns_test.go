package ns

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"syscall"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vishvananda/netns"
)

var _ = Describe("Ns", func() {
	It("create ns use netns lib", func() {
		// Lock the OS Thread so we don't accidentally switch namespaces
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		// Save the current network namespace
		origns, err := netns.Get()
		Expect(err).NotTo(HaveOccurred())
		defer origns.Close()
		log.Println(os.Getpid(), syscall.Gettid(), origns)
		// Create a new network namespace
		newns, err := netns.New()
		Expect(err).NotTo(HaveOccurred())
		netns.Set(newns)
		log.Println(os.Getpid(), syscall.Gettid(), newns)
		defer newns.Close()

		// Do something with the network namespace
		ifaces, _ := net.Interfaces()
		fmt.Printf("Interfaces: %v\n", ifaces)

		// Switch back to the original namespace
		netns.Set(origns)
	})

	// It("test create ns", func() {
	// 	fd, err := NewNS(testNS)
	// 	Expect(err).NotTo(HaveOccurred())
	// 	Expect(fd.Name()).To(Equal(path.Join(nsRootPath, testNS)))

	// })

})
