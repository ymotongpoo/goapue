package main

import (
	"os"
	"os/exec"
	"fmt"
	"log"
)

func main() {
	for {
		fmt.Printf("%% ")
		
		var buf string;
		_, err := fmt.Scanln(&buf)
		if err != nil {
			log.Fatal("read error")
		}

		cmd := exec.Command(buf)
		output, err := cmd.Output();
		if err != nil {
			log.Fatal("couldn't execute ", buf)
		}
		fmt.Printf("%v", string(output))
	}
	os.Exit(0)
}