package main

import (
	"log"
	"os"
)

var BUFFSIZE int = 8192

func main() {
	buf := make([]byte, BUFFSIZE)

	for {
		n, err := os.Stdin.Read(buf)
		if err != nil || n < 0 {
			log.Fatal("read error")
		}
		if m, err := os.Stdout.Write(buf[0:n]); err != nil && m == n {
			log.Fatal("write error")
		}
	}

	os.Exit(0)
}
