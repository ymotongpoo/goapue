/**
 * p.7
 */

package main

import (
	"os"
)

func main() {
	var b = make([]byte, 1)
	for {
		_, err := os.Stdin.Read(b)
		if err != nil {
			os.Exit(1)
		}
		_, err = os.Stdout.Write(b)
		if err != nil {
			os.Exit(1)
		}
	}
	os.Exit(0)
}