package ns

import (
	"fmt"
	"log"
	"os"
	"path"

	"golang.org/x/sys/unix"
)

var nsRootPath = "/var/run/mns"

func NewNS(ns string) (*os.File, error) {
	err := os.MkdirAll(nsRootPath, 0755)
	if err != nil {
		log.Println(err)
	}
	nsPath := path.Join(nsRootPath, ns)
	fd, err := os.Create(nsPath)

	err = unix.Unshare(unix.CLONE_NEWNET)
	if err != nil {
		log.Println(err)
	}
	log.Println(getCurrentThreadNetNSPath())
	err = unix.Mount(getCurrentThreadNetNSPath(), nsPath, "none", unix.MS_BIND, "")
	if err != nil {
		err = fmt.Errorf("failed to bind mount ns at %s: %v", nsPath, err)
	}
	return fd, err
}

// getCurrentThreadNetNSPath copied from pkg/ns
func getCurrentThreadNetNSPath() string {
	// /proc/self/ns/net returns the namespace of the main thread, not
	// of whatever thread this goroutine is running on.  Make sure we
	// use the thread's net namespace since the thread is switching around
	return fmt.Sprintf("/proc/%d/task/%d/ns/net", os.Getpid(), unix.Gettid())
}
