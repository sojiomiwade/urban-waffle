package main

import (
	"fmt"
	"os"
	"time"
)

// open a file for writing
// write some stuff
// check if it is there
func main() {
	of := os.Args[1]
	ofile, _ := os.Create(of)
	for {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(ofile, "hi")
	}
}
