/**
 * p.6
 */

package main

import (
	"syscall"
)

var BUFFSIZE int = 8192

func main() {
	buf := make([]byte, BUFFSIZE)
	for {
		n, err := syscall.Read(syscall.Stdin, buf)
		if err != nil || n < 0 {
			syscall.Exit(1)
		}
		if m, err := syscall.Write(syscall.Stdout, buf[0:n]); err != nil && m == n {
			syscall.Exit(1)
		}
	}
	syscall.Exit(0)
}

