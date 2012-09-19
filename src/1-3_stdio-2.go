/**
 * p.6
 */

package main

import (
	"os"
	"log"
	"bufio"
)

func main() {
	rb := bufio.NewReader(os.Stdin)
	wb := bufio.NewWriter(os.Stdout)
	for {
		c, err := rb.ReadByte()
		if err != nil {
			log.Fatal("input error")
		}
		if err := wb.WriteByte(c); err != nil {
			log.Fatal("output error")
		} else {
			wb.Flush()
		}
	}
	os.Exit(0)
}
