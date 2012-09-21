/**
 * p.8
 */

package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Printf("hello world from process ID %d\n", os.Getpid());
	os.Exit(0)
}