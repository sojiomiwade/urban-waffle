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
	for _,name := range names {

	}
}
/*
a 1 2 1 
c 1 2 3
b 1 2 1
d 5 1
names: a b
5 1 [a,c,b,d]
3 2 [a,b,c]

1 a {a:}
*/

func getNames(count map[string]int, filenames []string) {

	for _,name in range filenames
}

func countLines(file *os.File, count map[string]int, filename map[string]bool) {
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
		word=scanner.Text()
    count[word] += 1
		filename[word]=true
  }
}
