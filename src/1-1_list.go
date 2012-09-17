package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("a single argument (the directory name) i.s required")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("can't open %s\n", os.Args[1])
		os.Exit(1)
	}

	fis, err := file.Readdir(-1)
	if err == nil {
		for i := range(fis) {
			fmt.Printf("%s %s\n", fis[i].Mode(), fis[i].Name())
		}
	}

	file.Close()
	os.Exit(0)
}