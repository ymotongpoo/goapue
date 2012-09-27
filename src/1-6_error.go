package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	fmt.Fprintf(os.Stderr, "EACCES: %s\n", os.ErrPermission)

	log.Fatal(os.ErrNotExist)
}