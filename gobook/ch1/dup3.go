// gopl.io/ch1/dup1

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

// Now I try to do above myself in Go with help from Coderpad

package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

func main() {
  count := make(map[string]int)
  fnames := os.Args[1:]
  if len(fnames) > 0 {

		fmt.Println(fnames)
    for _, filename := range fnames {
      data, err := ioutil.ReadFile(filename)
      if err != nil {
        fmt.Fprintf(os.Stderr, "Error %q\n for file %s\n", err, filename)
      } else {
lines:=strings.Split(string(data),"\n")
fmt.Println(lines,len(lines))
        for _, line := range lines {
          count[line] ++
        }
      }
    }

    for line, linecount := range count {
      if linecount > 1 {
        fmt.Printf("%d\t%s\n", linecount, line)
      }
    }
  }
}
