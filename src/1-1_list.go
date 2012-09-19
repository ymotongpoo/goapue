/**
 * p.4
 */

package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("a single argument (the directory name) is required")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("can't open %s\n", os.Args[1])
	}

	fis, err := file.Readdir(-1)
	if err == nil {
		for i := range fis {
			fmt.Printf("%s %s\n", fis[i].Mode(), fis[i].Name())
		}
	}

	file.Close()
	os.Exit(0)
}
