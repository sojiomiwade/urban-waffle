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
	if len(os.Args[1:]) == 0 {
		countLines(os.Stdin, count)
	} else {
		for _, filename := range os.Args[1:] {
			fp, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error %q for file %s\n", err, filename)
			} else {
				countLines(fp, count)
				fp.Close()
			}
		}
	}
	for line, linecount := range count {
		if linecount > 1 {
			fmt.Printf("%d\t%s\n", linecount, line)
		}
	}
}

func countLines(file *os.File, count map[string]int) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count[scanner.Text()] += 1
	}
}
