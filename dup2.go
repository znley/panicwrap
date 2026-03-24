//go:build linux && (386 || amd64)
// +build linux,386 linux,amd64

package panicwrap

import (
	"syscall"
)

func dup2(oldfd, newfd int) error {
	return syscall.Dup2(oldfd, newfd)
}
