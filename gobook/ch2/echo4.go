// Echo4 prints its command-line arguments
package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep = flag.String("s", " ", "separator")
var n = flag.Bool("n", false, "omit trailing newline")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*sep))
	if !*n {
		fmt.Println()
	}
}
