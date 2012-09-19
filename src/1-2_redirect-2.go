/**
 * p.6
 */

package main

import (
	"log"
	"io"
	"os"
)

var BUFFSIZE int = 8192

func main() {
	buf := make([]byte, BUFFSIZE)
	r := io.TeeReader(os.Stdin, os.Stdout)
	for {
		if _, err := r.Read(buf); err != nil {
			log.Fatal("read error")
		}
	}
}