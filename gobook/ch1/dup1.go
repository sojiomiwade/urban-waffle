package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	count := make(map[string]int)
	files := make(map[string]map[string]bool)
	if len(os.Args[1:]) == 0 {
		countLines(os.Stdin, count, files, "Stdin")
	} else {
		for _, filename := range os.Args[1:] {
			fp, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error %q for file %s\n", err, filename)
			} else {
				countLines(fp, count, files, filename)
				fp.Close()
			}
		}
	}
	for line, linecount := range count {
		if linecount > 1 {
			keys := make([]string, 0, len(files[line]))
			for k := range files[line] {
				keys = append(keys, k)
			}
			fmt.Printf("%d\t%10s\t%s\n", linecount, line, keys)
		}

	}
}

/*
jj {filefoo:true}
kk {filefoo}
jj {filefoo}
kk
jj
jj
*/
func countLines(
	file io.Reader,
	count map[string]int,
	files map[string]map[string]bool,
	filename string,
) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		count[word] += 1
		if files[word] == nil {
			files[word] = make(map[string]bool)
		}
		files[word][filename] = true
	}
}
