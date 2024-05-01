// gopl.io/ch1/dup1

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

// Now I try to do above myself in Go with help from Coderpad

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		count[scanner.Text()] += 1
	}
	for line, linecount := range count {
		if linecount > 1 {
			fmt.Printf("%d\t%s\n", linecount, line)
		}
	}
}
